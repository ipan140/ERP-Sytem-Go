package elearning

import (
	"ERP-System/config"
)

func CreateCourse(data *Course) error {
	return config.DB.Create(data).Error
}

func GetAllCourse() ([]Course, error) {
	var list []Course
	err := config.DB.Find(&list).Error
	return list, err
}

func GetCourseByID(id uint) (*Course, error) {
	var data Course
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateCourse(data *Course) error {
	return config.DB.Save(data).Error
}

func DeleteCourse(id uint) error {
	return config.DB.Delete(&Course{}, id).Error
}
