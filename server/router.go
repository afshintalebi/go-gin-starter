package server

import (
	api "github.com/afshintalebi/go-gin-starter/server/api"
	apiv1 "github.com/afshintalebi/go-gin-starter/server/api/v1"
	"github.com/gin-gonic/gin"
)

func setupV1Routes(router *gin.Engine) {

	errors := router.Group("/errors-example")
	{
		errors.GET("400", api.Error400Handler)
		errors.GET("404", api.Error404Handler)
		errors.GET("500", api.Error500Handler)
	}

	v1 := router.Group("/v1")
	{
		v1.GET("", apiv1.HomeHandler)

		v1.GET("/health-check", apiv1.HealthCheckHandler)
	}
}
