package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
    Remark          string
	RepoTags		string
	Repository		string
}

func GetImageList() ([]Image, error) {
	var imageList []Image
	err := database.GetDatabase().Find(&imageList).Error
	if err != nil {
		return nil, err
	}
	return imageList, nil
}

func (image *Image) GetImage() error {
	return database.GetDatabase().Model(&Image{}).Where("ID = ?", image.ID).First(&image).Error
}

func (image *Image) CreateImage() error {
	return database.GetDatabase().Create(&image).Error
}

func (image *Image) UpdateImage() error {
	return database.GetDatabase().Model(&Image{}).Where("ID = ?", image.ID).Updates(&image).Error
}

func (image *Image) DeleteImage() error {
	return database.GetDatabase().Model(&Image{}).Where("ID = ?", image.ID).Delete(&image).Error
}
