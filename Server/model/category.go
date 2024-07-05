package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string
	Order uint
}

func GetCategoryList() ([]Category, error) {
	var categoryList []Category
	err := database.GetDatabase().Model(&Category{}).Find(&categoryList).Error
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}

func (category *Category) GetCategory() error {
	return database.GetDatabase().Model(&Category{}).Where("ID = ?", category.ID).First(&category).Error
}

func (category *Category) CreateCategory() error {
	return database.GetDatabase().Model(&Category{}).Create(&category).Error
}

func (category *Category) UpdateCategory() error {
	return database.GetDatabase().Model(&Category{}).Where("ID = ?", category.ID).Updates(&category).Error
}

func (category *Category) DeleteCategory() error {
	return database.GetDatabase().Model(&Category{}).Where("ID = ?", category.ID).Delete(&category).Error
}
