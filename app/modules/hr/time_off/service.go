package time_off

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

func CreateLeaveRequestService(data *LeaveRequest) error {
	err := CreateLeaveRequest(data)
	if err == nil && rabbitmq.Channel != nil {
		// Asumsi saat karyawan membuat pengajuan cuti pertama kali
		body, _ := json.Marshal(data)
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "hr_leave_notification", body)
		log.Println("📨 Event RabbitMQ: Pengajuan Cuti Baru dikirim ke antrean!")
	}
	return err
}

func GetAllLeaveRequestService() ([]LeaveRequest, error) {
	return GetAllLeaveRequest()
}

func GetLeaveRequestByIDService(id uint) (*LeaveRequest, error) {
	return GetLeaveRequestByID(id)
}

func UpdateLeaveRequestService(data *LeaveRequest) error {
	err := UpdateLeaveRequest(data)
	if err == nil && rabbitmq.Channel != nil {
		// Asumsi saat HR menyetujui/menolak pengajuan cuti
		body, _ := json.Marshal(data)
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "hr_leave_notification", body)
		log.Println("📨 Event RabbitMQ: Update Status Cuti dikirim ke antrean!")
	}
	return err
}

func DeleteLeaveRequestService(id uint) error {
	return DeleteLeaveRequest(id)
}

func CreateLeaveTypeService(data *LeaveType) error        { return CreateLeaveType(data) }
func GetAllLeaveTypeService() ([]LeaveType, error)        { return GetAllLeaveType() }
func GetLeaveTypeByIDService(id uint) (*LeaveType, error) { return GetLeaveTypeByID(id) }
func UpdateLeaveTypeService(data *LeaveType) error        { return UpdateLeaveType(data) }
func DeleteLeaveTypeService(id uint) error                { return DeleteLeaveType(id) }

func CreateLeaveAllocationService(data *LeaveAllocation) error { return CreateLeaveAllocation(data) }
func GetAllLeaveAllocationService() ([]LeaveAllocation, error) { return GetAllLeaveAllocation() }
func GetLeaveAllocationByIDService(id uint) (*LeaveAllocation, error) {
	return GetLeaveAllocationByID(id)
}
func UpdateLeaveAllocationService(data *LeaveAllocation) error { return UpdateLeaveAllocation(data) }
func DeleteLeaveAllocationService(id uint) error               { return DeleteLeaveAllocation(id) }
