package consolidation

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateConsolidationReport(data *ConsolidationReport) error {
	return config.DB.Create(data).Error
}

func GetAllConsolidationReports() ([]ConsolidationReport, error) {
	var list []ConsolidationReport
	err := config.DB.Order("created_at desc").Find(&list).Error
	return list, err
}

func GetConsolidationReportByID(id uint) (*ConsolidationReport, error) {
	var data ConsolidationReport
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateConsolidationReport(data *ConsolidationReport) error {
	return config.DB.Model(&ConsolidationReport{}).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"name":     data.Name,
		"period":   data.Period,
		"branches": data.Branches,
	}).Error
}

func DeleteConsolidationReport(id uint) error {
	config.DB.Where("report_id = ?", id).Delete(&ConsolidatedAccount{})
	return config.DB.Delete(&ConsolidationReport{}, id).Error
}
