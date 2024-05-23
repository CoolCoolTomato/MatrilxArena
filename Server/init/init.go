package init

import (
	"fmt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/config"
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	config.Config = viper.New()
	config.Config.SetConfigName("config")
	config.Config.SetConfigType("json")
	config.Config.AddConfigPath(".")
	config.Config.ReadInConfig()
	
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		config.GetConfig().GetString("Database.Username"),
		config.GetConfig().GetString("Database.Password"),
		config.GetConfig().GetString("Database.Host"),
		config.GetConfig().GetString("Database.Port"),
		config.GetConfig().GetString("Database.Database"),
		config.GetConfig().GetString("Database.Charset"),
		config.GetConfig().GetString("Database.Local"),
		)
	
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	conn.AutoMigrate(
		&model.User{},
		&model.Image{},
		&model.Challenge{},
		)
	
	database.Database = conn
	
	route.Route = gin.Default()
	route.SetUserRoute(route.Route)
	route.SetImageRoute(route.Route)
	route.SetChallengeRoute(route.Route)
}