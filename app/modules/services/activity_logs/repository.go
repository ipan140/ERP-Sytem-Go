package activity_logs

import (
	"ERP-System/config"
	"encoding/json"
)

func CreateActivityLog(log *ActivityLog) error {
	return config.DB.Create(log).Error
}

func GetActivityLogsByEntity(entityType string, entityID uint, limit int) ([]ActivityLog, error) {
	var logs []ActivityLog
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	err := config.DB.Where("entity_type = ? AND entity_id = ?", entityType, entityID).
		Order("created_at DESC").
		Limit(limit).
		Find(&logs).Error
	return logs, err
}

// Record is a helper to record activity easily from any service
func Record(companyID *uint, entityType string, entityID uint, action string, userID *uint, userName string, oldVal any, newVal any, notes string, ip string) error {
	var oldJson, newJson string
	if oldVal != nil {
		if b, err := json.Marshal(oldVal); err == nil {
			oldJson = string(b)
		}
	}
	if newVal != nil {
		if b, err := json.Marshal(newVal); err == nil {
			newJson = string(b)
		}
	}

	log := &ActivityLog{
		CompanyID:  companyID,
		EntityType: entityType,
		EntityID:   entityID,
		Action:     action,
		UserID:     userID,
		UserName:   userName,
		OldValue:   oldJson,
		NewValue:   newJson,
		Notes:      notes,
		IPAddress:  ip,
	}
	return CreateActivityLog(log)
}
