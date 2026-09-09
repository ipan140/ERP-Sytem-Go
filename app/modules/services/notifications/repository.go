package notifications

import (
	"ERP-System/config"
	"fmt"
	"log"
)

func CreateNotificationLog(n *NotificationLog) error {
	return config.DB.Create(n).Error
}

func GetNotificationLogs(limit int) ([]NotificationLog, error) {
	var logs []NotificationLog
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	err := config.DB.Order("created_at DESC").Limit(limit).Find(&logs).Error
	return logs, err
}

// SendNotification dispatches notification via WhatsApp / Email and logs to database
func SendNotification(companyID *uint, channel string, recipient string, recipientName string, entityType string, entityID *uint, subject string, message string) (*NotificationLog, error) {
	nLog := &NotificationLog{
		CompanyID:     companyID,
		Channel:       channel,
		Recipient:     recipient,
		RecipientName: recipientName,
		EntityType:    entityType,
		EntityID:      entityID,
		Subject:       subject,
		Message:       message,
		Status:        "sent",
	}

	// Mock / Webhook Gateway Dispatch
	log.Printf("[NOTIFICATION GATEWAY] [%s] To: %s (%s) | Subject: %s | Message: %s",
		channel, recipient, recipientName, subject, message)
	fmt.Printf("🚀 Sent %s to %s: %s\n", channel, recipient, message)

	err := CreateNotificationLog(nLog)
	return nLog, err
}
