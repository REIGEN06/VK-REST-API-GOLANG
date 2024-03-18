package repository

import (
	"github.com/reigen06/vk-rest-api/internal/models"
	"gorm.io/gorm"
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

type Repository struct {
	Actor
	Movie
}

func NewRepository(db *gorm.DB, logger *slog.Logger) *Repository {
	return &Repository{
		Actor: NewActorPostgresRepository(db, logger),
		Movie: NewMoviePostgresRepository(db, logger),
	}
}
