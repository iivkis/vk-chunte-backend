package entities

import (
	"testing"

	"github.com/iivkis/vk-chunte/internal/util"
	"github.com/stretchr/testify/require"
)

func createUser() *User {
	var (
		vkID uint   = uint(util.GetRandomInt(1, 99999))
		age  uint   = uint(util.GetRandomInt(1, 100))
		name string = util.GetRandomString(util.GetRandomInt(1, 200))
	)

	return &User{
		VkID: &vkID,
		Name: &name,
		Age:  &age,
	}
}

func TestUserValidation(t *testing.T) {
	var user1 = createUser()
	err1 := user1.Validation(true)
	require.NoError(t, err1)

	var user2 = createUser()
	err2 := user2.Validation(false)
	require.NoError(t, err2)

	var user3 = createUser()
	user3.Age = nil
	err3 := user3.Validation(false)
	require.Error(t, err3)

	var user4 = createUser()
	user4.Age = nil
	err4 := user4.Validation(true)
	require.NoError(t, err4)
}
