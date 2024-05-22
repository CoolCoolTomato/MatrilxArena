package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Challenge struct {
	gorm.Model
	Image 			Image
	Title 			string
	Description 	string
}

func (challenge *Challenge) GetChallenge() error {
	return database.GetDatabase().First(&challenge).Error
}

func (challenge *Challenge) CreateChallenge() error {
	return database.GetDatabase().Create(&challenge).Error
}

func (challenge *Challenge) UpdateChallenge() error {
	return database.GetDatabase().Updates(&challenge).Error
}

func (challenge *Challenge) DeleteChallenge() error {
	return database.GetDatabase().Delete(&challenge).Error
}
