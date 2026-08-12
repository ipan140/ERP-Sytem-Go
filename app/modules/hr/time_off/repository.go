package time_off

import (
	"ERP-System/config"
)

func CreateLeaveRequest(data *LeaveRequest) error {
	return config.DB.Create(data).Error
}

func GetAllLeaveRequest() ([]LeaveRequest, error) {
	var list []LeaveRequest
	err := config.DB.Find(&list).Error
	return list, err
}

func GetLeaveRequestByID(id uint) (*LeaveRequest, error) {
	var data LeaveRequest
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateLeaveRequest(data *LeaveRequest) error {
	return config.DB.Save(data).Error
}

func DeleteLeaveRequest(id uint) error {
	return config.DB.Delete(&LeaveRequest{}, id).Error
}
