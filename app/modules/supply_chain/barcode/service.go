package barcode

func CreateBarcodeConfigService(data *BarcodeConfig) error {
	return CreateBarcodeConfig(data)
}

func GetAllBarcodeConfigService() ([]BarcodeConfig, error) {
	return GetAllBarcodeConfig()
}

func GetBarcodeConfigByIDService(id uint) (*BarcodeConfig, error) {
	return GetBarcodeConfigByID(id)
}

func UpdateBarcodeConfigService(data *BarcodeConfig) error {
	return UpdateBarcodeConfig(data)
}

func DeleteBarcodeConfigService(id uint) error {
	return DeleteBarcodeConfig(id)
}
