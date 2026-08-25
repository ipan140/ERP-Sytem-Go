package quality

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateQualityCheck(data *QualityCheck) error {
	return config.DB.Create(data).Error
}

func GetAllQualityCheck() ([]QualityCheck, error) {
	var list []QualityCheck
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetQualityCheckByID(id uint) (*QualityCheck, error) {
	var data QualityCheck
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateQualityCheck(data *QualityCheck) error {
	return config.DB.Save(data).Error
}

func DeleteQualityCheck(id uint) error {
	return config.DB.Delete(&QualityCheck{}, id).Error
}
