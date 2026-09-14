package approvals

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateApprovalRequest(data *ApprovalRequest) error {
	return config.DB.Create(data).Error
}

func GetAllApprovalRequest() ([]ApprovalRequest, error) {
	var list []ApprovalRequest
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedApprovalRequests(offset int, limit int, search string, status string) ([]ApprovalRequest, int64, error) {
	var list []ApprovalRequest
	var total int64

	query := config.DB.Model(&ApprovalRequest{})

	if status != "" && status != "all" && status != "All" && status != "Semua" {
		query = query.Where("status = ?", status)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR requester_name ILIKE ? OR type ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetApprovalRequestByID(id uint) (*ApprovalRequest, error) {
	var data ApprovalRequest
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateApprovalRequest(data *ApprovalRequest) error {
	return config.DB.Save(data).Error
}

func DeleteApprovalRequest(id uint) error {
	return config.DB.Delete(&ApprovalRequest{}, id).Error
}
