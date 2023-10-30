package sentry

import "github.com/getsentry/sentry-go"

func CaptureError(err error) {
	sentry.CaptureException(err)
}
