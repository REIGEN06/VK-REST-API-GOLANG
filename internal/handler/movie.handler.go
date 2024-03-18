package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/reigen06/vk-rest-api/internal/models"
	"net/http"
	"strconv"
)

// @Summary Create movie
// @Tags movie
// @Description create movie
// @ID create-movie
// @Accept  json
// @Produce  json
// @Param input body models.Movie true "movie info"
// @Success 201 "OK"
// @Failure 400
// @Failure 500
// @Router /api/movie [post]
func (h *Handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	// @TODO: duplicate code, move it to a separate function
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// @TODO: duplicate code, move it to a separate function
	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to create movie", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary search with sort (default = DESC)
// @Tags movie
// @Description get movies: with sort (default = DESC)
// @ID get-movies-sortby
// @Accept json
// @Produce  json
// @Success 200 {object} []models.Movie
// @Failure 500
// @Router /api/movie/all/{sort_by} [get]
func (h *Handler) GetSortedMovies(w http.ResponseWriter, r *http.Request) {
	queryValue := r.URL.Query().Get("sort_by")
	movies, err := h.services.Movie.GetSorted(queryValue)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to get sorted movies", err)
		return
	}

	jsonResponse, _ := json.MarshalIndent(movies, "", " ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary Search by fragment of a movie's name
// @Tags movie
// @Description get movies: search by fragment of a movie's name
// @ID get-movies-movie-name
// @Accept json
// @Produce  json
// @Success 200 {object} []models.Movie
// @Failure 500
// @Router /api/movie/all/{movie_name} [get]
func (h *Handler) GetByMovieName(w http.ResponseWriter, r *http.Request) {
	requestedName := chi.URLParam(r, "movie_name")
	movies, err := h.services.Movie.GetByMovieName(requestedName)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to get movies by name", err)
		return
	}

	jsonResponse, _ := json.MarshalIndent(movies, "", " ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary Search by fragment of an actor's name
// @Tags movie
// @Description get movies: search by fragment of an actor's name
// @ID get-movies
// @Accept json
// @Produce  json
// @Success 200 {object} []models.Movie
// @Failure 500
// @Router /api/movie/all/{actor_name} [get]
func (h *Handler) GetByActorName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "actor_name")
	movies, err := h.services.Movie.GetByActorName(name)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "Failed to get movies by actor name", err)
		return
	}

	jsonResponse, _ := json.MarshalIndent(movies, "", " ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary Update movie
// @Tags movie
// @Description update movie
// @ID update-movie
// @Accept json
// @Produce  json
// @Param id path int true "movie id"
// @Success 200 "OK"
// @Failure 400
// @Failure 500
// @Router /api/movie/{id} [put]
func (h *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusBadRequest, "failed to read body", err)
		return
	}

	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		newErrorResponse(w, h.logger, http.StatusBadRequest, "failed to read body", err)
		return
	}

	requestedId := chi.URLParam(r, "id")
	movieId, err := strconv.Atoi(requestedId)

	if err := h.services.Movie.Update(movieId, &movie); err != nil {
		newErrorResponse(w, h.logger, http.StatusBadRequest, "failed to update movie", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary Delete movie
// @Tags movie
// @Description delete movies
// @ID delete-movies
// @Accept json
// @Produce  json
// @Param id path int true "movie id"
// @Success 200 "OK"
// @Failure 500
// @Router /api/movie/{id} [delete]
func (h *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	requestedId := chi.URLParam(r, "id")
	movieId, _ := strconv.Atoi(requestedId)

	if err := h.services.Movie.Delete(movieId); err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to delete movie", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}
