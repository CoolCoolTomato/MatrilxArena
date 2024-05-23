package route

import "github.com/gin-gonic/gin"

var Route *gin.Engine

func GetRoute() *gin.Engine {
	return Route
}
