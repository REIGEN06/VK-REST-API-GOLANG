package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/reigen06/vk-rest-api/docs"
	"github.com/reigen06/vk-rest-api/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
)

type Handler struct {
	services *service.Service
	logger   *slog.Logger
}

func NewHandler(services *service.Service, logger *slog.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) SetRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// @TODO: change hardcode URL
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4005/swagger/doc.json"),
	))

	router.Route("/api", func(r chi.Router) {

		r.Route("/actor", func(r chi.Router) {
			r.Post("/", h.CreateActor)
			r.Get("/all", h.GetAllActorsWithMovies)
			r.Put("/{id}", h.UpdateActor)
			r.Delete("/{id}", h.DeleteActor)
		})

		r.Route("/movie", func(r chi.Router) {
			r.Post("/", h.CreateMovie)
			r.Put("/{id}", h.UpdateMovie)
			r.Delete("/{id}", h.DeleteMovie)

			r.Route("/all", func(r chi.Router) {
				r.Get("/{sort_by}", h.GetSortedMovies)
				r.Get("/{movie_name}", h.GetByMovieName)
				r.Get("/{actor_name}", h.GetByActorName)
			})
		})
	})

	return router
}
