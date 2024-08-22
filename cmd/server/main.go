package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", helloWorld)

	if err := http.ListenAndServeTLS(":443", os.Getenv("CRT_FILE"), os.Getenv("KEY_FILE"), r); err != nil {
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("hello, world")); err != nil {
		log.Print(err)
	}
}
