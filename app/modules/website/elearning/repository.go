package elearning

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateCourse(data *Course) error {
	return config.DB.Create(data).Error
}

func GetAllCourse() ([]Course, error) {
	var list []Course
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedCourses(offset, limit int, search string) ([]Course, int64, error) {
	var list []Course
	var total int64
	query := config.DB.Model(&Course{}).Preload(clause.Associations)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetCourseByID(id uint) (*Course, error) {
	var data Course
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateCourse(data *Course) error {
	return config.DB.Save(data).Error
}

func DeleteCourse(id uint) error {
	return config.DB.Delete(&Course{}, id).Error
}

func CreateSlide(data *Slide) error { return config.DB.Create(data).Error }
func GetAllSlide() ([]Slide, error) {
	var list []Slide
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetSlideByID(id uint) (*Slide, error) {
	var data Slide
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateSlide(data *Slide) error { return config.DB.Save(data).Error }
func DeleteSlide(id uint) error     { return config.DB.Delete(&Slide{}, id).Error }

func CreateCertification(data *Certification) error { return config.DB.Create(data).Error }
func GetAllCertification() ([]Certification, error) {
	var list []Certification
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetCertificationByID(id uint) (*Certification, error) {
	var data Certification
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateCertification(data *Certification) error { return config.DB.Save(data).Error }
func DeleteCertification(id uint) error             { return config.DB.Delete(&Certification{}, id).Error }
