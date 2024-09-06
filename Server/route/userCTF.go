package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserCTFRoute(r *gin.Engine) {
	userCTF := r.Group("/userCTF")
	userCTF.Use(middleware.JWTAuthMiddleware())
	{
		userCTF.POST("/GetUserCTFList", api.GetUserCTFList)
		userCTF.POST("/GetVisibleUserCTFList", api.GetVisibleUserCTFList)
		userCTF.POST("/AddUserCTF", api.AddUserCTF)
		userCTF.POST("/RemoveUserCTF", api.RemoveUserCTF)
	}
}
