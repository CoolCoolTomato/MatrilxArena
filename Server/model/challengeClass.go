package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type ChallengeClass struct {
	gorm.Model
	Name string
}

func GetChallengeClassList() ([]ChallengeClass, error) {
	var challengeClassList []ChallengeClass
	err := database.GetDatabase().Model(&ChallengeClass{}).Find(&challengeClassList).Error
	if err != nil {
		return nil, err
	}
	return challengeClassList, nil
}

func (challengeClass *ChallengeClass) GetChallengeClass() error {
	return database.GetDatabase().Model(&ChallengeClass{}).Where("ID = ?", challengeClass.ID).First(&challengeClass).Error
}

func (challengeClass *ChallengeClass) CreateChallengeClass() error {
	return database.GetDatabase().Model(&ChallengeClass{}).Create(&challengeClass).Error
}

func (challengeClass *ChallengeClass) UpdateChallengeClass() error {
	return database.GetDatabase().Model(&ChallengeClass{}).Where("ID = ?", challengeClass.ID).Updates(&challengeClass).Error
}

func (challengeClass *ChallengeClass) DeleteChallengeClass() error {
	return database.GetDatabase().Model(&ChallengeClass{}).Where("ID = ?", challengeClass.ID).Delete(&challengeClass).Error
}
