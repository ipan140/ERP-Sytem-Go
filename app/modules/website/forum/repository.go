package forum

import (
	"ERP-System/config"
)

func CreateForumPost(data *ForumPost) error {
	return config.DB.Create(data).Error
}

func GetAllForumPost() ([]ForumPost, error) {
	var list []ForumPost
	err := config.DB.Find(&list).Error
	return list, err
}

func GetForumPostByID(id uint) (*ForumPost, error) {
	var data ForumPost
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateForumPost(data *ForumPost) error {
	return config.DB.Save(data).Error
}

func DeleteForumPost(id uint) error {
	return config.DB.Delete(&ForumPost{}, id).Error
}
