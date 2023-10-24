package i18n

import (
	"encoding/json"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var langLocalizer *i18n.Localizer

// TODO optimize or modifying of loading the I18n
func SetupI18n(req *http.Request) {
	// create bundle with default language
	bundle := i18n.NewBundle(language.English)

	// setup to load JSON files for translations
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("./assets/i18n/en.json")

	// get language from the header, if failed the default language will load
	accept := req.Header.Get("Accept-Language")

	// create localizer
	langLocalizer = i18n.NewLocalizer(bundle, accept)
}
