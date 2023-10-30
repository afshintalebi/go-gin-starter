package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(ctx *gin.Context) {
	if serverIsHealthy() {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server is healthy",
		})
	} else {
		panic("Server is not healthy")
	}
}

func serverIsHealthy() bool {
	// Check the health of the server and return true or false accordingly
	// For example, check if the server can connect to the database
	return true
}
