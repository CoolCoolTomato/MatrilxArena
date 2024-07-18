package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name            string
	Description     string
	GroupChallenges []GroupChallenge `gorm:"foreignKey:GroupID"`
	Users           []User           `gorm:"many2many:group_user"`
}

func GetGroupList() ([]Group, error) {
	var groupList []Group
	err := database.GetDatabase().Model(&Group{}).Find(&groupList).Error
	if err != nil {
		return nil, err
	}
	return groupList, nil
}

func (group *Group) GetGroup() error {
	return database.GetDatabase().Model(&Group{}).Where("ID = ?", group.ID).First(&group).Error
}

func (group *Group) CreateGroup() error {
	return database.GetDatabase().Model(&Group{}).Create(&group).Error
}

func (group *Group) UpdateGroup() error {
	return database.GetDatabase().Model(&Group{}).Where("ID = ?", group.ID).Updates(&group).Error
}

func (group *Group) DeleteGroup() error {
	return database.GetDatabase().Model(&Group{}).Where("ID = ?", group.ID).Delete(&group).Error
}