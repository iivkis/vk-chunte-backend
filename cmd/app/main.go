package main

import (
	"github.com/iivkis/vk-chunte/config"
	"github.com/iivkis/vk-chunte/internal/app"
)

func main() {
	config.Load(".env")
	app.Launch()
}
