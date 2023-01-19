package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request)  {
	// fmt.Fprintf(w, "Hello world from %s", app.Domain)

	var payload = struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`

	}{
		Status: "active",
		Message: "Go Movies up and running",
		Version: "6.6.6",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request)  {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movie {

		ID: 1,
		Title: "Highlander",
		ReleaseDate: rd,
		MPAARating: "R",
		RunTime: 116,
		Description: "A movie brat",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	}
	movies = append(movies, highlander)


	rd, _ = time.Parse("2006-01-02", "1981-06-12")

	rotla := models.Movie {

		ID: 1,
		Title: "Raiders of the Lost Ark",
		ReleaseDate: rd,
		MPAARating: "PG-13",
		RunTime: 115,
		Description: "Another  movie brat",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	}

	movies = append(movies, rotla)


	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}