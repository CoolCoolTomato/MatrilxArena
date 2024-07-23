package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name            string
	Description     string
	Public          bool
	GroupChallenges []GroupChallenge `gorm:"foreignKey:GroupID"`
	Users           []User           `gorm:"many2many:group_user"`
}

func GetGroupList(queryGroup Group) ([]Group, error) {
	query := database.GetDatabase().Model(&Group{})
	if queryGroup.Name != "" {
		query = query.Where("name LIKE ?", "%"+queryGroup.Name+"%")
	}

	var groupList []Group
	err := query.
		Preload("GroupChallenges").
		Preload("Users").
		Find(&groupList).
		Error
	if err != nil {
		return nil, err
	}

	return groupList, nil
}

func (group *Group) GetGroup() error {
	return database.GetDatabase().Model(&Group{}).
		Preload("GroupChallenges").
		Preload("Users.GroupChallenges", "group_id = ?", group.ID).
		Where("ID = ?", group.ID).
		First(&group).
		Error
}

func (group *Group) CreateGroup() error {
	return database.GetDatabase().Model(&Group{}).
		Create(&group).
		Error
}

func (group *Group) UpdateGroup() error {
	return database.GetDatabase().Model(&Group{}).
		Select("Name", "Description", "Public").
		Where("ID = ?", group.ID).
		Updates(&group).
		Error
}

func (group *Group) DeleteGroup() error {
	return database.GetDatabase().Model(&Group{}).
		Where("ID = ?", group.ID).
		Delete(&group).
		Error
}

func (group *Group) GetUserList(queryUser User) ([]User, error) {
	query := database.GetDatabase().Model(group)
	if queryUser.Username != "" {
		query = query.Where("users.username LIKE ?", "%"+queryUser.Username+"%")
	}

	var userList []User
	err := query.
		Preload("GroupChallenges").
		Association("Users").
		Find(&userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (group *Group) AddUser(user *User) error {
	return database.GetDatabase().Model(group).Association("Users").Append(user)
}

func (group *Group) DeleteUser(user *User) error {
	return database.GetDatabase().Model(group).Association("Users").Delete(user)
}
