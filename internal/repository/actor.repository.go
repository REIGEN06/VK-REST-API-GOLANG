package repository

import (
	"github.com/reigen06/vk-rest-api/internal/models"
	"gorm.io/gorm"
	"log/slog"
)

type ActorPostgresRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewActorPostgresRepository(db *gorm.DB, logger *slog.Logger) *ActorPostgresRepository {
	return &ActorPostgresRepository{
		db:     db,
		logger: logger}
}

func (r *ActorPostgresRepository) Create(actor *models.Actor) error {
	if err := r.db.Create(actor).Error; err != nil {
		r.logger.Error("error while creating actor", err)
		return err
	}
	return nil
}

func (r *ActorPostgresRepository) Update(actorId int, actor *models.Actor) error {
	if err := r.db.Model(&models.Actor{}).Where("id = ?", actorId).Updates(actor).Error; err != nil {
		r.logger.Error("error while updating actor", err)
		return err
	}
	return nil
}

func (r *ActorPostgresRepository) Delete(actorId int) error {
	if err := r.db.Delete(&models.Actor{}, actorId).Error; err != nil {
		r.logger.Error("error while deleting actor", err)
		return err
	}
	return nil
}

func (r *ActorPostgresRepository) GetAllWithMovies() ([]models.Actor, error) {
	var actors []models.Actor
	if err := r.db.Preload("Movies").Find(&actors).Error; err != nil {
		r.logger.Error("error while fetching actors with movies", err)
		return nil, err
	}
	return actors, nil
}
