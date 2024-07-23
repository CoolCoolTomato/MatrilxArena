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

func GetAttachmentList(queryAttachment Attachment) ([]Attachment, error) {
	query := database.GetDatabase().Model(&Attachment{})
	if queryAttachment.FileName != "" {
		query = query.Where("file_name LIKE ?", "%"+queryAttachment.FileName+"%")
	}

	var attachmentList []Attachment
	err := query.
		Find(&attachmentList).
		Error
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
