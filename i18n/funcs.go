package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetI18nMessage(messageId string, templateData map[string]interface{}) string {
	langLocalizer := langLocalizer

	return langLocalizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: templateData,
	})
}
