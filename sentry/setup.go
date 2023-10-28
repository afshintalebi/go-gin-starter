package sentry

import (
	"fmt"
	"log"

	"github.com/afshintalebi/go-gin-starter/config"
	"github.com/getsentry/sentry-go"
)

func init() {
	if !config.IsProductionMode() {
		fmt.Println("Sentry doesn't work on the DEBUG mode")
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.GetEnvSentryDSN(),
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	fmt.Println("Sentry has been loaded")
}