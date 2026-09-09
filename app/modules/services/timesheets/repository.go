package timesheets

import (
	"ERP-System/config"
	"time"
	"gorm.io/gorm/clause"
)

func CreateTimesheet(data *Timesheet) error {
	return config.DB.Create(data).Error
}

func GetAllTimesheet() ([]Timesheet, error) {
	var list []Timesheet
	err := config.DB.Preload(clause.Associations).Order("date DESC, id DESC").Find(&list).Error
	return list, err
}

func GetPaginatedTimesheets(offset int, limit int, search string, projectID uint, employeeID uint, billable string, status string, companyID uint) ([]Timesheet, int64, error) {
	var list []Timesheet
	var total int64

	query := config.DB.Model(&Timesheet{}).Preload(clause.Associations)

	if companyID > 0 {
		query = query.Where("company_id = ?", companyID)
	}
	if projectID > 0 {
		query = query.Where("project_id = ?", projectID)
	}
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}
	if billable == "true" {
		query = query.Where("is_billable = ?", true)
	} else if billable == "false" {
		query = query.Where("is_billable = ?", false)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("description ILIKE ?", s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("date DESC, id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetTimesheetByID(id uint) (*Timesheet, error) {
	var data Timesheet
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateTimesheet(data *Timesheet) error {
	return config.DB.Save(data).Error
}

func DeleteTimesheet(id uint) error {
	return config.DB.Delete(&Timesheet{}, id).Error
}

// Fase 2: Workflow Approval Methods
func UpdateTimesheetStatus(id uint, updates map[string]interface{}) error {
	return config.DB.Model(&Timesheet{}).Where("id = ?", id).Updates(updates).Error
}

func BulkApproveTimesheets(ids []uint, approverID uint) error {
	now := time.Now()
	updates := map[string]interface{}{
		"status":          "approved",
		"approved_by_id": approverID,
		"approved_at":     now,
		"rejection_reason": "",
	}
	return config.DB.Model(&Timesheet{}).Where("id IN ?", ids).Updates(updates).Error
}

