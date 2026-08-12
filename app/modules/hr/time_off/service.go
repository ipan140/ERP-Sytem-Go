package time_off

func CreateLeaveRequestService(data *LeaveRequest) error {
	return CreateLeaveRequest(data)
}

func GetAllLeaveRequestService() ([]LeaveRequest, error) {
	return GetAllLeaveRequest()
}

func GetLeaveRequestByIDService(id uint) (*LeaveRequest, error) {
	return GetLeaveRequestByID(id)
}

func UpdateLeaveRequestService(data *LeaveRequest) error {
	return UpdateLeaveRequest(data)
}

func DeleteLeaveRequestService(id uint) error {
	return DeleteLeaveRequest(id)
}
