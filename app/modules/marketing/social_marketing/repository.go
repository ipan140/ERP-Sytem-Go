package social_marketing

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateSocialPost(data *SocialPost) error {
	return config.DB.Create(data).Error
}

func GetAllSocialPost() ([]SocialPost, error) {
	var list []SocialPost
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedSocialPosts(offset, limit int, search string) ([]SocialPost, int64, error) {
	var list []SocialPost
	var total int64
	query := config.DB.Model(&SocialPost{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("message ILIKE ? OR channels ILIKE ? OR status ILIKE ?", s, s, s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetSocialPostByID(id uint) (*SocialPost, error) {
	var data SocialPost
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateSocialPost(data *SocialPost) error {
	return config.DB.Save(data).Error
}

func DeleteSocialPost(id uint) error {
	return config.DB.Delete(&SocialPost{}, id).Error
}
