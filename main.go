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
			indexTmpl.Execute(w, nil)
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
		resultTmpl.Execute(w, details)
	})

	log.Print("Application is available at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
