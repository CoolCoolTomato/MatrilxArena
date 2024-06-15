package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Role     int
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

func (user *User) GetUserByUsernameOrEmail() error {
	if user.Username != "" {
		return database.GetDatabase().Model(&User{}).Where("Username = ?", user.Username).First(&user).Error
	}
	if user.Email != "" {
		return database.GetDatabase().Model(&User{}).Where("Email = ?", user.Email).First(&user).Error
	}
	return gorm.ErrRecordNotFound
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
