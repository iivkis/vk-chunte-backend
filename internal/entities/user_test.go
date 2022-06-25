package entities

import (
	"testing"
	"time"

	"github.com/iivkis/vk-chunte/internal/util"
	"github.com/stretchr/testify/require"
)

type userTest struct {
	ID        int
	VkID      uint
	Name      string
	Age       uint
	CreatedAt time.Time
}

func TestUserValidateCorrect(t *testing.T) {
	users := make([]userTest, 5)

	for i := range users {
		users[i] = userTest{
			ID:   util.GetRandomInt(1, 1000),
			VkID: uint(util.GetRandomInt(1, 1000)),
			Name: util.GetRandomString(util.GetRandomInt(1, 100)),
			Age:  uint(util.GetRandomInt(14, 100)),
		}
	}

	for _, user := range users {
		u := User{
			ID:   &user.ID,
			VkID: &user.VkID,
			Name: &user.Name,
			Age:  &user.Age,
		}

		err := util.ValidateStruct(u)
		require.NoError(t, err)
	}
}

func TestUserValidateOnNilFields(t *testing.T) {
	var (
		ID   = util.GetRandomInt(1, 1000)
		VkID = uint(util.GetRandomInt(1, 1000))
	)

	user := User{
		ID:   &ID,
		VkID: &VkID,
	}

	err := util.ValidateStruct(user)
	require.NoError(t, err)
}

func TestUserValidateIncorrect(t *testing.T) {
	users := []userTest{
		{
			ID:   util.GetRandomInt(1, 1000),
			VkID: uint(util.GetRandomInt(1, 1000)),
			Name: util.GetRandomString(util.GetRandomInt(101, 1000)), //incorrect
			Age:  uint(util.GetRandomInt(101, 1000)),                 //incorrect
		},

		{
			ID:   util.GetRandomInt(1, 1000),
			VkID: uint(util.GetRandomInt(1, 1000)),
			Name: "", //incorrect
			Age:  0,  //incorrect
		},
	}

	for _, user := range users {
		u := User{
			ID:   &user.ID,
			VkID: &user.VkID,
			Name: &user.Name,
			Age:  &user.Age,
		}

		err := util.ValidateStruct(u)
		require.Error(t, err)
	}
}
