package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	categoryList, err := model.GetCategoryList()
	if err != nil {
		response.Fail(err, "Get category list fail", c)
		return
	}

	response.OK(categoryList, "Get category list success", c)
}

func GetCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil || category.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = category.GetCategory()
	if err != nil {
		response.Fail(err, "Get category fail", c)
		return
	}

	response.OK(category, "Get category success", c)
}

func CreateCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = category.CreateCategory()
	if err != nil {
		response.Fail(err, "Create category fail", c)
		return
	}

	response.OK(nil, "Create category success", c)
}

func UpdateCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil || category.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = category.UpdateCategory()
	if err != nil {
		response.Fail(err, "Update category fail", c)
		return
	}

	response.OK(nil, "Update category success", c)
}

func DeleteCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil || category.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = category.DeleteCategory()
	if err != nil {
		response.Fail(err, "Delete category fail", c)
		return
	}

	response.OK(nil, "Delete category success", c)
}
