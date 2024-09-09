package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	initDB()

	r := chi.NewRouter()

	r.Get("/students", getStudents)
	r.Post("/students", createStudent)
	r.Get("/students/{id}", getStudent)
	r.Put("/students/{id}", updateStudent)
	r.Delete("/students/{id}", deleteStudent)

	fmt.Println("School API server starting...")
	http.ListenAndServe(":8080", r)
}
