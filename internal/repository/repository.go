package repository

import (
	"context"

	"github.com/iivkis/vk-chunte/internal/entities"
)

type Users interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, id uint, user *entities.User) (*entities.User, error)
}

type Repository struct {
	Users Users
}

func NewRespository() *Repository {
	store := NewStore()
	return &Repository{
		Users: NewUsersRepo(store),
	}
}
