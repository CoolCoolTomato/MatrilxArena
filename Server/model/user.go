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

func GetUserList() ([]User, error) {
	var userList []User
	err := database.GetDatabase().Select("ID", "Username", "Email").Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (user *User) GetUser() error {
	return database.GetDatabase().Model(&User{}).Where("ID = ?", user.ID).Select("ID", "Username", "Email").First(&user).Error
}

func (user *User) CreateUser() error {
	return database.GetDatabase().Create(&user).Error
}

func (user *User) UpdateUser() error {
	return database.GetDatabase().Model(&User{}).Where("ID = ?", user.ID).Updates(&user).Error
}

func (user *User) DeleteUser() error {
	return database.GetDatabase().Model(&User{}).Where("ID = ?", user.ID).Delete(&user).Error
}
