package appraisals

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAppraisal(data *Appraisal) error {
	return config.DB.Create(data).Error
}

func GetAllAppraisal() ([]Appraisal, error) {
	var list []Appraisal
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedAppraisals(offset, limit int, search string, employeeID string, state string) ([]Appraisal, int64, error) {
	var list []Appraisal
	var total int64

	query := config.DB.Model(&Appraisal{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.appraisals.employee_id = ?", employeeID)
	}

	if state != "" && state != "all" {
		query = query.Where("hrd.appraisals.state = ?", state)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.appraisals.employee_id").
			Where("hrd.employees.name ILIKE ? OR hrd.appraisals.feedback ILIKE ? OR hrd.appraisals.state ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.appraisals.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetAppraisalByID(id uint) (*Appraisal, error) {
	var data Appraisal
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAppraisal(data *Appraisal) error {
	return config.DB.Save(data).Error
}

func DeleteAppraisal(id uint) error {
	return config.DB.Delete(&Appraisal{}, id).Error
}
