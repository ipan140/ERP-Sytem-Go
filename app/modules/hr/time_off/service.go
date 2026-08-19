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

func CreateLeaveTypeService(data *LeaveType) error { return CreateLeaveType(data) }
func GetAllLeaveTypeService() ([]LeaveType, error) { return GetAllLeaveType() }
func GetLeaveTypeByIDService(id uint) (*LeaveType, error) { return GetLeaveTypeByID(id) }
func UpdateLeaveTypeService(data *LeaveType) error { return UpdateLeaveType(data) }
func DeleteLeaveTypeService(id uint) error { return DeleteLeaveType(id) }

func CreateLeaveAllocationService(data *LeaveAllocation) error { return CreateLeaveAllocation(data) }
func GetAllLeaveAllocationService() ([]LeaveAllocation, error) { return GetAllLeaveAllocation() }
func GetLeaveAllocationByIDService(id uint) (*LeaveAllocation, error) { return GetLeaveAllocationByID(id) }
func UpdateLeaveAllocationService(data *LeaveAllocation) error { return UpdateLeaveAllocation(data) }
func DeleteLeaveAllocationService(id uint) error { return DeleteLeaveAllocation(id) }
