package blog

import (
	"ERP-System/config"
)

func CreateBlogPost(data *BlogPost) error {
	return config.DB.Create(data).Error
}

func GetAllBlogPost() ([]BlogPost, error) {
	var list []BlogPost
	err := config.DB.Find(&list).Error
	return list, err
}

func GetBlogPostByID(id uint) (*BlogPost, error) {
	var data BlogPost
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateBlogPost(data *BlogPost) error {
	return config.DB.Save(data).Error
}

func DeleteBlogPost(id uint) error {
	return config.DB.Delete(&BlogPost{}, id).Error
}
