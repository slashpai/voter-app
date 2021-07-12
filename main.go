package main

import (
	"github.com/slashpai/voter-app/pkg/redis"
	"html/template"
	"log"
	"net/http"
)

type Result struct {
	DogVote     int
	CatVote     int
	NeutralVote int
	Success     bool
}

func main() {
	indexTmpl := template.Must(template.ParseFiles("views/index.html"))
	resultTmpl := template.Must(template.ParseFiles("views/results.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Starting request")
		if r.Method != http.MethodPost {
			err := indexTmpl.Execute(w, nil)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	})

	http.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Vote for dog")
		redis.VoteDog()
	})

	http.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Vote for cat")
		redis.VoteCat()
	})

	http.HandleFunc("/neutral", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Vote for neutral")
		redis.VoteNeutral()
	})

	http.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Vote Results")
		res := redis.VoteResult()
		details := Result{
			DogVote:     res["dog"],
			CatVote:     res["cat"],
			NeutralVote: res["neutral"],
			Success:     true,
		}
		log.Print("details", details)
		err := resultTmpl.Execute(w, details)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Print("Application is available at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
