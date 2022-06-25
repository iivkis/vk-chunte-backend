package entities

import (
	"time"
)

type User struct {
	ID        *int      `db:"db" json:"id"`
	VkID      *uint     `db:"vk_id" json:"vk_id"`
	Name      *string   `db:"name" json:"name"  validate:"omitempty,min=1,max=100"`
	Age       *uint     `db:"age" json:"age" validate:"omitempty,min=14,max=100"`
	CreatedAt time.Time `json:"created_at"`
}
