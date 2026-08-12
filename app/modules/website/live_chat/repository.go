package live_chat

import (
	"ERP-System/config"
)

func CreateChatSession(data *ChatSession) error {
	return config.DB.Create(data).Error
}

func GetAllChatSession() ([]ChatSession, error) {
	var list []ChatSession
	err := config.DB.Find(&list).Error
	return list, err
}

func GetChatSessionByID(id uint) (*ChatSession, error) {
	var data ChatSession
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateChatSession(data *ChatSession) error {
	return config.DB.Save(data).Error
}

func DeleteChatSession(id uint) error {
	return config.DB.Delete(&ChatSession{}, id).Error
}
