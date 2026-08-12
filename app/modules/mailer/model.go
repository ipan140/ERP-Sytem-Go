package mailer

import (
	"ERP-System/config"
	"time"
)

// EmailLog mencatat riwayat semua email yang pernah dikirim oleh sistem
type EmailLog struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Recipient string     `gorm:"type:varchar(255);not null" json:"recipient"`
	Subject   string     `gorm:"type:varchar(255);not null" json:"subject"`
	Body      string     `gorm:"type:text" json:"body"`
	Status    string     `gorm:"type:varchar(50);default:'pending'" json:"status"`
	SentAt    *time.Time `json:"sent_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &EmailLog{})
}
