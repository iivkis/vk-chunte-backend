package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/iivkis/vk-chunte/internal/entities"
	"github.com/iivkis/vk-chunte/internal/util"
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

func (r *UsersRepo) Update(ctx context.Context, id int, user *entities.User) (*entities.User, error) {
	fields, values, err := util.StructToSQLParams(user)
	if err != nil {
		return nil, err
	}

	query := "UPDATE users SET " + fields + " WHERE id = " + strconv.Itoa(id) + " RETURNING *;"
	row := r.db.QueryRowContext(ctx, query, values...)

	if err := row.Err(); err != nil {
		return nil, err
	}

	u, err := r.scan(row)
	return u, err
}
