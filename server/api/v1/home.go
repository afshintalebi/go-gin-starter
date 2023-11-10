package apiv1

import (
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.String(200, "API Version 1")
}
