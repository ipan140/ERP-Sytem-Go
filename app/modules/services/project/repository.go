package project

import (
	"ERP-System/config"
)

func CreateProject(data *Project) error {
	return config.DB.Create(data).Error
}

func GetAllProject() ([]Project, error) {
	var list []Project
	err := config.DB.Find(&list).Error
	return list, err
}

func GetProjectByID(id uint) (*Project, error) {
	var data Project
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateProject(data *Project) error {
	return config.DB.Save(data).Error
}

func DeleteProject(id uint) error {
	return config.DB.Delete(&Project{}, id).Error
}
