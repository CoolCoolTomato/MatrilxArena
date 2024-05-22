package main

import (
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route.SetRoutes(router)
	router.Run()
}
