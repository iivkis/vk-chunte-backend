package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/iivkis/vk-chunte/internal/entities"
	"github.com/iivkis/vk-chunte/internal/util"
)

type UsersRepo struct {
	store *store
}

func NewUsersRepo(store *store) Users {
	return &UsersRepo{store: store}
}

func (r *UsersRepo) scan(row *sql.Row) (user *entities.User, err error) {
	err = row.Scan(
		&user.ID,
		&user.VkID,
		&user.Name,
		&user.Age,
	)
	return
}

func (r *UsersRepo) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := `INSERT INTO users (
		vk_id,
		name,
		age 
	) VALUES ($1, $2, $3) RETURNING *;`

	row := r.store.db.QueryRowContext(ctx, query, user.VkID, user.Name, user.Age)
	if err := row.Err(); err != nil {
		return nil, err
	}

	return r.scan(row)
}

func (r *UsersRepo) Update(ctx context.Context, id int, user *entities.User) (*entities.User, error) {
	fields, values, err := util.StructToInsertableSQLParams(user)
	if err != nil {
		return nil, err
	}

	query := "UPDATE users SET " + fields + " WHERE id = " + strconv.Itoa(id) + " RETURNING *;"

	row := r.store.db.QueryRowContext(ctx, query, values...)
	if err := row.Err(); err != nil {
		return nil, err
	}

	return r.scan(row)
}
