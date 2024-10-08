package model

import (
	"encoding/base64"
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/gormType"
	"gorm.io/gorm"
)

type Challenge struct {
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
	Users          []User `gorm:"many2many:challenge_user"`
}

func GetChallengeList(queryChallenge Challenge) ([]Challenge, error) {
	query := database.GetDatabase().Model(&Challenge{})
	if queryChallenge.CategoryID != 0 {
		query = query.Where("category_id = ?", queryChallenge.CategoryID)
	}
	if queryChallenge.Title != "" {
		query = query.Where("title LIKE ?", "%"+queryChallenge.Title+"%")
	}

	var challengeList []Challenge
	err := query.
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Find(&challengeList).
		Error
	if err != nil {
		return nil, err
	}

	for i, _ := range challengeList {
		challengeList[i].DecodeCommands()
	}

	return challengeList, nil
}

func (challenge *Challenge) GetChallenge() error {
	err := database.GetDatabase().Model(&Challenge{}).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Where("ID = ?", challenge.ID).
		First(&challenge).
		Error
	challenge.DecodeCommands()
	return err
}

func (challenge *Challenge) CreateChallenge() error {
	challenge.EncodeCommands()
	createFields := []string{
		"title",
		"description",
		"category_id",
		"image_id",
		"attachment_id",
		"specified_ports",
		"commands",
		"flag",
	}

	creates := map[string]interface{}{
		"title":           challenge.Title,
		"description":     challenge.Description,
		"category_id":     challenge.CategoryID,
		"image_id":        challenge.ImageID,
		"attachment_id":   challenge.AttachmentID,
		"specified_ports": challenge.SpecifiedPorts,
		"commands":        challenge.Commands,
		"flag":            challenge.Flag,
	}

	if challenge.CategoryID == 0 {
		creates["category_id"] = nil
	}

	if challenge.ImageID == 0 {
		creates["image_id"] = nil
	}

	if challenge.AttachmentID == 0 {
		creates["attachment_id"] = nil
	}

	return database.GetDatabase().
		Model(&Challenge{}).
		Select(createFields).
		Create(creates).
		Error
}

func (challenge *Challenge) UpdateChallenge() error {
	challenge.EncodeCommands()
	updateFields := []string{
		"title",
		"description",
		"category_id",
		"image_id",
		"attachment_id",
		"specified_ports",
		"commands",
		"flag",
	}

	updates := map[string]interface{}{
		"ID":              challenge.ID,
		"title":           challenge.Title,
		"description":     challenge.Description,
		"category_id":     challenge.CategoryID,
		"image_id":        challenge.ImageID,
		"attachment_id":   challenge.AttachmentID,
		"specified_ports": challenge.SpecifiedPorts,
		"commands":        challenge.Commands,
		"flag":            challenge.Flag,
	}

	if challenge.CategoryID == 0 {
		updates["category_id"] = nil
	}

	if challenge.ImageID == 0 {
		updates["image_id"] = nil
	}

	if challenge.AttachmentID == 0 {
		updates["attachment_id"] = nil
	}

	return database.GetDatabase().
		Model(&Challenge{}).
		Where("ID = ?", challenge.ID).
		Select(updateFields).
		Updates(updates).
		Error
}

func (challenge *Challenge) DeleteChallenge() error {
	return database.GetDatabase().Model(&Challenge{}).
		Where("ID = ?", challenge.ID).
		Delete(&challenge).
		Error
}

func (challenge *Challenge) EncodeCommands() {
	var encodedCommands gormType.StringSlice
	for _, command := range challenge.Commands {
		encodedCommand := base64.StdEncoding.EncodeToString([]byte(command))
		encodedCommands = append(encodedCommands, encodedCommand)
	}
	challenge.Commands = encodedCommands
}

func (challenge *Challenge) DecodeCommands() {
	var decodedCommands gormType.StringSlice
	for _, command := range challenge.Commands {
		decodedCommand, _ := base64.StdEncoding.DecodeString(command)
		decodedCommands = append(decodedCommands, string(decodedCommand))
	}
	challenge.Commands = decodedCommands
}
