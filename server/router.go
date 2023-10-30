package server

import (
	apiv1 "github.com/afshintalebi/go-gin-starter/server/api/v1"
	"github.com/gin-gonic/gin"
)

func setupV1Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("", apiv1.HomeHandler)
		v1.GET("/health-check", apiv1.HealthCheckHandler)
	}
}

