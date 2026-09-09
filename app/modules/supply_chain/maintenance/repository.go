package maintenance

import (
	"ERP-System/config"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaintenanceRequest(data *MaintenanceRequest) error {
	return config.DB.Create(data).Error
}

func GetAllMaintenanceRequest() ([]MaintenanceRequest, error) {
	var list []MaintenanceRequest
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetMaintenanceRequestByID(id uint) (*MaintenanceRequest, error) {
	var data MaintenanceRequest
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateMaintenanceRequest(data *MaintenanceRequest) error {
	return config.DB.Save(data).Error
}

func DeleteMaintenanceRequest(id uint) error {
	return config.DB.Delete(&MaintenanceRequest{}, id).Error
}

func GetPaginatedMaintenanceRequests(page, limit int, search, state, reqType string) ([]MaintenanceRequest, int64, error) {
	var list []MaintenanceRequest
	var total int64

	db := config.DB.Model(&MaintenanceRequest{})

	if search != "" {
		s := "%" + strings.ToLower(search) + "%"
		db = db.Where("LOWER(name) LIKE ? OR LOWER(code) LIKE ?", s, s)
	}

	if state != "" && state != "all" {
		db = db.Where("state = ?", state)
	}

	if reqType != "" && reqType != "all" {
		db = db.Where("type = ?", reqType)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Preload(clause.Associations).
		Order("id desc").
		Offset(offset).
		Limit(limit).
		Find(&list).Error

	return list, total, err
}

func GetMaintenanceSummary() (*MaintenanceSummary, error) {
	var summary MaintenanceSummary

	config.DB.Model(&MaintenanceRequest{}).Count(&summary.TotalRequests)
	config.DB.Model(&MaintenanceRequest{}).Where("state = ?", "todo").Count(&summary.TodoCount)
	config.DB.Model(&MaintenanceRequest{}).Where("state = ?", "progress").Count(&summary.InProgressCount)
	config.DB.Model(&MaintenanceRequest{}).Where("state = ?", "done").Count(&summary.DoneCount)
	config.DB.Model(&MaintenanceEquipment{}).Count(&summary.TotalEquipments)

	return &summary, nil
}

func CreateMaintenanceRequestWithSequence(req *CreateMaintenanceRequestDto) (*MaintenanceRequest, error) {
	var mr MaintenanceRequest
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		dateStr := time.Now().Format("20060102")
		prefix := fmt.Sprintf("MR-%s-", dateStr)

		var last MaintenanceRequest
		var nextNum int64 = 1
		if err := tx.Unscoped().Where("code LIKE ?", prefix+"%").Order("code desc").First(&last).Error; err == nil {
			var lastNum int64
			fmt.Sscanf(last.Code, prefix+"%d", &lastNum)
			nextNum = lastNum + 1
		} else {
			var count int64
			tx.Model(&MaintenanceRequest{}).Count(&count)
			nextNum = count + 1
		}

		code := fmt.Sprintf("MR-%s-%04d", dateStr, nextNum)

		mr = MaintenanceRequest{
			Name:        req.Name,
			Code:        code,
			EquipmentID: req.EquipmentID,
			Type:        req.Type,
			Priority:    req.Priority,
			State:       "todo",
			Notes:       req.Notes,
		}

		if req.Priority == "" {
			mr.Priority = "normal"
		}
		if req.Type == "" {
			mr.Type = "corrective"
		}

		if req.Duration != nil {
			mr.Duration = *req.Duration
		}

		if req.ScheduleDate != nil && *req.ScheduleDate != "" {
			parsedDate, err := time.Parse(time.RFC3339, *req.ScheduleDate)
			if err != nil {
				parsedDate, err = time.Parse("2006-01-02", *req.ScheduleDate)
			}
			if err == nil {
				mr.ScheduleDate = &parsedDate
			}
		}

		return tx.Create(&mr).Error
	})

	if err != nil {
		return nil, err
	}

	// Preload relations
	config.DB.Preload(clause.Associations).First(&mr, mr.ID)
	return &mr, nil
}

func UpdateMaintenanceState(id uint, state string, duration *float64, notes *string) (*MaintenanceRequest, error) {
	var mr MaintenanceRequest
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&mr, id).Error; err != nil {
			return err
		}

		mr.State = state
		if duration != nil {
			mr.Duration = *duration
		}
		if notes != nil && *notes != "" {
			mr.Notes = *notes
		}
		if state == "done" {
			now := time.Now()
			mr.DateDone = &now
		}

		return tx.Save(&mr).Error
	})

	if err != nil {
		return nil, err
	}

	config.DB.Preload(clause.Associations).First(&mr, mr.ID)
	return &mr, nil
}

func GetPaginatedEquipments(page, limit int, search string) ([]MaintenanceEquipment, int64, error) {
	var list []MaintenanceEquipment
	var total int64

	db := config.DB.Model(&MaintenanceEquipment{})

	if search != "" {
		s := "%" + strings.ToLower(search) + "%"
		db = db.Where("LOWER(name) LIKE ? OR LOWER(category) LIKE ?", s, s)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Preload(clause.Associations).
		Order("id desc").
		Offset(offset).
		Limit(limit).
		Find(&list).Error

	return list, total, err
}

func CreateMaintenanceEquipment(dto *CreateMaintenanceEquipmentDto) (*MaintenanceEquipment, error) {
	eq := MaintenanceEquipment{
		Name:         dto.Name,
		Category:     dto.Category,
		WorkcenterID: dto.WorkcenterID,
		Cost:         dto.Cost,
	}

	if dto.NextActionDate != nil && *dto.NextActionDate != "" {
		parsedDate, err := time.Parse(time.RFC3339, *dto.NextActionDate)
		if err != nil {
			parsedDate, err = time.Parse("2006-01-02", *dto.NextActionDate)
		}
		if err == nil {
			eq.NextActionDate = &parsedDate
		}
	}

	if err := config.DB.Create(&eq).Error; err != nil {
		return nil, err
	}

	config.DB.Preload(clause.Associations).First(&eq, eq.ID)
	return &eq, nil
}

func DeleteMaintenanceEquipment(id uint) error {
	return config.DB.Delete(&MaintenanceEquipment{}, id).Error
}
