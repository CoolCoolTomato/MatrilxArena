package localizer

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetMessage(messageID string, c *gin.Context) string {
	localizer := c.MustGet("localizer").(*i18n.Localizer)
	message, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	return message
}
