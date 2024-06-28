package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Password   string
	Email      string
	Role       int
	Challenges []Challenge `gorm:"many2many:challenge_user"`
}

func GetUserList() ([]User, error) {
	var userList []User
	err := database.GetDatabase().Model(&User{}).Select("ID", "Username", "Email", "Role").Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (user *User) GetUser() error {
	return database.GetDatabase().Model(&User{}).Where("ID = ?", user.ID).Select("ID", "Username", "Email").First(&user).Error
}

func (user *User) CreateUser() error {
	return database.GetDatabase().Model(&User{}).Create(&user).Error
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

func (user *User) GetChallengeList() ([]Challenge, error) {
	var challengeList []Challenge
	err := database.GetDatabase().Model(user).Association("Challenges").Find(&challengeList)
	if err != nil {
		return nil, err
	}
	return challengeList, nil
}

func (user *User) AddChallenge(challenge *Challenge) error {
	return database.GetDatabase().Model(user).Association("Challenges").Append(challenge)
}

func (user *User) DeleteChallenge(challenge *Challenge) error {
	return database.GetDatabase().Model(user).Association("Challenges").Delete(challenge)
}
