package field_service

import (
	"ERP-System/config"
)

func CreateFieldServiceTask(data *FieldServiceTask) error {
	return config.DB.Create(data).Error
}

func GetAllFieldServiceTask() ([]FieldServiceTask, error) {
	var list []FieldServiceTask
	err := config.DB.Find(&list).Error
	return list, err
}

func GetFieldServiceTaskByID(id uint) (*FieldServiceTask, error) {
	var data FieldServiceTask
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateFieldServiceTask(data *FieldServiceTask) error {
	return config.DB.Save(data).Error
}

func DeleteFieldServiceTask(id uint) error {
	return config.DB.Delete(&FieldServiceTask{}, id).Error
}
