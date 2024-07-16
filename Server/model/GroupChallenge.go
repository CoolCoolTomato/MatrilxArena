package model

import (
	"encoding/base64"
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/gormType"
	"gorm.io/gorm"
)

type GroupChallenge struct {
	gorm.Model
	Title          string
	Description    string
	CategoryID     uint
	Category       Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL"`
	ImageID        uint
	Image          Image `gorm:"foreignKey:ImageID;constraint:OnDelete:SET NULL"`
	AttachmentID   uint
	Attachment     Attachment           `gorm:"foreignKey:AttachmentID;constraint:OnDelete:SET NULL"`
	SpecifiedPorts gormType.StringSlice `gorm:"type:json"`
	Commands       gormType.StringSlice `gorm:"type:json"`
	Flag           string
	Users          []User `gorm:"many2many:group_challenge_user"`
	GroupID        uint
	Group          Group `gorm:"foreignKey:GroupID;constraint:OnDelete:SET NULL"`
}

func GetGroupChallengeList() ([]GroupChallenge, error) {
	var groupChallengeList []GroupChallenge
	err := database.GetDatabase().Model(&GroupChallenge{}).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Group").
		Find(&groupChallengeList).
		Error
	if err != nil {
		return nil, err
	}

	for i, _ := range groupChallengeList {
		groupChallengeList[i].DecodeCommands()
	}

	return groupChallengeList, nil
}

func GetGroupChallengeListByQuery(queryGroupChallenge GroupChallenge) ([]GroupChallenge, error) {
	query := database.GetDatabase().Model(&GroupChallenge{})
	if queryGroupChallenge.CategoryID != 0 {
		query = query.Where("category_id = ?", queryGroupChallenge.CategoryID)
	}
	if queryGroupChallenge.Title != "" {
		query = query.Where("title LIKE ?", "%"+queryGroupChallenge.Title+"%")
	}

	var groupChallengeList []GroupChallenge
	err := query.
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Group").
		Find(&groupChallengeList).
		Error
	if err != nil {
		return nil, err
	}

	for i, _ := range groupChallengeList {
		groupChallengeList[i].DecodeCommands()
	}

	return groupChallengeList, nil
}

func (groupChallenge *GroupChallenge) GetGroupChallenge() error {
	err := database.GetDatabase().Model(&GroupChallenge{}).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Group").
		Where("ID = ?", groupChallenge.ID).
		First(&groupChallenge).
		Error
	groupChallenge.DecodeCommands()
	return err
}

func (groupChallenge *GroupChallenge) CreateGroupChallenge() error {
	groupChallenge.EncodeCommands()
	createFields := []string{
		"title",
		"description",
		"category_id",
		"image_id",
		"attachment_id",
		"specified_ports",
		"commands",
		"flag",
		"group_id",
	}

	creates := map[string]interface{}{
		"title":           groupChallenge.Title,
		"description":     groupChallenge.Description,
		"category_id":     groupChallenge.CategoryID,
		"image_id":        groupChallenge.ImageID,
		"attachment_id":   groupChallenge.AttachmentID,
		"specified_ports": groupChallenge.SpecifiedPorts,
		"commands":        groupChallenge.Commands,
		"flag":            groupChallenge.Flag,
		"group_id":        groupChallenge.GroupID,
	}

	if groupChallenge.CategoryID == 0 {
		creates["category_id"] = nil
	}

	if groupChallenge.ImageID == 0 {
		creates["image_id"] = nil
	}

	if groupChallenge.AttachmentID == 0 {
		creates["attachment_id"] = nil
	}

	if groupChallenge.GroupID == 0 {
		creates["group_id"] = nil
	}

	return database.GetDatabase().
		Model(&GroupChallenge{}).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Group").
		Select(createFields).
		Create(creates).
		Error
}

func (groupChallenge *GroupChallenge) UpdateGroupChallenge() error {
	groupChallenge.EncodeCommands()
	updateFields := []string{
		"title",
		"description",
		"category_id",
		"image_id",
		"attachment_id",
		"specified_ports",
		"commands",
		"flag",
		"group_id",
	}

	updates := map[string]interface{}{
		"ID":              groupChallenge.ID,
		"title":           groupChallenge.Title,
		"description":     groupChallenge.Description,
		"category_id":     groupChallenge.CategoryID,
		"image_id":        groupChallenge.ImageID,
		"attachment_id":   groupChallenge.AttachmentID,
		"specified_ports": groupChallenge.SpecifiedPorts,
		"commands":        groupChallenge.Commands,
		"flag":            groupChallenge.Flag,
		"group_id":        groupChallenge.GroupID,
	}

	if groupChallenge.CategoryID == 0 {
		updates["category_id"] = nil
	}

	if groupChallenge.ImageID == 0 {
		updates["image_id"] = nil
	}

	if groupChallenge.AttachmentID == 0 {
		updates["attachment_id"] = nil
	}

	if groupChallenge.GroupID == 0 {
		updates["group_id"] = nil
	}

	return database.GetDatabase().
		Model(&GroupChallenge{}).
		Preload("Image").
		Preload("Category").
		Preload("Attachment").
		Preload("Group").
		Where("ID = ?", groupChallenge.ID).
		Select(updateFields).
		Updates(updates).
		Error
}

func (groupChallenge *GroupChallenge) DeleteGroupChallenge() error {
	return database.GetDatabase().Model(&GroupChallenge{}).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Group").
		Where("ID = ?", groupChallenge.ID).
		Delete(&groupChallenge).
		Error
}

func (groupChallenge *GroupChallenge) EncodeCommands() {
	var encodedCommands gormType.StringSlice
	for _, command := range groupChallenge.Commands {
		encodedCommand := base64.StdEncoding.EncodeToString([]byte(command))
		encodedCommands = append(encodedCommands, encodedCommand)
	}
	groupChallenge.Commands = encodedCommands
}

func (groupChallenge *GroupChallenge) DecodeCommands() {
	var decodedCommands gormType.StringSlice
	for _, command := range groupChallenge.Commands {
		decodedCommand, _ := base64.StdEncoding.DecodeString(command)
		decodedCommands = append(decodedCommands, string(decodedCommand))
	}
	groupChallenge.Commands = decodedCommands
}
