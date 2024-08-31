package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetCTFUserRoute(r *gin.Engine) {
	adminCTFUser := r.Group("/ctfUser")
	adminCTFUser.Use(middleware.AdminAuthMiddleware())
	{
		adminCTFUser.POST("/GetCTFUserList", api.GetCTFUserList)
		adminCTFUser.POST("/AddCTFUser", api.AddCTFUser)
		adminCTFUser.POST("/RemoveCTFUser", api.RemoveCTFUser)
	}
}
