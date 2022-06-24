package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_NAME string

	HOST string
	PORT string

	DB_URL string
)

func Load(dir string) {
	if err := godotenv.Load(dir); err != nil {
		panic(err)
	}
	set()
}

func set() {
	//meta
	APP_NAME = os.Getenv("APP_NAME")

	//server settings
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")

	//db
	DB_URL = os.Getenv("DB_URL")
}
