package users

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

type userRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewRepo(writeDB *gorm.DB, readDB *gorm.DB) UserRepo {
	return &userRepo{writeDB: writeDB, readDB: readDB}
}

type UserRepo interface {
	Create(ctx context.Context, user entities.Users) (entities.Users, error)
	GetByUsername(ctx context.Context, username string) (entities.Users, error)
}
