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

func CreateLeaveType(data *LeaveType) error { return config.DB.Create(data).Error }
func GetAllLeaveType() ([]LeaveType, error) {
	var list []LeaveType
	err := config.DB.Find(&list).Error
	return list, err
}
func GetLeaveTypeByID(id uint) (*LeaveType, error) {
	var data LeaveType
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateLeaveType(data *LeaveType) error { return config.DB.Save(data).Error }
func DeleteLeaveType(id uint) error         { return config.DB.Delete(&LeaveType{}, id).Error }

func CreateLeaveAllocation(data *LeaveAllocation) error { return config.DB.Create(data).Error }
func GetAllLeaveAllocation() ([]LeaveAllocation, error) {
	var list []LeaveAllocation
	err := config.DB.Find(&list).Error
	return list, err
}
func GetLeaveAllocationByID(id uint) (*LeaveAllocation, error) {
	var data LeaveAllocation
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateLeaveAllocation(data *LeaveAllocation) error { return config.DB.Save(data).Error }
func DeleteLeaveAllocation(id uint) error               { return config.DB.Delete(&LeaveAllocation{}, id).Error }
