package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type CTFTeam struct {
	gorm.Model
	Name        string
	Description string
	CTFID       uint
	CTF         CTF    `gorm:"foreignKey:CTFID;constraint:OnDelete:SET NULL"`
	Users       []User `gorm:"many2many:ctf_team_user"`
}

func GetCTFTeamList(queryCTFTeam CTFTeam) ([]CTFTeam, error) {
	query := database.GetDatabase().Model(&CTFTeam{})
	if queryCTFTeam.Name != "" {
		query = query.Where("name LIKE ?", "%"+queryCTFTeam.Name+"%")
	}

	var ctfTeamList []CTFTeam
	err := query.
		Preload("Users").
		Find(&ctfTeamList).
		Error
	if err != nil {
		return nil, err
	}

	return ctfTeamList, nil
}

func (ctfTeam *CTFTeam) GetCTFTeam() error {
	return database.GetDatabase().Model(&CTFTeam{}).
		Preload("Users").
		Where("ID = ?", ctfTeam.ID).
		First(&ctfTeam).
		Error
}

func (ctfTeam *CTFTeam) CreateCTFTeam() error {
	return database.GetDatabase().Model(&CTFTeam{}).
		Create(&ctfTeam).
		Error
}

func (ctfTeam *CTFTeam) UpdateCTFTeam() error {
	return database.GetDatabase().Model(&CTFTeam{}).
		Select("Name", "Description").
		Where("ID = ?", ctfTeam.ID).
		Updates(&ctfTeam).
		Error
}

func (ctfTeam *CTFTeam) DeleteCTFTeam() error {
	return database.GetDatabase().Model(&CTFTeam{}).
		Where("ID = ?", ctfTeam.ID).
		Delete(&ctfTeam).
		Error
}

func (ctfTeam *CTFTeam) GetUserList(queryUser User) ([]User, error) {
	query := database.GetDatabase().Model(ctfTeam)
	if queryUser.Username != "" {
		query = query.Where("users.username LIKE ?", "%"+queryUser.Username+"%")
	}

	var userList []User
	err := query.
		Association("Users").
		Find(&userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (ctfTeam *CTFTeam) AddUser(user *User) error {
	return database.GetDatabase().Model(ctfTeam).Association("Users").Append(user)
}

func (ctfTeam *CTFTeam) DeleteUser(user *User) error {
	return database.GetDatabase().Model(ctfTeam).Association("Users").Delete(user)
}
