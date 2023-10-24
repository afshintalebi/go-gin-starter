package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvDevMode() string {
	return GetEnv("GIN_MODE")
}

func GetEnvSentryDSN() string {
	return GetEnv("SENTRY_DSN")
}

func IsProductionMode() bool {
	devMode := GetEnvDevMode()

	return devMode == "release"
}

func SetToReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}
