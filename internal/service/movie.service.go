package service

import (
	"github.com/reigen06/vk-rest-api/internal/models"
	"github.com/reigen06/vk-rest-api/internal/repository"
	"log/slog"
)

type MovieService struct {
	repository repository.Movie
	logger     *slog.Logger
}

func NewMovieService(repository repository.Movie, logger *slog.Logger) *MovieService {
	return &MovieService{
		repository: repository,
		logger:     logger,
	}
}

func (s *MovieService) Create(movie *models.Movie) error {
	s.logger.Debug("movie has been created", slog.String("name", movie.Name))

	if err := s.repository.Create(movie); err != nil {
		s.logger.Error("failed to create movie", slog.String("cause", err.Error()))
		return err
	}

	return nil
}

func (s *MovieService) GetSorted(sortBy string) ([]models.Movie, error) {
	s.logger.Debug("movies have been requested with sorting:", slog.String("sortBy", sortBy))

	movies, err := s.repository.GetSorted(sortBy)
	if err != nil {
		s.logger.Error("failed to get sorted movies", slog.String("cause", err.Error()))
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) GetByMovieName(name string) ([]models.Movie, error) {
	s.logger.Debug("searching movies by name:", slog.String("name", name))

	movies, err := s.repository.GetByMovieName(name)
	if err != nil {
		s.logger.Error("failed to search movies by name", slog.String("cause", err.Error()))
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) GetByActorName(name string) ([]models.Movie, error) {
	s.logger.Debug("searching movies by actor name:", slog.String("name", name))

	movies, err := s.repository.GetByActorName(name)
	if err != nil {
		s.logger.Error("failed to search movies by actor name", slog.String("cause", err.Error()))
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) Update(movieId int, movie *models.Movie) error {
	s.logger.Debug("movie has been updated by id:", slog.Int("movieId", movieId))

	if err := s.repository.Update(movieId, movie); err != nil {
		s.logger.Error("failed to update movie", slog.String("cause", err.Error()))
		return err
	}

	return nil
}

func (s *MovieService) Delete(movieId int) error {
	s.logger.Debug("movie has been deleted by id:", slog.Int("movieId", movieId))

	if err := s.repository.Delete(movieId); err != nil {
		s.logger.Error("failed to delete movie", slog.String("cause", err.Error()))
		return err
	}

	return nil
}
