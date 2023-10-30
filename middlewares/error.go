package middlewares

import (
	"github.com/afshintalebi/go-gin-starter/errutils"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func ErrorMiddleware(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	httpResponse := errutils.HttpResponse{Message: "Internal server error", Status: 500, Description: goErr.Error()}

	c.AbortWithStatusJSON(500, httpResponse)
}
