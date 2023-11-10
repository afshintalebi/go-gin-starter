package middlewares

import (
	"errors"
	"slices"

	"github.com/gin-gonic/gin"
)

func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedLanguages := []string{"en"}

		language := c.Request.Header.Get("Accept-Language")

		if language == "" {
			language = allowedLanguages[0]
			c.Request.Header.Add("Accept-Language", language)
		} else if language != "" && !slices.Contains(allowedLanguages, language) {
			panic(errors.New("language value in header is invalid"))
		}

		// save language to the context
		c.Set("language", language)

		// before request

		c.Next()

		// after request
	}
}
