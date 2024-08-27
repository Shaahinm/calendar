package lang

import (
	"encoding/json"
	"log/slog"
	"path"

	"github.com/Shaahinm/calendar/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
)

var supportedLanguages = map[string]string{
	"en": "en",
	"de": "de",
}

func getBundle() *i18n.Bundle {
	if bundle == nil {
		Init()
	}
	return bundle
}

func getLocalizer() *i18n.Localizer {
	if localizer == nil {
		if bundle == nil {
			Init()
		}
		envLang := config.Envs.UserLang
		if lang, ok := supportedLanguages[envLang]; ok {
			envLang = lang
		} else {
			envLang = "en"
		}
		localizer = i18n.NewLocalizer(bundle, envLang)
	}
	return localizer
}

func Init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, lang := range supportedLanguages {
		bundle.MustLoadMessageFile(path.Join("internal", "locale", "lang", lang+".json"))
	}
}

func GetWithTemplate(id string, args map[string]interface{}) string {
	loc := getLocalizer()
	personUnreadEmails := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
		TemplateData: args,
	})

	return personUnreadEmails
}

func Get(id string) string {
	loc := getLocalizer()
	str, err := loc.LocalizeMessage(&i18n.Message{ID: id})
	if err != nil {
		slog.Warn("error while getting lang string", "err", err)
		// str = id
	}
	if str == " " {
		str = ""
	}
	return str
}
