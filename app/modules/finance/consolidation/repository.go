package consolidation

import (
	"ERP-System/config"
)

func CreateConsolidationReport(data *ConsolidationReport) error {
	return config.DB.Create(data).Error
}

func GetAllConsolidationReports() ([]ConsolidationReport, error) {
	var list []ConsolidationReport
	err := config.DB.Find(&list).Error
	return list, err
}
