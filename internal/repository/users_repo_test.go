package repository

import (
	"context"
	"testing"

	"github.com/iivkis/vk-chunte/internal/entities"
	"github.com/iivkis/vk-chunte/internal/util"
	"github.com/stretchr/testify/require"
)

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
	current := createUser(t)

	var (
		newName string = util.GetRandomString(util.GetRandomInt(1, 100))
		newAge  uint   = uint(util.GetRandomInt(14, 100))
	)

	updatedFields := &entities.User{
		Name: &newName,
		Age:  &newAge,
	}

	err := util.ValidateStruct(updatedFields)
	require.NoError(t, err)

	updatedUser, err := repo.Users.Update(context.Background(), *current.ID, updatedFields)

	require.NoError(t, err)
	require.Equal(t, updatedUser.ID, current.ID)
	require.Equal(t, updatedUser.VkID, current.VkID)
	require.Equal(t, updatedUser.Name, updatedFields.Name)
	require.Equal(t, updatedUser.Age, updatedFields.Age)
	require.Equal(t, updatedUser.CreatedAt, current.CreatedAt)
}

func TestGetUserByID(t *testing.T) {
	current := createUser(t)
	found, err := repo.Users.GetByID(context.Background(), *current.ID)

	require.NoError(t, err)
	require.Equal(t, current, found)
}
