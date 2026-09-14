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

func GetPaginatedTaxReportsRepo(offset int, limit int, search string, taxType string) ([]TaxReportSummary, int64, float64, error) {
	var list []TaxReportSummary
	var total int64

	var summary struct {
		TotalTax float64
	}
	config.DB.Model(&TaxReportSummary{}).Select("COALESCE(SUM(tax_amount), 0) as total_tax").Scan(&summary)

	query := config.DB.Model(&TaxReportSummary{})
	if taxType != "" && taxType != "all" && taxType != "All" && taxType != "Semua" {
		query = query.Where("tax_type = ?", taxType)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("partner_name ILIKE ? OR npwp ILIKE ? OR tax_period ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	err := query.Order("id asc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, summary.TotalTax, err
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
