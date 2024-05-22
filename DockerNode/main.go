package main

import (
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/config"
	_ "github.com/CoolCoolTomato/MatrilxArena/DockerNode/config"
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route.SetRoutes(router)

	port := config.GetConfig().GetString("Port")
	router.Run(":" + port)
}
