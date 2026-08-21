package base

import (
	"ERP-System/config"
)

func CreateCurrency(data *Currency) error {
	return config.DB.Create(data).Error
}

func GetAllCurrency() ([]Currency, error) {
	var list []Currency
	err := config.DB.Find(&list).Error
	return list, err
}

func GetCurrencyByID(id uint) (*Currency, error) {
	var data Currency
	err := config.DB.First(&data, id).Error
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
	err := config.DB.Find(&list).Error
	return list, err
}
func GetCountryByID(id uint) (*Country, error) {
	var data Country
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateCountry(data *Country) error { return config.DB.Save(data).Error }
func DeleteCountry(id uint) error       { return config.DB.Delete(&Country{}, id).Error }

func CreateCountryState(data *CountryState) error { return config.DB.Create(data).Error }
func GetAllCountryState() ([]CountryState, error) {
	var list []CountryState
	err := config.DB.Find(&list).Error
	return list, err
}
func GetCountryStateByID(id uint) (*CountryState, error) {
	var data CountryState
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateCountryState(data *CountryState) error { return config.DB.Save(data).Error }
func DeleteCountryState(id uint) error            { return config.DB.Delete(&CountryState{}, id).Error }

func CreatePartner(data *Partner) error { return config.DB.Create(data).Error }
func GetAllPartner() ([]Partner, error) {
	var list []Partner
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPartnerByID(id uint) (*Partner, error) {
	var data Partner
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePartner(data *Partner) error { return config.DB.Save(data).Error }
func DeletePartner(id uint) error       { return config.DB.Delete(&Partner{}, id).Error }
