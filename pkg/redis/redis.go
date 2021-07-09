package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
)

type dbDetails struct {
	addr     string
	password string
	db       int
}

var ctx = context.Background()

var rdbDetails = NewDBDetails()

func NewDBDetails() *dbDetails {
	rdb, err := strconv.Atoi(os.Getenv("VOTER_REDIS_DB"))

	if err != nil {
		panic(err)
	}
	return &dbDetails{
		addr:     os.Getenv("VOTER_REDIS_ADDR"),
		password: os.Getenv("VOTER_REDIS_PASSWORD"),
		db:       rdb,
	}
}

func (db *dbDetails) RDBConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     db.addr,
		Password: db.password,
		DB:       db.db,
	})
}

func IncrementCount(rdb *redis.Client, key string) {
	_, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("No votes yet for", key)
		err := rdb.Set(ctx, key, 0, 0).Err()
		if err != nil {
			panic(err)
		}
	}
	logIncr := rdb.IncrBy(ctx, key, 1)
	log.Print(logIncr)
}

func VoteDog() {
	rdb := rdbDetails.RDBConnection()
	IncrementCount(rdb, "dog")
}

func VoteCat() {
	rdb := rdbDetails.RDBConnection()
	IncrementCount(rdb, "cat")
}

func VoteNeutral() {
	rdb := rdbDetails.RDBConnection()
	IncrementCount(rdb, "neutral")
}

func VoteResult() map[string]int {
	rdb := rdbDetails.RDBConnection()
	results := map[string]int{"dog": 0, "cat": 0, "neutral": 0}
	for key := range results {
		votes, err := rdb.Get(ctx, key).Result()
		results[key], _ = strconv.Atoi(votes)
		log.Print(key, " ", votes)
		if err != nil {
			log.Print(err)
		}
	}
	return results
}
