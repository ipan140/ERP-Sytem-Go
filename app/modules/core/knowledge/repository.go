package knowledge

import (
	"ERP-System/config"
)

func CreateArticle(data *Article) error {
	return config.DB.Create(data).Error
}

func GetAllArticle() ([]Article, error) {
	var list []Article
	err := config.DB.Find(&list).Error
	return list, err
}

func GetArticleByID(id uint) (*Article, error) {
	var data Article
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateArticle(data *Article) error {
	return config.DB.Save(data).Error
}

func DeleteArticle(id uint) error {
	return config.DB.Delete(&Article{}, id).Error
}
