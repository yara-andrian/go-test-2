package main

import (
	"log"
	"net/http"

	"./stringutil"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stringutil.Reverse("DBS")))
	})

	r.Get("/a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stringutil.Reverse("!oG ,olleH")))
	})

	r.Get("/b", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stringutil.Reverse("!ob")))
	})

	log.Println("Server listening to PORT: 8081")

	http.ListenAndServe(":8080", r)
}
