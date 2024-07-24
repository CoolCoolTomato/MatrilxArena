package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
	"strings"
)

type Image struct {
	gorm.Model
	Remark     string
	RepoTags   string
	Repository string
}

func GetImageList(queryImage Image) ([]Image, error) {
	query := database.GetDatabase().Model(&Image{})
	var conditions []string
	var values []interface{}

	if queryImage.Remark != "" {
		conditions = append(conditions, "remark LIKE ?")
		values = append(values, "%"+queryImage.Remark+"%")
	}

	if queryImage.RepoTags != "" {
		conditions = append(conditions, "repo_tags LIKE ?")
		values = append(values, "%"+queryImage.RepoTags+"%")
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	var imageList []Image
	err := query.
		Find(&imageList).
		Error
	if err != nil {
		return nil, err
	}
	return imageList, nil
}

func (image *Image) GetImage() error {
	return database.GetDatabase().Model(&Image{}).Where("ID = ?", image.ID).First(&image).Error
}

func (image *Image) CreateImage() error {
	return database.GetDatabase().Model(&Image{}).Create(&image).Error
}

func (image *Image) UpdateImage() error {
	return database.GetDatabase().Model(&Image{}).Where("ID = ?", image.ID).Updates(&image).Error
}

func (image *Image) DeleteImage() error {
	return database.GetDatabase().Model(&Image{}).Where("ID = ?", image.ID).Delete(&image).Error
}
