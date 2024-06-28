package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	FileName string
	FilePath string
}

func GetAttachmentList() ([]Attachment, error) {
	var attachmentList []Attachment
	err := database.GetDatabase().Model(&Attachment{}).Find(&attachmentList).Error
	if err != nil {
		return nil, err
	}
	return attachmentList, nil
}

func (attachment *Attachment) GetAttachment() error {
	return database.GetDatabase().Model(&Attachment{}).Where("ID = ?", attachment.ID).First(&attachment).Error
}

func (attachment *Attachment) CreateAttachment() error {
	return database.GetDatabase().Model(&Attachment{}).Create(&attachment).Error
}

func (attachment *Attachment) UpdateAttachment() error {
	return database.GetDatabase().Model(&Attachment{}).Where("ID = ?", attachment.ID).Updates(&attachment).Error
}

func (attachment *Attachment) DeleteAttachment() error {
	return database.GetDatabase().Model(&Attachment{}).Where("ID = ?", attachment.ID).Delete(&attachment).Error
}
