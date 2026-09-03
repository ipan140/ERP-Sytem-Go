package tax

import (
	"ERP-System/config"
)

func GetAllTaxMasterConfigsRepo() ([]TaxMasterConfig, error) {
	var list []TaxMasterConfig
	err := config.DB.Order("id asc").Find(&list).Error
	return list, err
}

func GetTaxMasterConfigByIDRepo(id uint) (*TaxMasterConfig, error) {
	var item TaxMasterConfig
	err := config.DB.First(&item, id).Error
	return &item, err
}

func CreateTaxMasterConfigRepo(item *TaxMasterConfig) error {
	return config.DB.Create(item).Error
}

func UpdateTaxMasterConfigRepo(item *TaxMasterConfig) error {
	return config.DB.Save(item).Error
}

func DeleteTaxMasterConfigRepo(id uint) error {
	return config.DB.Delete(&TaxMasterConfig{}, id).Error
}

func GetAllTaxReportsRepo() ([]TaxReportSummary, error) {
	var list []TaxReportSummary
	err := config.DB.Order("id asc").Find(&list).Error
	return list, err
}

func GetTaxReportByIDRepo(id uint) (*TaxReportSummary, error) {
	var item TaxReportSummary
	err := config.DB.First(&item, id).Error
	return &item, err
}

func CreateTaxReportRepo(t *TaxReportSummary) error {
	return config.DB.Create(t).Error
}

func UpdateTaxReportRepo(t *TaxReportSummary) error {
	return config.DB.Save(t).Error
}

func DeleteTaxReportRepo(id uint) error {
	return config.DB.Delete(&TaxReportSummary{}, id).Error
}
