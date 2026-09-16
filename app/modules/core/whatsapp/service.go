package whatsapp

func CreateWaTemplateService(data *WaTemplate) error {
	return CreateWaTemplate(data)
}

func GetAllWaTemplateService() ([]WaTemplate, error) {
	return GetAllWaTemplate()
}

func GetPaginatedWaTemplateService(offset, limit int, search string) ([]WaTemplate, int64, error) {
	return GetPaginatedWaTemplate(offset, limit, search)
}

func GetWaTemplateByIDService(id uint) (*WaTemplate, error) {
	return GetWaTemplateByID(id)
}

func UpdateWaTemplateService(data *WaTemplate) error {
	return UpdateWaTemplate(data)
}

func DeleteWaTemplateService(id uint) error {
	return DeleteWaTemplate(id)
}
