package utils

import "github.com/afshintalebi/go-gin-starter/config"

func GetEnvironmentText() string {
	if config.IsProductionMode() {
		return "PRODUCTION"
	} else {
		return "DEBUG"
	}
}
