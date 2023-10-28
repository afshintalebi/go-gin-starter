package server

import (
	"net/http"

	"github.com/afshintalebi/go-gin-starter/i18n"
	"github.com/afshintalebi/go-gin-starter/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupServer() *http.Server {
	// force log's color
	gin.ForceConsoleColor()

	router := gin.Default()

	// use a custom recovery function to format the global errors
	router.Use(gin.CustomRecovery(middlewares.ErrorMiddleware))

	// setup middlewares
	router.Use(middlewares.LanguageMiddleware())

	// call some requirements
	router.Use(func(ctx *gin.Context) {
		i18n.SetupI18n(ctx.Request)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	return server
}
