package middlewares

import (
	"github.com/afshintalebi/go-gin-starter/errutils"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func ErrorMiddleware(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	httpResponse := errutils.GerErrorData(500, "Internal server error", goErr.Error())

	c.AbortWithStatusJSON(500, httpResponse)
}
