package live_chat

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateChatSession(data *ChatSession) error {
	return config.DB.Create(data).Error
}

func GetAllChatSession() ([]ChatSession, error) {
	var list []ChatSession
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedChatSessions(offset, limit int, search string) ([]ChatSession, int64, error) {
	var list []ChatSession
	var total int64
	query := config.DB.Model(&ChatSession{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("visitor_ip ILIKE ?", s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetChatSessionByID(id uint) (*ChatSession, error) {
	var data ChatSession
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateChatSession(data *ChatSession) error {
	return config.DB.Save(data).Error
}

func DeleteChatSession(id uint) error {
	return config.DB.Delete(&ChatSession{}, id).Error
}
