package recruitment

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateApplicant(data *Applicant) error {
	return config.DB.Create(data).Error
}

func GetAllApplicant() ([]Applicant, error) {
	var list []Applicant
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedApplicants(offset, limit int, search string, jobPositionID string, stageID string, state string) ([]Applicant, int64, error) {
	var list []Applicant
	var total int64

	query := config.DB.Model(&Applicant{})

	if jobPositionID != "" && jobPositionID != "all" && jobPositionID != "0" {
		query = query.Where("job_position_id = ?", jobPositionID)
	}

	if stageID != "" && stageID != "all" && stageID != "0" {
		query = query.Where("stage_id = ?", stageID)
	}

	if state != "" && state != "all" {
		query = query.Where("state = ?", state)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ? OR phone ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetApplicantByID(id uint) (*Applicant, error) {
	var data Applicant
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateApplicant(data *Applicant) error {
	return config.DB.Save(data).Error
}

func DeleteApplicant(id uint) error {
	return config.DB.Delete(&Applicant{}, id).Error
}
