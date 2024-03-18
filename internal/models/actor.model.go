package models

import (
	"gorm.io/gorm"
	"time"
)

type Actor struct {
	gorm.Model `swaggerignore:"true"`
	ID         int    `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(255)"`
	Sex        string `gorm:"type:varchar(64)"`
	BirthDate  time.Time
	Movies     []Movie `gorm:"many2many:actor_movies" json:"movies,omitempty"`
}
