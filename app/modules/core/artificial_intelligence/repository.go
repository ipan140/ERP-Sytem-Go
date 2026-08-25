package artificial_intelligence

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAIPrompt(data *AIPrompt) error {
	return config.DB.Create(data).Error
}

func GetAllAIPrompt() ([]AIPrompt, error) {
	var list []AIPrompt
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetAIPromptByID(id uint) (*AIPrompt, error) {
	var data AIPrompt
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAIPrompt(data *AIPrompt) error {
	return config.DB.Save(data).Error
}

func DeleteAIPrompt(id uint) error {
	return config.DB.Delete(&AIPrompt{}, id).Error
}
