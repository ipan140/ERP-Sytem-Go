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

func CreateCountryService(data *Country) error        { return CreateCountry(data) }
func GetAllCountryService() ([]Country, error)        { return GetAllCountry() }
func GetCountryByIDService(id uint) (*Country, error) { return GetCountryByID(id) }
func UpdateCountryService(data *Country) error        { return UpdateCountry(data) }
func DeleteCountryService(id uint) error              { return DeleteCountry(id) }

func CreateCountryStateService(data *CountryState) error        { return CreateCountryState(data) }
func GetAllCountryStateService() ([]CountryState, error)        { return GetAllCountryState() }
func GetCountryStateByIDService(id uint) (*CountryState, error) { return GetCountryStateByID(id) }
func UpdateCountryStateService(data *CountryState) error        { return UpdateCountryState(data) }
func DeleteCountryStateService(id uint) error                   { return DeleteCountryState(id) }

func CreatePartnerService(data *Partner) error        { return CreatePartner(data) }
func GetAllPartnerService() ([]Partner, error)        { return GetAllPartner() }
func GetPaginatedPartnerService(offset, limit int, search string) ([]Partner, int64, error) {
	return GetPaginatedPartner(offset, limit, search)
}
func GetPartnerByIDService(id uint) (*Partner, error) { return GetPartnerByID(id) }
func UpdatePartnerService(data *Partner) error        { return UpdatePartner(data) }
func DeletePartnerService(id uint) error              { return DeletePartner(id) }
