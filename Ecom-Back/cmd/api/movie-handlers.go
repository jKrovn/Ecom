package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}
	movie, err := app.models.DB.Get(id)
	if err != nil {
		fmt.Println(err)
	}
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
	}
}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJSON(w, err)
	}
	err = app.writeJSON(w, http.StatusOK, genres, "genres")
}

func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJSON(w, err)
	}
	movies, err := app.models.DB.All(genreID)
	if err != nil {
		app.errorJSON(w, err)
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {

}
