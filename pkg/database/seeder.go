package database

import (
	"github.com/mamtaharris/mini-aspire/pkg/database/seeds"
	"gorm.io/gorm"
)

type Seed struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AllSeeds(db *gorm.DB) []Seed {
	users := Seed{TableName: "users", Run: func(db *gorm.DB) error { return seeds.Users(db) }}
	return []Seed{
		users,
	}
}
