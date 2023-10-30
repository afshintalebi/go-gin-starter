package apiv1

import "github.com/gin-gonic/gin"

func HomeHandler(ctx *gin.Context) {
	ctx.String(200, "API Version 1")
}
