package init

import (
	"encoding/json"
	"github.com/CoolCoolTomato/MatrilxArena/Server/bundle"
	"github.com/CoolCoolTomato/MatrilxArena/Server/locales"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"strings"
)

func BundleInit() {
	bundle.Bundle = i18n.NewBundle(language.English)
	bundle.Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	localeFiles, err := locales.FS.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range localeFiles {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			data, err := locales.FS.ReadFile(file.Name())
			if err != nil {
				panic(err)
			}
			_, err = bundle.Bundle.ParseMessageFileBytes(data, file.Name())
			if err != nil {
				panic(err)
			}
		}
	}
}
