package surveys

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateSurvey(data *Survey) error {
	return config.DB.Create(data).Error
}

func GetAllSurvey() ([]Survey, error) {
	var list []Survey
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedSurveys(offset, limit int, search string) ([]Survey, int64, error) {
	var list []Survey
	var total int64
	query := config.DB.Model(&Survey{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ? OR state ILIKE ?", s, s, s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetSurveyByID(id uint) (*Survey, error) {
	var data Survey
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateSurvey(data *Survey) error {
	return config.DB.Save(data).Error
}

func DeleteSurvey(id uint) error {
	return config.DB.Delete(&Survey{}, id).Error
}

// RecordSurveyResponse mencatat respon (dari GForm webhook atau form publik), lalu menghitung ulang NPS Score
func RecordSurveyResponse(id uint, rating int) (*Survey, error) {
	var survey Survey
	if err := config.DB.First(&survey, id).Error; err != nil {
		return nil, err
	}

	if rating >= 9 {
		survey.PromotersCount++
	} else if rating >= 7 {
		survey.PassivesCount++
	} else {
		survey.DetractorsCount++
	}
	survey.ResponsesCount++

	// Hitung ulang NPS = (% Promoter - % Detractor) * 100
	if survey.ResponsesCount > 0 {
		promoterPct := float64(survey.PromotersCount) / float64(survey.ResponsesCount) * 100
		detractorPct := float64(survey.DetractorsCount) / float64(survey.ResponsesCount) * 100
		survey.NpsScore = int(promoterPct - detractorPct)
	}

	if err := config.DB.Save(&survey).Error; err != nil {
		return nil, err
	}
	return &survey, nil
}

