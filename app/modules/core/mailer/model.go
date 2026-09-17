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

func (EmailLog) TableName() string {
	return "setting.email_logs"
}

// SmtpConfig menyimpan kredensial dan konfigurasi SMTP outgoing mail server
type SmtpConfig struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Host       string    `gorm:"type:varchar(255);not null;default:'smtp.gmail.com'" json:"host"`
	Port       int       `gorm:"default:587" json:"port"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	SenderName string    `gorm:"type:varchar(255);default:'ERP Notification System'" json:"sender_name"`
	Encryption string    `gorm:"type:varchar(50);default:'STARTTLS'" json:"encryption"` // STARTTLS, SSL_TLS, NONE
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (SmtpConfig) TableName() string {
	return "setting.smtp_configs"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &EmailLog{}, &SmtpConfig{})
}
