package notifications

func GetNotificationLogsService(limit int) ([]NotificationLog, error) {
	return GetNotificationLogs(limit)
}

func DispatchNotificationService(channel string, recipient string, recipientName string, entityType string, entityID *uint, subject string, message string) (*NotificationLog, error) {
	return SendNotification(nil, channel, recipient, recipientName, entityType, entityID, subject, message)
}
