package base

func CreateCurrencyService(data *Currency) error {
	return CreateCurrency(data)
}

func GetAllCurrencyService() ([]Currency, error) {
	return GetAllCurrency()
}

func GetCurrencyByIDService(id uint) (*Currency, error) {
	return GetCurrencyByID(id)
}

func UpdateCurrencyService(data *Currency) error {
	return UpdateCurrency(data)
}

func DeleteCurrencyService(id uint) error {
	return DeleteCurrency(id)
}
