package entities

import (
	"errors"
	"time"
)

type User struct {
	ID        *uint     `json:"id"`
	VkID      *uint     `json:"vk_id"`
	Name      *string   `json:"name"`
	Age       *uint     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

var _ entity = (*User)(nil)

func (e *User) Validation(allowNil bool) error {
	if validate(isNil(e.Name), allowNil, func() bool {
		return len(*e.Name) < 1 || len(*e.Name) > 200
	}) {
		return errors.New("invalid name: name length should been beetwen 1 <= len <= 200")
	}

	if validate(isNil(e.Age), allowNil, func() bool {
		return *e.Age < 1 || *e.Age > 100
	}) {
		return errors.New("invalid age: age should been beetwen 1 <= age <= 100")
	}

	return nil
}
