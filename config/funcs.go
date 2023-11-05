package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	return viper.GetString(key)
}

func GetAppEnv() string {
	var configuration Configurations
	mode := GetEnv("GIN_MODE")
	if mode == "" {
		mode = configuration.App.Env
	}

	return mode
}

func GetSentryDSN() string {
	var configuration Configurations
	sentryDsn := GetEnv("SENTRY_DSN")
	fmt.Println(fmt.Printf("<<< sentryDsn is %s >>>", sentryDsn))

	if sentryDsn == "" {
		sentryDsn = configuration.Sentry.DSN
	}

	return sentryDsn
}

func IsProductionMode() bool {
	devMode := GetAppEnv()

	return devMode == "release"
}

func SetToReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}
