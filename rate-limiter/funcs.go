package ratelimiter

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/afshintalebi/go-gin-starter/errutils"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.AbortWithStatusJSON(
		429,
		errutils.GerErrorData(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String(), "Rate Limit Error"),
	)
}

func GetMiddlewareHandler() gin.HandlerFunc {
	return rateLimitMiddleware
}
