package social_marketing

import (
	"ERP-System/config"
)

func CreateSocialPost(data *SocialPost) error {
	return config.DB.Create(data).Error
}

func GetAllSocialPost() ([]SocialPost, error) {
	var list []SocialPost
	err := config.DB.Find(&list).Error
	return list, err
}

func GetSocialPostByID(id uint) (*SocialPost, error) {
	var data SocialPost
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateSocialPost(data *SocialPost) error {
	return config.DB.Save(data).Error
}

func DeleteSocialPost(id uint) error {
	return config.DB.Delete(&SocialPost{}, id).Error
}
