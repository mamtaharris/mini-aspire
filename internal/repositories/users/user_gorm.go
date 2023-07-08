package users

import (
	"context"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"github.com/pkg/errors"
)

func (u *userRepo) Create(ctx context.Context, user entities.Users) (entities.Users, error) {

	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	result := u.writeDB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return user, errors.Wrap(result.Error, "[userRepo][Create]")
	}
	return user, nil
}

func (u *userRepo) GetByUsername(ctx context.Context, username string) (entities.Users, error) {

	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	result := u.readDB.WithContext(ctx).Where("username = ?", username).Find()
	if result.Error != nil {
		return user, errors.Wrap(result.Error, "[userRepo][Create]")
	}
	return user, nil
}
