package forum

func CreateForumPostService(data *ForumPost) error {
	return CreateForumPost(data)
}

func GetAllForumPostService() ([]ForumPost, error) {
	return GetAllForumPost()
}

func GetPaginatedForumPostService(offset, limit int, search string) ([]ForumPost, int64, error) {
	return GetPaginatedForumPosts(offset, limit, search)
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
