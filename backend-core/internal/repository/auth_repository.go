package repository

import (
	"context"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db gorm.DB
}

func NewAuthRepository(db gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r AuthRepository) GetUserById(ctx context.Context, userId string) (model.User, error) {
	user, err := gorm.G[model.User](&r.db).Where("id = ?", userId).First(ctx)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
