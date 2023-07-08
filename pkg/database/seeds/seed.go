package seeds

import (
	"time"

	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

func Users(db *gorm.DB) error {
	var err error = nil
	err = db.Create(&entities.Users{
		Username:  "Mamta",
		Password:  "Mamta123",
		Role:      "USER",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		return err
	}

	err = db.Create(&entities.Users{
		Username:  "Aspire",
		Password:  "Aspire123",
		Role:      "ADMIN",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		return err
	}

	return nil
}
