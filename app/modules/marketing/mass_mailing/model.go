package mass_mailing

import (
	"ERP-System/config"
	"time"
)

type MailingCampaign struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Subject   string    `gorm:"type:varchar(255)" json:"subject"`
	BodyHtml  string    `gorm:"type:text" json:"body_html"`
	State     string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, in_queue, sending, done
	CreatedAt time.Time `json:"created_at"`
}

type MailingContact struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MailingCampaign{}, &MailingContact{})
}
