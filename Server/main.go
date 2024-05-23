package main

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/config"
	_ "github.com/CoolCoolTomato/MatrilxArena/Server/init"
	"github.com/CoolCoolTomato/MatrilxArena/Server/route"
)

func main() {
	r := route.GetRoute()
	port := config.GetConfig().GetString("Server.Port")
	r.Run(":" + port)
}
