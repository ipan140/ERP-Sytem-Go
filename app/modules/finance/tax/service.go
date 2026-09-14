package tax

import (
	"fmt"
)

// Master Tax Config Service
func GetAllTaxMasterConfigsService() ([]TaxMasterConfig, error) {
	return GetAllTaxMasterConfigsRepo()
}

func CreateTaxMasterConfigService(item *TaxMasterConfig) error {
	item.IsActive = true
	return CreateTaxMasterConfigRepo(item)
}

func UpdateTaxMasterConfigService(item *TaxMasterConfig) error {
	return UpdateTaxMasterConfigRepo(item)
}

func DeleteTaxMasterConfigService(id uint) error {
	return DeleteTaxMasterConfigRepo(id)
}

// Transaction Summary Service
func GetAllTaxReportsService() ([]TaxReportSummary, float64, error) {
	list, err := GetAllTaxReportsRepo()
	if err != nil {
		return nil, 0, err
	}

	var totalTaxAmount float64
	for _, it := range list {
		totalTaxAmount += it.TaxAmount
	}
	return list, totalTaxAmount, nil
}

func GetPaginatedTaxReportsService(offset int, limit int, search string, taxType string) ([]TaxReportSummary, int64, float64, error) {
	return GetPaginatedTaxReportsRepo(offset, limit, search, taxType)
}

func CreateTaxReportService(t *TaxReportSummary) error {
	if t.TaxAmount == 0 && t.TaxBase > 0 && t.TaxRate > 0 {
		t.TaxAmount = (t.TaxBase * t.TaxRate) / 100
	}
	if t.Status == "" {
		t.Status = "Siap Lapor"
	}
	return CreateTaxReportRepo(t)
}

func UpdateTaxReportService(t *TaxReportSummary) error {
	if t.TaxBase > 0 && t.TaxRate > 0 {
		t.TaxAmount = (t.TaxBase * t.TaxRate) / 100
	}
	return UpdateTaxReportRepo(t)
}

func DeleteTaxReportService(id uint) error {
	return DeleteTaxReportRepo(id)
}

func GenerateEBupotCsvStringService() (string, error) {
	list, err := GetAllTaxReportsRepo()
	if err != nil {
		return "", err
	}

	csvContent := "Masa Pajak;Jenis Pajak;NPWP Pemotong;Nama Entitas;DPP;Tarif;Pajak Terutang;Status Validasi\n"
	for _, it := range list {
		csvContent += fmt.Sprintf("%s;%s;%s;%s;%.0f;%.1f%%;%.0f;VALID\n",
			it.TaxPeriod, it.TaxType, it.NPWP, it.PartnerName, it.TaxBase, it.TaxRate, it.TaxAmount)
	}
	return csvContent, nil
}
