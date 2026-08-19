package barcode

func CreateBarcodeConfigService(data *BarcodeNomenclature) error {
	return CreateBarcodeConfig(data)
}

func GetAllBarcodeConfigService() ([]BarcodeNomenclature, error) {
	return GetAllBarcodeConfig()
}

func GetBarcodeConfigByIDService(id uint) (*BarcodeNomenclature, error) {
	return GetBarcodeConfigByID(id)
}

func UpdateBarcodeConfigService(data *BarcodeNomenclature) error {
	return UpdateBarcodeConfig(data)
}

func DeleteBarcodeConfigService(id uint) error {
	return DeleteBarcodeConfig(id)
}
