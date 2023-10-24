package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	loadEnv()
}

func loadEnv(path ...string) {
	fileName := "./.env"
	if len(path) > 0 && path[0] != "" {
		fileName = path[0]
	}

	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	} else {
		fmt.Println(".env file has been loaded")
	}
}
