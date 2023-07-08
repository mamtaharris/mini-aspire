package database

import (
	"fmt"

	"github.com/mamtaharris/mini-aspire/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%d sslmode=disable TimeZone=UTC",
		config.DB.Host, config.DB.Username, config.DB.Password, config.DB.Name, config.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
