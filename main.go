package main

import (
	"time"

	"github.com/afshintalebi/go-gin-starter/config"
	"github.com/afshintalebi/go-gin-starter/logger"
	_ "github.com/afshintalebi/go-gin-starter/sentry"
	"github.com/afshintalebi/go-gin-starter/server"
	"github.com/getsentry/sentry-go"
)

func init() {
	// check the configironment to set the App on the release mode
	if config.IsProductionMode() {
		config.SetToReleaseMode()
	}
}

func main() {
	defer logger.GetLogger().Sync()
	defer sentry.Flush(2 * time.Second)

	server := server.SetupServer()

	server.ListenAndServe()
}
