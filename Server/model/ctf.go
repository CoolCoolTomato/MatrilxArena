package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
	"time"
)

type CTF struct {
	gorm.Model
	Name          string
	Description   string
	Public        bool
	TeamSignIn    bool
	CTFChallenges []CTFChallenge `gorm:"foreignKey:CTFID"`
	Users         []User         `gorm:"many2many:ctf_user"`
	CTFTeams      []CTFTeam      `gorm:"foreignKey:CTFID"`
	StartTime     time.Time
	EndTime       time.Time
}

func GetCTFList(queryCTF CTF) ([]CTF, error) {
	query := database.GetDatabase().Model(&CTF{})
	if queryCTF.Name != "" {
		query = query.Where("name LIKE ?", "%"+queryCTF.Name+"%")
	}

	var ctfList []CTF
	err := query.
		Preload("CTFChallenges").
		Preload("Users").
		Find(&ctfList).
		Error
	if err != nil {
		return nil, err
	}

	return ctfList, nil
}

func (ctf *CTF) GetCTF() error {
	return database.GetDatabase().Model(&CTF{}).
		Preload("CTFChallenges").
		Preload("Users.CTFChallenges", "ctf_id = ?", ctf.ID).
		Where("ID = ?", ctf.ID).
		First(&ctf).
		Error
}

func (ctf *CTF) CreateCTF() error {
	return database.GetDatabase().Model(&CTF{}).
		Create(&ctf).
		Error
}

func (ctf *CTF) UpdateCTF() error {
	return database.GetDatabase().Model(&CTF{}).
		Select("Name", "Description", "Public", "TeamSignIn", "StartTime", "EndTime").
		Where("ID = ?", ctf.ID).
		Updates(&ctf).
		Error
}

func (ctf *CTF) DeleteCTF() error {
	return database.GetDatabase().Model(&CTF{}).
		Where("ID = ?", ctf.ID).
		Delete(&ctf).
		Error
}

func (ctf *CTF) GetUserList(queryUser User) ([]User, error) {
	query := database.GetDatabase().Model(ctf)
	if queryUser.Username != "" {
		query = query.Where("users.username LIKE ?", "%"+queryUser.Username+"%")
	}

	var userList []User
	err := query.
		Preload("CTFChallenges").
		Association("Users").
		Find(&userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (ctf *CTF) AddUser(user *User) error {
	return database.GetDatabase().Model(ctf).Association("Users").Append(user)
}

func (ctf *CTF) DeleteUser(user *User) error {
	return database.GetDatabase().Model(ctf).Association("Users").Delete(user)
}
