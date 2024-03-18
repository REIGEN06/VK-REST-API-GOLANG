package service

import (
	"github.com/reigen06/vk-rest-api/internal/models"
	"github.com/reigen06/vk-rest-api/internal/repository"
	"log/slog"
)

// Actor @TODO: add mocks
type Actor interface {
	Create(actor *models.Actor) error
	GetAllWithMovies() ([]models.Actor, error)
	Update(actorId int, actor *models.Actor) error
	Delete(actorId int) error
}

type Movie interface {
	Create(movie *models.Movie) error
	GetSorted(sortBy string) ([]models.Movie, error)
	GetByMovieName(name string) ([]models.Movie, error)
	GetByActorName(name string) ([]models.Movie, error)
	Update(movieId int, movie *models.Movie) error
	Delete(movieId int) error
}

type Service struct {
	Actor
	Movie
}

func NewService(repository *repository.Repository, logger *slog.Logger) *Service {
	return &Service{
		Actor: NewActorService(repository.Actor, logger),
		Movie: NewMovieService(repository.Movie, logger),
	}
}
