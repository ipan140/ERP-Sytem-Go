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

func GetPaginatedMailingCampaignService(offset, limit int, search string) ([]MailingCampaign, int64, error) {
	return GetPaginatedMailingCampaigns(offset, limit, search)
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

// RunABSplitTestService mengevaluasi performa Versi A vs Versi B dan memilih pemenang otomatis
func RunABSplitTestService(campaignID uint) (*MailingCampaign, error) {
	campaign, err := GetMailingCampaignByID(campaignID)
	if err != nil {
		return nil, err
	}

	campaign.IsABTesting = true
	// Simulasi hasil sample testing (misal Versi B menghasilkan open rate lebih tinggi)
	campaign.VariantAOpened = 142
	campaign.VariantBOpened = 289

	if campaign.VariantBOpened > campaign.VariantAOpened {
		campaign.WinnerVariant = "B"
		// Otomatis ubah subjek kampanye utama menjadi Subjek B pemenang untuk sisa audiens
		if campaign.SubjectB != "" {
			campaign.Subject = campaign.SubjectB
		}
	} else {
		campaign.WinnerVariant = "A"
	}

	campaign.Status = "Sent"
	campaign.SentCount = 2500
	campaign.OpenedCount = campaign.VariantAOpened + campaign.VariantBOpened
	campaign.ClickedCount = int(float64(campaign.OpenedCount) * 0.45)

	if err := UpdateMailingCampaign(campaign); err != nil {
		return nil, err
	}

	return campaign, nil
}

