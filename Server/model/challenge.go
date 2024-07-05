package model

import (
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

func GetChallengeList() ([]Challenge, error) {
	var challengeList []Challenge
	err := database.GetDatabase().Model(&Challenge{}).Preload("Category").Preload("Image").Preload("Attachment").Find(&challengeList).Error
	if err != nil {
		return nil, err
	}
	return challengeList, nil
}

func GetChallengeListByQuery(challenge Challenge) ([]Challenge, error) {
	var challengeList []Challenge
	query := database.GetDatabase().Model(&Challenge{})
	if challenge.CategoryID != 0 {
		query = query.Where("category_id = ?", challenge.CategoryID)
	}
	if challenge.Title != "" {
		query = query.Where("title LIKE ?", "%"+challenge.Title+"%")
	}
	err := query.
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Find(&challengeList).Error
	if err != nil {
		return nil, err
	}
	return challengeList, nil
}

func (challenge *Challenge) GetChallenge() error {
	return database.GetDatabase().Model(&Challenge{}).Preload("Category").Preload("Image").Preload("Attachment").Where("ID = ?", challenge.ID).First(&challenge).Error
}

func (challenge *Challenge) CreateChallenge() error {
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
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Select(createFields).
		Create(creates).
		Error
}

func (challenge *Challenge) UpdateChallenge() error {
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
		Preload("Image").
		Preload("Category").
		Preload("Attachment").
		Where("ID = ?", challenge.ID).
		Select(updateFields).
		Updates(updates).
		Error
}

func (challenge *Challenge) DeleteChallenge() error {
	return database.GetDatabase().Model(&Challenge{}).Preload("Category").Preload("Image").Preload("Attachment").Where("ID = ?", challenge.ID).Delete(&challenge).Error
}
