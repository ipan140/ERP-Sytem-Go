package forum

func CreateForumPostService(data *ForumPost) error {
	return CreateForumPost(data)
}

func GetAllForumPostService() ([]ForumPost, error) {
	return GetAllForumPost()
}

func GetForumPostByIDService(id uint) (*ForumPost, error) {
	return GetForumPostByID(id)
}

func UpdateForumPostService(data *ForumPost) error {
	return UpdateForumPost(data)
}

func DeleteForumPostService(id uint) error {
	return DeleteForumPost(id)
}
