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

	
	sentryDSN := config.GetSentryDSN()
	if sentryDSN == "" {
		panic("Sentry DSN is empty")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		// Debug:            true,
		// AttachStacktrace: true,
		// EnableTracing: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	fmt.Println("Sentry has been loaded")
}
