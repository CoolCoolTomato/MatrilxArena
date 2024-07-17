package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	categoryList, err := model.GetCategoryList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Category.GetCategoryListFail", c), c)
		return
	}

	response.OK(categoryList, localizer.GetMessage("Category.GetCategoryListSuccess", c), c)
}

func GetCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil || category.ID == 0 {
		response.Fail(err, localizer.GetMessage("Category.InvalidArgument", c), c)
		return
	}

	err = category.GetCategory()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Category.GetCategoryFail", c), c)
		return
	}

	response.OK(category, localizer.GetMessage("Category.GetCategorySuccess", c), c)
}

func CreateCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Category.InvalidArgument", c), c)
		return
	}

	err = category.CreateCategory()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Category.CreateCategoryFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Category.CreateCategorySuccess", c), c)
}

func UpdateCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil || category.ID == 0 {
		response.Fail(err, localizer.GetMessage("Category.InvalidArgument", c), c)
		return
	}

	err = category.UpdateCategory()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Category.UpdateCategoryFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Category.UpdateCategorySuccess", c), c)
}

func DeleteCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil || category.ID == 0 {
		response.Fail(err, localizer.GetMessage("Category.InvalidArgument", c), c)
		return
	}

	err = category.DeleteCategory()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Category.DeleteCategoryFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Category.DeleteCategorySuccess", c), c)
}
