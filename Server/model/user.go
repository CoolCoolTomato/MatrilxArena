package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string
	Password        string
	Email           string
	Role            int
	Challenges      []Challenge      `gorm:"many2many:challenge_user"`
	Groups          []Group          `gorm:"many2many:group_user"`
	GroupChallenges []GroupChallenge `gorm:"many2many:group_challenge_user"`
	CTFs            []CTF            `gorm:"many2many:ctf_user"`
	CTFTeams        []CTFTeam        `gorm:"many2many:ctf_team_user"`
	CTFChallenges   []CTFChallenge   `gorm:"many2many:ctf_challenge_user"`
}

func GetUserList(queryUser User) ([]User, error) {
	query := database.GetDatabase().Model(&User{})
	if queryUser.Username != "" {
		query = query.Where("username LIKE ?", "%"+queryUser.Username+"%")
	}

	var userList []User
	err := query.
		Select("ID", "Username", "Email", "Role").
		Find(&userList).
		Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (user *User) GetUser() error {
	return database.GetDatabase().Model(&User{}).Where("ID = ?", user.ID).Select("ID", "Username", "Email", "Role").First(&user).Error
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
		return database.GetDatabase().Model(&User{}).
			Where("Username = ?", user.Username).
			First(&user).
			Error
	}
	if user.Email != "" {
		return database.GetDatabase().Model(&User{}).
			Where("Email = ?", user.Email).
			First(&user).
			Error
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
	err := database.GetDatabase().Model(user).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Association("Challenges").
		Find(&challengeList)
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

func (user *User) GetGroupList(queryGroup Group) ([]Group, error) {
	query := database.GetDatabase().Model(user)
	if queryGroup.Name != "" {
		query = query.Where("groups.name LIKE ?", "%"+queryGroup.Name+"%")
	}

	var groupList []Group
	err := query.
		Association("Groups").
		Find(&groupList)
	if err != nil {
		return nil, err
	}
	return groupList, nil
}

func (user *User) GetVisibleGroupList(queryGroup Group) ([]Group, error) {
	query := database.GetDatabase().Model(&Group{})
	if queryGroup.Name != "" {
		query = query.Where("name LIKE ?", "%"+queryGroup.Name+"%")
	}

	var groupList []Group
	err := query.
		Where("id IN (?) OR public = ?",
			database.GetDatabase().Table("group_user").Select("group_id").Where("user_id = ?", user.ID), true).
		Find(&groupList).Error
	if err != nil {
		return nil, err
	}
	return groupList, nil
}

func (user *User) AddGroup(group *Group) error {
	return database.GetDatabase().Model(user).Association("Groups").Append(group)
}

func (user *User) DeleteGroup(group *Group) error {
	return database.GetDatabase().Model(user).Association("Groups").Delete(group)
}

func (user *User) GetGroupChallengeList() ([]GroupChallenge, error) {
	var groupChallengeList []GroupChallenge
	err := database.GetDatabase().Model(user).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Group").
		Association("GroupChallenges").
		Find(&groupChallengeList)
	if err != nil {
		return nil, err
	}
	return groupChallengeList, nil
}

func (user *User) AddGroupChallenge(groupChallenge *GroupChallenge) error {
	return database.GetDatabase().Model(user).Association("GroupChallenges").Append(groupChallenge)
}

func (user *User) DeleteGroupChallenge(groupChallenge *GroupChallenge) error {
	return database.GetDatabase().Model(user).Association("GroupChallenges").Delete(groupChallenge)
}
