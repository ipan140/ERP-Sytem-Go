package activity_logs

func GetActivityLogsService(entityType string, entityID uint, limit int) ([]ActivityLog, error) {
	return GetActivityLogsByEntity(entityType, entityID, limit)
}
