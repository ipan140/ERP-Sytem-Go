package repairs

import (
	"ERP-System/config"
	"time"
	"gorm.io/gorm/clause"
)

func CreateRepairOrder(data *RepairOrder) error {
	return config.DB.Create(data).Error
}

func GetAllRepairOrder() ([]RepairOrder, error) {
	var list []RepairOrder
	err := config.DB.Preload(clause.Associations).Order("id DESC").Find(&list).Error
	return list, err
}

func GetPaginatedRepairOrders(offset int, limit int, search string, state string, warranty string, companyID uint, technicianID uint) ([]RepairOrder, int64, error) {
	var list []RepairOrder
	var total int64

	query := config.DB.Model(&RepairOrder{}).Preload(clause.Associations)

	if companyID > 0 {
		query = query.Where("company_id = ?", companyID)
	}
	if technicianID > 0 {
		query = query.Where("technician_id = ?", technicianID)
	}
	if state != "" && state != "all" {
		query = query.Where("state = ?", state)
	}
	if warranty != "" && warranty != "all" {
		query = query.Where("warranty_status = ?", warranty)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR serial_number ILIKE ? OR diagnosis ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetRepairOrderByID(id uint) (*RepairOrder, error) {
	var data RepairOrder
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateRepairOrder(data *RepairOrder) error {
	return config.DB.Save(data).Error
}

func DeleteRepairOrder(id uint) error {
	return config.DB.Delete(&RepairOrder{}, id).Error
}

// Fase 2: Quality Control Gate
func PassQCRepairOrder(id uint, inspectorID uint, notes string) error {
	now := time.Now()
	updates := map[string]interface{}{
		"qc_passed":        true,
		"qc_inspector_id": inspectorID,
		"qc_passed_at":    now,
		"qc_notes":        notes,
		"state":           "ready",
	}
	return config.DB.Model(&RepairOrder{}).Where("id = ?", id).Updates(updates).Error
}

// Fase 4: Public Tracking & Customer Approval
func GetRepairByRMAAndSerial(id uint, serialNumber string) (*RepairOrder, error) {
	var data RepairOrder
	query := config.DB.Preload(clause.Associations).Where("id = ?", id)
	if serialNumber != "" {
		query = query.Where("LOWER(TRIM(serial_number)) = LOWER(TRIM(?))", serialNumber)
	}
	err := query.First(&data).Error
	return &data, err
}

func ApproveRepairEstimate(id uint, note string) (*RepairOrder, error) {
	now := time.Now()
	updates := map[string]interface{}{
		"customer_approved_at":   now,
		"customer_approval_note": note,
		"state":                  "under_repair",
	}
	if err := config.DB.Model(&RepairOrder{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return GetRepairOrderByID(id)
}

