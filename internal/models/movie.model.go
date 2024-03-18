package models

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model  `swaggerignore:"true"`
	ID          int       `gorm:"primaryKey;autoIncrement" `
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Rating      int       `gorm:"check:(rating >=1 and rating <= 10)" json:"rating"`
	ReleaseDate time.Time `json:"release-date"`
	Actors      []Actor   `gorm:"many2many:actor_movies" json:"actors"`
}
