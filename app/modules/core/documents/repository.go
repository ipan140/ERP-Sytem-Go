package documents

import (
	"ERP-System/config"
)

func CreateWorkspace(data *Workspace) error {
	return config.DB.Create(data).Error
}

func GetAllWorkspace() ([]Workspace, error) {
	var list []Workspace
	err := config.DB.Find(&list).Error
	return list, err
}

func GetWorkspaceByID(id uint) (*Workspace, error) {
	var data Workspace
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateWorkspace(data *Workspace) error {
	return config.DB.Save(data).Error
}

func DeleteWorkspace(id uint) error {
	return config.DB.Delete(&Workspace{}, id).Error
}
