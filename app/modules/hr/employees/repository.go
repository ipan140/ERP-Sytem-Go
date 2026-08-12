package employees

import (
	"ERP-System/config"
)

func CreateEmployee(data *Employee) error {
	return config.DB.Create(data).Error
}

func GetAllEmployee() ([]Employee, error) {
	var list []Employee
	err := config.DB.Find(&list).Error
	return list, err
}

func GetEmployeeByID(id uint) (*Employee, error) {
	var data Employee
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateEmployee(data *Employee) error {
	return config.DB.Save(data).Error
}

func DeleteEmployee(id uint) error {
	return config.DB.Delete(&Employee{}, id).Error
}
