package model

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"gorm.io/gorm"
)

type DockerNode struct {
	gorm.Model
	Host string
	Port string
}

func GetDockerNodeList() ([]DockerNode, error) {
	var dockerNodeList []DockerNode
	err := database.GetDatabase().Find(&dockerNodeList).Error
	if err != nil {
		return nil, err
	}
	return dockerNodeList, nil
}

func (dockerNode *DockerNode) GetDockerNode() error {
	return database.GetDatabase().Model(&DockerNode{}).Where("ID = ?", dockerNode.ID).First(&dockerNode).Error
}

func (dockerNode *DockerNode) CreateDockerNode() error {
	return database.GetDatabase().Create(&dockerNode).Error
}

func (dockerNode *DockerNode) UpdateDockerNode() error {
	return database.GetDatabase().Model(&DockerNode{}).Where("ID = ?", dockerNode.ID).Updates(&dockerNode).Error
}

func (dockerNode *DockerNode) DeleteDockerNode() error {
	return database.GetDatabase().Model(&DockerNode{}).Where("ID = ?", dockerNode.ID).Delete(&dockerNode).Error
}