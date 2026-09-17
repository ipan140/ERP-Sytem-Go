package base

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateCurrency(data *Currency) error {
	return config.DB.Create(data).Error
}

func GetAllCurrency() ([]Currency, error) {
	var list []Currency
	err := config.DB.Preload(clause.Associations).Order("is_base desc, id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SyncBankIndonesiaRates()
		_ = config.DB.Order("is_base desc, id asc").Find(&list).Error
	}
	return list, err
}

func SyncBankIndonesiaRates() error {
	defaults := []Currency{
		{Code: "IDR", Name: "Indonesian Rupiah", Symbol: "Rp", Rate: 1.00, IsBase: true},
		{Code: "USD", Name: "US Dollar", Symbol: "$", Rate: 16250.00, IsBase: false},
		{Code: "EUR", Name: "Euro", Symbol: "€", Rate: 17680.00, IsBase: false},
		{Code: "SGD", Name: "Singapore Dollar", Symbol: "S$", Rate: 12540.00, IsBase: false},
		{Code: "JPY", Name: "Japanese Yen", Symbol: "¥", Rate: 104.20, IsBase: false},
		{Code: "GBP", Name: "British Pound Sterling", Symbol: "£", Rate: 21100.00, IsBase: false},
		{Code: "AUD", Name: "Australian Dollar", Symbol: "A$", Rate: 10750.00, IsBase: false},
		{Code: "MYR", Name: "Malaysian Ringgit", Symbol: "RM", Rate: 3740.00, IsBase: false},
		{Code: "CNY", Name: "Chinese Yuan", Symbol: "¥", Rate: 2270.00, IsBase: false},
	}
	for _, d := range defaults {
		var existing Currency
		if err := config.DB.Where("code = ?", d.Code).First(&existing).Error; err == nil {
			existing.Rate = d.Rate
			existing.Symbol = d.Symbol
			existing.Name = d.Name
			existing.IsBase = d.IsBase
			config.DB.Save(&existing)
		} else {
			config.DB.Create(&d)
		}
	}
	return nil
}

func GetCurrencyByID(id uint) (*Currency, error) {
	var data Currency
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateCurrency(data *Currency) error {
	return config.DB.Save(data).Error
}

func DeleteCurrency(id uint) error {
	return config.DB.Delete(&Currency{}, id).Error
}

func CreateCountry(data *Country) error { return config.DB.Create(data).Error }
func GetAllCountry() ([]Country, error) {
	var list []Country
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetCountryByID(id uint) (*Country, error) {
	var data Country
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateCountry(data *Country) error { return config.DB.Save(data).Error }
func DeleteCountry(id uint) error       { return config.DB.Delete(&Country{}, id).Error }

func CreateCountryState(data *CountryState) error { return config.DB.Create(data).Error }
func GetAllCountryState() ([]CountryState, error) {
	var list []CountryState
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetCountryStateByID(id uint) (*CountryState, error) {
	var data CountryState
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateCountryState(data *CountryState) error { return config.DB.Save(data).Error }
func DeleteCountryState(id uint) error            { return config.DB.Delete(&CountryState{}, id).Error }

func CreatePartner(data *Partner) error { return config.DB.Create(data).Error }
func GetAllPartner() ([]Partner, error) {
	var list []Partner
	err := config.DB.Preload(clause.Associations).Order("id asc").Find(&list).Error
	return list, err
}
func GetPaginatedPartner(offset, limit int, search string) ([]Partner, int64, error) {
	var list []Partner
	var total int64
	db := config.DB.Model(&Partner{})
	if search != "" {
		db = db.Where("name ILIKE ? OR email ILIKE ? OR phone ILIKE ? OR city ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	db.Count(&total)
	err := db.Preload(clause.Associations).Order("id asc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}
func GetPartnerByID(id uint) (*Partner, error) {
	var data Partner
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePartner(data *Partner) error { return config.DB.Save(data).Error }
func DeletePartner(id uint) error       { return config.DB.Delete(&Partner{}, id).Error }
