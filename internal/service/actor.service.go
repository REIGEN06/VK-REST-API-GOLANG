package service

import (
	"github.com/reigen06/vk-rest-api/internal/models"
	"github.com/reigen06/vk-rest-api/internal/repository"
	"log/slog"
)

type ActorService struct {
	repository repository.Actor
	logger     *slog.Logger
}

func NewActorService(repository repository.Actor, logger *slog.Logger) *ActorService {
	return &ActorService{
		repository: repository,
		logger:     logger}
}

func (s *ActorService) Create(actor *models.Actor) error {
	s.logger.Debug("actor has been created", actor.Name)
	if err := s.repository.Create(actor); err != nil {
		s.logger.Error("failed to create actor", slog.String("cause", err.Error()))
		return err
	}

	return nil
}

func (s *ActorService) GetAllWithMovies() ([]models.Actor, error) {
	s.logger.Debug("films have been requested")
	actors, err := s.repository.GetAllWithMovies()
	if err != nil {
		s.logger.Error("failed to get actors with movies", slog.String("cause", err.Error()))
		return nil, err
	}

	return actors, nil
}

func (s *ActorService) Update(actorId int, actor *models.Actor) error {
	s.logger.Debug("actor has been updated", actor.Name)
	if err := s.repository.Update(actorId, actor); err != nil {
		s.logger.Error("failed to update actor", slog.String("cause", err.Error()))
		return err
	}

	return nil

}

func (s *ActorService) Delete(actorId int) error {
	s.logger.Debug("actor has been deleted by id:", actorId)
	if err := s.repository.Delete(actorId); err != nil {
		s.logger.Error("failed to delete actor", slog.Int("actor_id", actorId), slog.String("cause", err.Error()))
		return err
	}

	return nil
}
