package mass_mailing

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateMailingCampaign(data *MailingCampaign) error {
	return config.DB.Create(data).Error
}

func GetAllMailingCampaign() ([]MailingCampaign, error) {
	var list []MailingCampaign
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedMailingCampaigns(offset, limit int, search string) ([]MailingCampaign, int64, error) {
	var list []MailingCampaign
	var total int64
	query := config.DB.Model(&MailingCampaign{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR subject ILIKE ? OR target_audience ILIKE ?", s, s, s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetMailingCampaignByID(id uint) (*MailingCampaign, error) {
	var data MailingCampaign
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateMailingCampaign(data *MailingCampaign) error {
	return config.DB.Save(data).Error
}

func DeleteMailingCampaign(id uint) error {
	return config.DB.Delete(&MailingCampaign{}, id).Error
}

func CreateUtmTracker(data *UtmTracker) error { return config.DB.Create(data).Error }
func GetAllUtmTracker() ([]UtmTracker, error) {
	var list []UtmTracker
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetUtmTrackerByID(id uint) (*UtmTracker, error) {
	var data UtmTracker
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateUtmTracker(data *UtmTracker) error { return config.DB.Save(data).Error }
func DeleteUtmTracker(id uint) error          { return config.DB.Delete(&UtmTracker{}, id).Error }
