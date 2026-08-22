package mass_mailing

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

func CreateMailingCampaignService(data *MailingCampaign) error {
	err := CreateMailingCampaign(data)
	if err == nil && rabbitmq.Channel != nil {
		body, _ := json.Marshal(data)
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "marketing_broadcast", body)
		log.Println("📢 Event RabbitMQ: Kampanye Email Massal dikirim ke antrean!")
	}
	return err
}

func GetAllMailingCampaignService() ([]MailingCampaign, error) {
	return GetAllMailingCampaign()
}

func GetMailingCampaignByIDService(id uint) (*MailingCampaign, error) {
	return GetMailingCampaignByID(id)
}

func UpdateMailingCampaignService(data *MailingCampaign) error {
	return UpdateMailingCampaign(data)
}

func DeleteMailingCampaignService(id uint) error {
	return DeleteMailingCampaign(id)
}

func CreateUtmTrackerService(data *UtmTracker) error        { return CreateUtmTracker(data) }
func GetAllUtmTrackerService() ([]UtmTracker, error)        { return GetAllUtmTracker() }
func GetUtmTrackerByIDService(id uint) (*UtmTracker, error) { return GetUtmTrackerByID(id) }
func UpdateUtmTrackerService(data *UtmTracker) error        { return UpdateUtmTracker(data) }
func DeleteUtmTrackerService(id uint) error                 { return DeleteUtmTracker(id) }
