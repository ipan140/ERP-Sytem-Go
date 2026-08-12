package blog

func CreateBlogPostService(data *BlogPost) error {
	return CreateBlogPost(data)
}

func GetAllBlogPostService() ([]BlogPost, error) {
	return GetAllBlogPost()
}

func GetBlogPostByIDService(id uint) (*BlogPost, error) {
	return GetBlogPostByID(id)
}

func UpdateBlogPostService(data *BlogPost) error {
	return UpdateBlogPost(data)
}

func DeleteBlogPostService(id uint) error {
	return DeleteBlogPost(id)
}
