package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string
	Password	     string
	Email            string
}

func (user *User) GetUser() error {
	return database.GetDatabase().First(&user).Error
}

func (user *User) CreateUser() error {
	return database.GetDatabase().Create(&user).Error
}

func (user *User) UpdateUser() error {
	return database.GetDatabase().Updates(&user).Error
}

func (user *User) DeleteUser() error {
	return database.GetDatabase().Delete(&user).Error
}
