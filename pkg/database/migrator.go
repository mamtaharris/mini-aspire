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
	var loans entities.Loans
	var repayments entities.Repayments

	usersM := Migrate{TableName: "users", Run: func(db *gorm.DB) error { return db.AutoMigrate(&users) }}
	loansM := Migrate{TableName: "loans", Run: func(db *gorm.DB) error { return db.AutoMigrate(&loans) }}
	repaymentsM := Migrate{TableName: "repayments", Run: func(db *gorm.DB) error { return db.AutoMigrate(&repayments) }}

	return []Migrate{
		usersM,
		loansM,
		repaymentsM,
	}
}
