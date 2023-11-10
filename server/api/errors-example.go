package api

import (
	"github.com/afshintalebi/go-gin-starter/errutils"
	"github.com/afshintalebi/go-gin-starter/i18n"
	"github.com/gin-gonic/gin"
)

func Error400Handler(c *gin.Context) {

	httpResponse := errutils.GerErrorData(
		400,
		i18n.GetI18nMessage("Error400Message", nil),
		i18n.GetI18nMessage("ErrorDescription", nil),
	)

	c.AbortWithStatusJSON(400, httpResponse)
}

func Error404Handler(c *gin.Context) {
	httpResponse := errutils.GerErrorData(
		404,
		i18n.GetI18nMessage("Error404Message", nil),
		i18n.GetI18nMessage("ErrorDescription", nil),
	)

	c.AbortWithStatusJSON(404, httpResponse)
}

func Error500Handler(c *gin.Context) {
	httpResponse := errutils.GerErrorData(
		500,
		i18n.GetI18nMessage("Error500Message", nil),
		i18n.GetI18nMessage("ErrorDescription", nil),
	)

	c.AbortWithStatusJSON(500, httpResponse)
}
