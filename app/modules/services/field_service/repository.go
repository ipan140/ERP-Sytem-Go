package field_service

import (
	"ERP-System/config"
	"time"
	"gorm.io/gorm/clause"
)

func CreateFieldServiceTask(data *FieldServiceTask) error {
	return config.DB.Create(data).Error
}

func GetAllFieldServiceTask() ([]FieldServiceTask, error) {
	var list []FieldServiceTask
	err := config.DB.Preload(clause.Associations).Order("id DESC").Find(&list).Error
	return list, err
}

func GetPaginatedFieldServiceTasks(offset int, limit int, search string, state string, priority string, employeeID uint, companyID uint) ([]FieldServiceTask, int64, error) {
	var list []FieldServiceTask
	var total int64

	query := config.DB.Model(&FieldServiceTask{}).Preload(clause.Associations)

	if companyID > 0 {
		query = query.Where("company_id = ?", companyID)
	}
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if state != "" && state != "all" {
		query = query.Where("state = ?", state)
	}
	if priority != "" && priority != "all" {
		query = query.Where("priority = ?", priority)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR address ILIKE ? OR notes ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetFieldServiceTaskByID(id uint) (*FieldServiceTask, error) {
	var data FieldServiceTask
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateFieldServiceTask(data *FieldServiceTask) error {
	return config.DB.Save(data).Error
}

func DeleteFieldServiceTask(id uint) error {
	return config.DB.Delete(&FieldServiceTask{}, id).Error
}

// Fase 2: e-BAST Validation Gate
func ValidateBastFieldServiceTask(id uint, validatorID uint) error {
	now := time.Now()
	updates := map[string]interface{}{
		"bast_validated":   true,
		"validated_by_id": validatorID,
		"validated_at":     now,
		"state":            "completed",
	}
	return config.DB.Model(&FieldServiceTask{}).Where("id = ?", id).Updates(updates).Error
}

// Fase 4: GPS Geotagging
func RecordGPSCheckIn(id uint, lat float64, lng float64) (*FieldServiceTask, error) {
	now := time.Now()
	updates := map[string]interface{}{
		"check_in_lat": lat,
		"check_in_lng": lng,
		"check_in_at":  now,
		"state":        "in_progress",
	}
	if err := config.DB.Model(&FieldServiceTask{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return GetFieldServiceTaskByID(id)
}

func RecordGPSCheckOut(id uint, lat float64, lng float64) (*FieldServiceTask, error) {
	now := time.Now()
	updates := map[string]interface{}{
		"check_out_lat": lat,
		"check_out_lng": lng,
		"check_out_at":  now,
	}
	if err := config.DB.Model(&FieldServiceTask{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return GetFieldServiceTaskByID(id)
}

