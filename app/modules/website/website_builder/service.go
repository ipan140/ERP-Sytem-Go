package website_builder

func CreatePageService(data *Page) error {
	return CreatePage(data)
}

func GetAllPageService() ([]Page, error) {
	return GetAllPage()
}

func GetPageByIDService(id uint) (*Page, error) {
	return GetPageByID(id)
}

func UpdatePageService(data *Page) error {
	return UpdatePage(data)
}

func DeletePageService(id uint) error {
	return DeletePage(id)
}
