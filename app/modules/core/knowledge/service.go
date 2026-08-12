package knowledge

func CreateArticleService(data *Article) error {
	return CreateArticle(data)
}

func GetAllArticleService() ([]Article, error) {
	return GetAllArticle()
}

func GetArticleByIDService(id uint) (*Article, error) {
	return GetArticleByID(id)
}

func UpdateArticleService(data *Article) error {
	return UpdateArticle(data)
}

func DeleteArticleService(id uint) error {
	return DeleteArticle(id)
}
