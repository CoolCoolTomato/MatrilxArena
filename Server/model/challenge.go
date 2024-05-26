package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Challenge struct {
	gorm.Model
	ImageID 		uint
	Image 			Image
	Title 			string
	Description 	string
}

func GetChallengeList() ([]Challenge, error) {
	var challengeList []Challenge
	err := database.GetDatabase().Find(&challengeList).Error
	if err != nil {
		return nil, err
	}
	return challengeList, nil
}

func (challenge *Challenge) GetChallenge() error {
	return database.GetDatabase().Model(&Challenge{}).Where("ID = ?", challenge.ID).First(&challenge).Error
}

func (challenge *Challenge) CreateChallenge() error {
	return database.GetDatabase().Create(&challenge).Error
}

func (challenge *Challenge) UpdateChallenge() error {
	return database.GetDatabase().Model(&Challenge{}).Where("ID = ?", challenge.ID).Updates(&challenge).Error
}

func (challenge *Challenge) DeleteChallenge() error {
	return database.GetDatabase().Model(&Challenge{}).Where("ID = ?", challenge.ID).Delete(&challenge).Error
}
