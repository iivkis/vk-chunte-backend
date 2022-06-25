package repository

import (
	"github.com/iivkis/vk-chunte/config"
)

var repo *Repository

func init() {
	config.Load("./../../.env")
	repo = NewRespository()
}
