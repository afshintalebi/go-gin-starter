package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	loadEnv()
	loadConfigs()
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

func loadConfigs() {
	var configuration Configurations

	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading config file, %s", err))
	}

	// Set undefined variables
	viper.SetDefault("app.name", "go-gin-starter")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Sprintf("Unable to decode into struct, %v", err))
	}
}
