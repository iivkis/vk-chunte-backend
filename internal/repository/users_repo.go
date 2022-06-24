package repository

import (
	"context"
	"database/sql"

	"github.com/iivkis/vk-chunte/internal/entities"
)

type UsersRepo struct {
	*Store
}

func NewUsersRepo(store *Store) *UsersRepo {
	return &UsersRepo{store}
}

func (r *UsersRepo) scan(row *sql.Row) (*entities.User, error) {
	u := new(entities.User)
	err := row.Scan(&u.ID, &u.VkID, &u.Name, &u.Age)
	return u, err
}

func (r *UsersRepo) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := `INSERT INTO users (
		vk_id,
		name,
		age 
	) VALUES ($1, $2, $3) RETURNING *;`

	row := r.db.QueryRowContext(ctx, query, user.VkID, user.Name, user.Age)
	if err := row.Err(); err != nil {
		return nil, err
	}

	u, err := r.scan(row)
	return u, err
}

func (r *UsersRepo) Update(ctx context.Context, id uint, user *entities.User) (*entities.User, error) {
	query := `UPDATE users SET name = $2,  age = $3 WHERE id = $1 RETURNING *;`

	row := r.db.QueryRowContext(ctx, query, id,
		user.Name,
		user.Age,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	u, err := r.scan(row)
	return u, err
}
