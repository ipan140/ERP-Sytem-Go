package blog

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateBlogPost(data *BlogPost) error {
	return config.DB.Create(data).Error
}

func GetAllBlogPost() ([]BlogPost, error) {
	var list []BlogPost
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedBlogPosts(offset, limit int, search string) ([]BlogPost, int64, error) {
	var list []BlogPost
	var total int64
	query := config.DB.Model(&BlogPost{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("title ILIKE ? OR state ILIKE ?", s, s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetBlogPostByID(id uint) (*BlogPost, error) {
	var data BlogPost
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateBlogPost(data *BlogPost) error {
	return config.DB.Save(data).Error
}

func DeleteBlogPost(id uint) error {
	return config.DB.Delete(&BlogPost{}, id).Error
}
