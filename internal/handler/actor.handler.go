package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/reigen06/vk-rest-api/internal/models"
	"net/http"
	"strconv"
)

// @Summary Create actor
// @Tags actor
// @Description create actor
// @ID create-actor
// @Accept  json
// @Produce  json
// @Param input body models.Actor true "actor info"
// @Success 201 "OK"
// @Failure 400
// @Failure 500
// @Router /api/actor [post]
func (h *Handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	// @TODO: duplicate code, move it to a separate function
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusBadRequest, "invalid input body", err)
		return
	}

	// @TODO: duplicate code, move it to a separate function
	var actor models.Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		newErrorResponse(w, h.logger, http.StatusBadRequest, "invalid data in request body", err)
		return
	}

	if err := h.services.Actor.Create(&actor); err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to create actor", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary Get actors
// @Tags actor
// @Description get actors
// @ID get-actors
// @Accept json
// @Produce  json
// @Success 200 {object} []models.Actor
// @Failure 500
// @Router /api/actor/all [get]
func (h *Handler) GetAllActorsWithMovies(w http.ResponseWriter, r *http.Request) {
	actorsWithMovies, err := h.services.GetAllWithMovies()
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to get actors with movies", err)
		return
	}

	jsonResponse, _ := json.MarshalIndent(actorsWithMovies, "", " ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary Update actor
// @Tags actor
// @Description update actor
// @ID update-actor
// @Accept json
// @Produce  json
// @Param id path int true "actor id"
// @Success 200 "OK"
// @Failure 400
// @Failure 500
// @Router /api/actor/{id} [put]
func (h *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusBadRequest, "failed to read body", err)
		return
	}

	var actor models.Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		newErrorResponse(w, h.logger, http.StatusBadRequest, "invalid data in request body", err)
		return
	}

	requestedId := chi.URLParam(r, "id")
	actorId, err := strconv.Atoi(requestedId)

	if err := h.services.Actor.Update(actorId, &actor); err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to update actor", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}

// @Summary Delete actor
// @Tags actor
// @Description delete actors
// @ID delete-actors
// @Accept json
// @Produce  json
// @Param id path int true "actor id"
// @Success 200 "OK"
// @Failure 500
// @Router /api/actor/{id} [delete]
func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	requestedId := chi.URLParam(r, "id")
	actorId, _ := strconv.Atoi(requestedId)

	if err := h.services.Actor.Delete(actorId); err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to delete actor", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		newErrorResponse(w, h.logger, http.StatusInternalServerError, "failed to write response", err)
		return
	}
}
