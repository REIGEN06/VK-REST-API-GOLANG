package repository

import (
	"github.com/reigen06/vk-rest-api/internal/models"
	"gorm.io/gorm"
	"log/slog"
)

type MoviePostgresRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewMoviePostgresRepository(db *gorm.DB, logger *slog.Logger) *MoviePostgresRepository {
	return &MoviePostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (r *MoviePostgresRepository) Create(movie *models.Movie) error {
	if err := r.db.Create(movie).Error; err != nil {
		r.logger.Error("error while creating movie", err)
		return err
	}
	return nil
}

// GetSorted Select the sorting criterion depending on the passed sortBy value
func (r *MoviePostgresRepository) GetSorted(sortBy string) ([]models.Movie, error) {
	var movies []models.Movie
	// @TODO: change to ENUM
	// @TODO: add order parameter
	switch sortBy {
	case "name":
		if err := r.db.Order("name DESC").Find(&movies).Error; err != nil {
			r.logger.Error("error while sorting by name", err)
			return nil, err
		}
	case "rating":
		if err := r.db.Order("rating DESC").Find(&movies).Error; err != nil {
			r.logger.Error("error while sorting by rating", err)
			return nil, err
		}
	case "release_date":
		if err := r.db.Order("release_date DESC").Find(&movies).Error; err != nil {
			r.logger.Error("error while sorting by release date", err)
			return nil, err
		}
	default:
		if err := r.db.Order("rating DESC").Find(&movies).Error; err != nil {
			r.logger.Error("error while sorting by rating (default)", err)
			return nil, err
		}
	}
	return movies, nil
}

// SearchByMovieName search by fragment of a movie's name
func (r *MoviePostgresRepository) GetByMovieName(name string) ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&movies).Error; err != nil {
		r.logger.Error("error while searching by movie name", err)
		return nil, err
	}
	return movies, nil
}

// SearchByActorName search by fragment of an actor's name
func (r *MoviePostgresRepository) GetByActorName(name string) ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Preload("Actors", "name LIKE ?", "%"+name+"%").Find(&movies).Error; err != nil {
		r.logger.Error("error while searching by actor name", err)
		return nil, err
	}
	return movies, nil
}

func (r *MoviePostgresRepository) Update(movieId int, movie *models.Movie) error {
	if err := r.db.Where("id = ?", movieId).Updates(movie).Error; err != nil {
		r.logger.Error("error while updating movie", err)
		return err
	}
	return nil
}

func (r *MoviePostgresRepository) Delete(movieId int) error {
	if err := r.db.Delete(&models.Movie{}, movieId).Error; err != nil {
		r.logger.Error("error while deleting movie", err)
		return err
	}
	return nil
}
