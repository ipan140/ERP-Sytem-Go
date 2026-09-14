package time_off

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateLeaveRequest(data *LeaveRequest) error {
	return config.DB.Create(data).Error
}

func GetAllLeaveRequest() ([]LeaveRequest, error) {
	var list []LeaveRequest
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedLeaveRequests(offset, limit int, search, employeeID, status string) ([]LeaveRequest, int64, error) {
	var list []LeaveRequest
	var total int64

	query := config.DB.Model(&LeaveRequest{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.leave_requests.employee_id = ?", employeeID)
	}

	if status != "" && status != "all" {
		query = query.Where("hrd.leave_requests.status = ?", status)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.leave_requests.employee_id").
			Where("hrd.employees.name ILIKE ? OR hrd.leave_requests.status ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.leave_requests.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetLeaveRequestByID(id uint) (*LeaveRequest, error) {
	var data LeaveRequest
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
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
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetLeaveTypeByID(id uint) (*LeaveType, error) {
	var data LeaveType
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateLeaveType(data *LeaveType) error { return config.DB.Save(data).Error }
func DeleteLeaveType(id uint) error         { return config.DB.Delete(&LeaveType{}, id).Error }

func CreateLeaveAllocation(data *LeaveAllocation) error { return config.DB.Create(data).Error }
func GetAllLeaveAllocation() ([]LeaveAllocation, error) {
	var list []LeaveAllocation
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetLeaveAllocationByID(id uint) (*LeaveAllocation, error) {
	var data LeaveAllocation
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateLeaveAllocation(data *LeaveAllocation) error { return config.DB.Save(data).Error }
func DeleteLeaveAllocation(id uint) error               { return config.DB.Delete(&LeaveAllocation{}, id).Error }
