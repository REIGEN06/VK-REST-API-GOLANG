package models

import (
	"gorm.io/gorm"
	"time"
)

type Actor struct {
	gorm.Model `swaggerignore:"true"`
	ID         int       `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	Sex        string    `gorm:"type:varchar(64)" json:"sex"`
	BirthDate  time.Time `json:"birth_date"`
	Movies     []Movie   `gorm:"many2many:actor_movies" json:"movies"`
}
