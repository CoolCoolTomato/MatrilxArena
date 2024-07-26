package init

import (
	"fmt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/config"
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInit() {
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
		&model.DockerNode{},
	)

	database.Database = conn

    username := config.GetConfig().GetString("Database.Admin.Username")
    password := config.GetConfig().GetString("Database.Admin.Password")
    user := model.User{
        Username: username,
        Password: password,
        Role: 1,
    }
    err = user.GetUserByUsernameOrEmail()
    if err != nil {
        err = user.SetPassword(user.Password)
        err = user.CreateUser()
        if err != nil {
            panic(err)
        }
    }

	database.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GetConfig().GetString("Redis.Host"), config.GetConfig().GetString("Redis.Port")),
		Password: config.GetConfig().GetString("Redis.Password"),
		DB:       config.GetConfig().GetInt("Redis.DB"),
	})
}
