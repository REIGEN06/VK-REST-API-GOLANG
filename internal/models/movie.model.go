package models

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model  `swaggerignore:"true"`
	ID          int    `gorm:"primaryKey;autoIncrement" `
	Name        string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text" `
	Rating      int    `gorm:"check:(rating >=1 and rating <= 10)"`
	ReleaseDate time.Time
	Actors      []Actor `gorm:"many2many:actor_movies"`
}
