package repositories

import (
	"context"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

type userRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewUserRepo(writeDB *gorm.DB, readDB *gorm.DB) UserRepo {
	return &userRepo{writeDB: writeDB, readDB: readDB}
}

type UserRepo interface {
	Create(ctx context.Context, user entities.Users) (entities.Users, error)
	GetByUsername(ctx context.Context, username string) (entities.Users, error)
}

func (u *userRepo) Create(ctx context.Context, user entities.Users) (entities.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()
	result := u.writeDB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *userRepo) GetByUsername(ctx context.Context, username string) (entities.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()
	var user entities.Users
	result := u.readDB.WithContext(ctx).Where("username = ?", username).Take(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
