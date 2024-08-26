package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetCTFRoute(r *gin.Engine) {
	adminCTF := r.Group("/ctf")
	adminCTF.Use(middleware.AdminAuthMiddleware())
	{
		adminCTF.POST("/GetCTFList", api.GetCTFList)
		adminCTF.POST("/GetCTF", api.GetCTF)
		adminCTF.POST("/CreateCTF", api.CreateCTF)
		adminCTF.POST("/UpdateCTF", api.UpdateCTF)
		adminCTF.POST("/DeleteCTF", api.DeleteCTF)
	}
}
