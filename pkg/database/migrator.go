package database

import (
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var users entities.Users

	usersM := Migrate{TableName: "users", Run: func(db *gorm.DB) error { return db.AutoMigrate(&users) }}

	return []Migrate{
		usersM,
	}
}
