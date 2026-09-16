package website_builder

func CreatePageService(data *Page) error {
	return CreatePage(data)
}

func GetAllPageService() ([]Page, error) {
	return GetAllPage()
}

func GetPaginatedPageService(offset, limit int, search string) ([]Page, int64, error) {
	return GetPaginatedPages(offset, limit, search)
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
