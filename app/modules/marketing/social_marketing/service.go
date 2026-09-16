package social_marketing

func CreateSocialPostService(data *SocialPost) error {
	return CreateSocialPost(data)
}

func GetAllSocialPostService() ([]SocialPost, error) {
	return GetAllSocialPost()
}

func GetPaginatedSocialPostService(offset, limit int, search string) ([]SocialPost, int64, error) {
	return GetPaginatedSocialPosts(offset, limit, search)
}

func GetSocialPostByIDService(id uint) (*SocialPost, error) {
	return GetSocialPostByID(id)
}

func UpdateSocialPostService(data *SocialPost) error {
	return UpdateSocialPost(data)
}

func DeleteSocialPostService(id uint) error {
	return DeleteSocialPost(id)
}
