package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetAttachmentRoute(r *gin.Engine) {
	attachment := r.Group("/attachment")
	attachment.Use(middleware.JWTAuthMiddleware())
	{
		attachment.POST("GetAttachmentList", api.GetAttachmentList)
		attachment.POST("GetAttachment", api.GetAttachment)
	}
	adminAttachment := r.Group("/attachment")
	adminAttachment.Use(middleware.AdminAuthMiddleware())
	{
		adminAttachment.POST("CreateAttachment", api.CreateAttachment)
		adminAttachment.POST("UpdateAttachment", api.UpdateAttachment)
		adminAttachment.POST("DeleteAttachment", api.DeleteAttachment)
		adminAttachment.POST("UploadAttachment", api.UploadAttachment)
	}
}
