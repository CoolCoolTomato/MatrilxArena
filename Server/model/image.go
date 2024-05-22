package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	DockerID		string
	RepoTags		string
	Repository		string
}

func (image *Image) GetImage() error {
	return database.GetDatabase().First(&image).Error
}

func (image *Image) CreateImage() error {
	return database.GetDatabase().Create(&image).Error
}

func (image *Image) UpdateImage() error {
	return database.GetDatabase().Updates(&image).Error
}

func (image *Image) DeleteImage() error {
	return database.GetDatabase().Delete(&image).Error
}
