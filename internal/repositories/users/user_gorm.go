package users

import (
	"context"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
)

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
