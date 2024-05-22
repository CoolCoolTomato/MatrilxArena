package database

import (
	"fmt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/config"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init()  {
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
		)
	
	Database = conn
}

func GetDatabase() *gorm.DB {
	return Database
}
