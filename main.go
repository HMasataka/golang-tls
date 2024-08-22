package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", helloWorld)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("hello, world")); err != nil {
		log.Print(err)
	}
}
