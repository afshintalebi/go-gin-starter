package server

import (
	"net/http"

	"github.com/afshintalebi/go-gin-starter/i18n"
	"github.com/afshintalebi/go-gin-starter/middlewares"
	ratelimiter "github.com/afshintalebi/go-gin-starter/rate-limiter"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupServer() *http.Server {
	// force log's color
	gin.ForceConsoleColor()

	router := gin.Default()

	// setup middlewares
	router.Use(gin.CustomRecovery(middlewares.ErrorMiddleware))
	router.Use(middlewares.LanguageMiddleware())
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	router.Use(cors.Default())
	router.Use(ratelimiter.GetMiddlewareHandler())

	// call some requirements
	router.Use(func(c *gin.Context) {
		i18n.SetupI18n(c.Request)
	})

	setupV1Routes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	return server
}
