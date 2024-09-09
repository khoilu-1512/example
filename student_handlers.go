package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func initDB() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// cSpell:disable
	dbSSLMode := os.Getenv("DB_SSLMODE")
	// cSpell:disable
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbSSLMode)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
			log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
			log.Fatal(err)
	}
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, age FROM students")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error fetching students")
		return
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Age); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error scanning student")
			return
		}
		students = append(students, s)
	}

	respondWithJSON(w, http.StatusOK, students)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := db.QueryRow("INSERT INTO students (name, age) VALUES ($1, $2) RETURNING id", student.Name, student.Age).Scan(&student.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating student")
		return
	}

	respondWithJSON(w, http.StatusCreated, student)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var student Student
	err := db.QueryRow("SELECT id, name, age FROM students WHERE id = $1", id).Scan(&student.ID, &student.Name, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Error fetching student")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err := db.Exec("UPDATE students SET name = $1, age = $2 WHERE id = $3", student.Name, student.Age, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating student")
		return
	}

	student.ID = id
	respondWithJSON(w, http.StatusOK, student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := db.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting student")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error checking deleted rows")
		return
	}

	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Student not found")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}