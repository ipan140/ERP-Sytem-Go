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
