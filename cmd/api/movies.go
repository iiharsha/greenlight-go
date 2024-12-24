package main

import (
	"fmt"
	"net/http"
	"time"

	data "github.com/iiharsha/gogreenlight/internal"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Interstellar",
		Runtime:   169,
		Genres:    []string{"Sci-Fi", "Adeventure"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a problem and could not process your requet", http.StatusInternalServerError)
	}
}
