package repository

import (
	"context"
	"testing"

	"github.com/iivkis/vk-chunte/config"
	"github.com/iivkis/vk-chunte/internal/entities"
	"github.com/iivkis/vk-chunte/internal/util"
	"github.com/stretchr/testify/require"
)

var repo *Repository

func init() {
	config.Load("./../../.env")
	repo = NewRespository()
}

func createUser(t *testing.T) *entities.User {
	var (
		VkID uint   = uint(util.GetRandomInt(1, 1_000_000))
		Name string = util.GetRandomString(util.GetRandomInt(10, 200))
		Age  uint   = uint(util.GetRandomInt(1, 100))
	)

	entity := &entities.User{
		VkID: &VkID,
		Name: &Name,
		Age:  &Age,
	}

	err := entity.Validation(false)
	require.NoError(t, err)

	user, err := repo.Users.Create(context.Background(), entity)
	require.NoError(t, err)
	require.Equal(t, user.VkID, entity.VkID)
	require.Equal(t, user.Name, entity.Name)
	require.Equal(t, user.Age, entity.Age)

	return user
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestUpdateUser(t *testing.T) {
	user := createUser(t)

	newAge := uint(54)

	updatedUserFields := &entities.User{
		Age: &newAge,
	}

	user1, err := repo.Users.Update(context.Background(), *user.ID, updatedUserFields)

	require.NoError(t, err)
	require.Equal(t, *user1.ID, *user.ID)
	require.Equal(t, *user1.Name, *user.Name)
	require.Equal(t, *user1.Age, *updatedUserFields.Age)
}
