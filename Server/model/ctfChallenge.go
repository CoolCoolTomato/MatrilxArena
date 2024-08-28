package model

import (
    "encoding/base64"
    "github.com/CoolCoolTomato/MatrilxArena/Server/database"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/gormType"
    "gorm.io/gorm"
)

type CTFChallenge struct {
	gorm.Model
	Title          string
	Description    string
	CategoryID     uint
	Category       Category             `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL"`
	ImageID        uint
	Image          Image                `gorm:"foreignKey:ImageID;constraint:OnDelete:SET NULL"`
	AttachmentID   uint
	Attachment     Attachment           `gorm:"foreignKey:AttachmentID;constraint:OnDelete:SET NULL"`
	SpecifiedPorts gormType.StringSlice `gorm:"type:json"`
	Commands       gormType.StringSlice `gorm:"type:json"`
	Flag           string
    Score          int
	Users          []User               `gorm:"many2many:ctf_challenge_user"`
	CTFID          uint
	CTF            CTF                  `gorm:"foreignKey:CTFID;constraint:OnDelete:SET NULL"`
}

func GetCTFChallengeList(queryCTFChallenge CTFChallenge) ([]CTFChallenge, error) {
    query := database.GetDatabase().Model(&CTFChallenge{})
	if queryCTFChallenge.CTFID != 0 {
		query = query.Where("ctf_id = ?", queryCTFChallenge.CTFID)
	}
	if queryCTFChallenge.CategoryID != 0 {
		query = query.Where("category_id = ?", queryCTFChallenge.CategoryID)
	}
	if queryCTFChallenge.Title != "" {
		query = query.Where("title LIKE ?", "%"+queryCTFChallenge.Title+"%")
	}

	var ctfChallengeList []CTFChallenge
	err := query.
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Users").
        Preload("CTF").
		Find(&ctfChallengeList).
		Error
	if err != nil {
		return nil, err
	}

	for i, _ := range ctfChallengeList {
		ctfChallengeList[i].DecodeCommands()
	}

	return ctfChallengeList, nil
}

func (ctfChallenge *CTFChallenge) GetCTFChallenge() error {
	err := database.GetDatabase().Model(&CTFChallenge{}).
		Preload("Category").
		Preload("Image").
		Preload("Attachment").
		Preload("Users").
        Preload("CTF").
		Where("ID = ?", ctfChallenge.ID).
		First(&ctfChallenge).
		Error
	ctfChallenge.DecodeCommands()
	return err
}

func (ctfChallenge *CTFChallenge) CreateCTFChallenge() error {
	ctfChallenge.EncodeCommands()
	createFields := []string{
		"title",
		"description",
		"category_id",
		"image_id",
		"attachment_id",
		"specified_ports",
		"commands",
		"flag",
        "score",
		"ctf_id",
	}

	creates := map[string]interface{}{
		"title":           ctfChallenge.Title,
		"description":     ctfChallenge.Description,
		"category_id":     ctfChallenge.CategoryID,
		"image_id":        ctfChallenge.ImageID,
		"attachment_id":   ctfChallenge.AttachmentID,
		"specified_ports": ctfChallenge.SpecifiedPorts,
		"commands":        ctfChallenge.Commands,
		"flag":            ctfChallenge.Flag,
        "score":           ctfChallenge.Score,
		"ctf_id":          ctfChallenge.CTFID,
	}

	if ctfChallenge.CategoryID == 0 {
		creates["category_id"] = nil
	}

	if ctfChallenge.ImageID == 0 {
		creates["image_id"] = nil
	}

	if ctfChallenge.AttachmentID == 0 {
		creates["attachment_id"] = nil
	}

	if ctfChallenge.CTFID == 0 {
		creates["ctf_id"] = nil
	}

	return database.GetDatabase().
		Model(&CTFChallenge{}).
		Select(createFields).
		Create(creates).
		Error
}

func (ctfChallenge *CTFChallenge) UpdateCTFChallenge() error {
	ctfChallenge.EncodeCommands()
	updateFields := []string{
		"title",
		"description",
		"category_id",
		"image_id",
		"attachment_id",
		"specified_ports",
		"commands",
		"flag",
        "score",
		"ctf_id",
	}

	updates := map[string]interface{}{
		"ID":              ctfChallenge.ID,
		"title":           ctfChallenge.Title,
		"description":     ctfChallenge.Description,
		"category_id":     ctfChallenge.CategoryID,
		"image_id":        ctfChallenge.ImageID,
		"attachment_id":   ctfChallenge.AttachmentID,
		"specified_ports": ctfChallenge.SpecifiedPorts,
		"commands":        ctfChallenge.Commands,
		"flag":            ctfChallenge.Flag,
        "score":           ctfChallenge.Score,
		"ctf_id":          ctfChallenge.CTFID,
	}

	if ctfChallenge.CategoryID == 0 {
		updates["category_id"] = nil
	}

	if ctfChallenge.ImageID == 0 {
		updates["image_id"] = nil
	}

	if ctfChallenge.AttachmentID == 0 {
		updates["attachment_id"] = nil
	}

	if ctfChallenge.CTFID == 0 {
		updates["ctf_id"] = nil
	}

	return database.GetDatabase().
		Model(&CTFChallenge{}).
		Where("ID = ?", ctfChallenge.ID).
		Select(updateFields).
		Updates(updates).
		Error
}

func (ctfChallenge *CTFChallenge) DeleteCTFChallenge() error {
	return database.GetDatabase().Model(&CTFChallenge{}).
		Where("ID = ?", ctfChallenge.ID).
		Delete(&ctfChallenge).
		Error
}

func (ctfChallenge *CTFChallenge) EncodeCommands() {
	var encodedCommands gormType.StringSlice
	for _, command := range ctfChallenge.Commands {
		encodedCommand := base64.StdEncoding.EncodeToString([]byte(command))
		encodedCommands = append(encodedCommands, encodedCommand)
	}
	ctfChallenge.Commands = encodedCommands
}

func (ctfChallenge *CTFChallenge) DecodeCommands() {
	var decodedCommands gormType.StringSlice
	for _, command := range ctfChallenge.Commands {
		decodedCommand, _ := base64.StdEncoding.DecodeString(command)
		decodedCommands = append(decodedCommands, string(decodedCommand))
	}
	ctfChallenge.Commands = decodedCommands
}
