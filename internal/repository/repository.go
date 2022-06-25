package repository

import (
	"context"

	"github.com/iivkis/vk-chunte/internal/entities"
)

type Users interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, id int, user *entities.User) (*entities.User, error)
	GetByID(ctx context.Context, id int) (*entities.User, error)
	// GetByVkID(ctx context.Context, vkID int) (*entities.User, error)
	// Delete(ctx context.Context, id int) (*entities.User, error)
}

type Repository struct {
	Users Users
}

func NewRespository() *Repository {
	store := newStore()

	return &Repository{
		Users: NewUsersRepo(store),
	}
}
