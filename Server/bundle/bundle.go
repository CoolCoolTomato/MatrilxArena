package bundle

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var Bundle *i18n.Bundle

func GetBundle() *i18n.Bundle {
	return Bundle
}
