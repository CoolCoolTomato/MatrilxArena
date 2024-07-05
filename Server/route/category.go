package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetCategoryRoute(r *gin.Engine) {
	category := r.Group("/category")
	category.Use(middleware.JWTAuthMiddleware())
	{
		category.POST("/GetCategoryList", api.GetCategoryList)
		category.POST("/GetCategory", api.GetCategory)
	}
	adminCategory := r.Group("/category")
	adminCategory.Use(middleware.AdminAuthMiddleware())
	{
		adminCategory.POST("/CreateCategory", api.CreateCategory)
		adminCategory.POST("/UpdateCategory", api.UpdateCategory)
		adminCategory.POST("/DeleteCategory", api.DeleteCategory)
	}
}
