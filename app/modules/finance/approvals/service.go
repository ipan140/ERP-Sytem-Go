package approvals

func CreateApprovalRequestService(data *ApprovalRequest) error {
	return CreateApprovalRequest(data)
}

func GetAllApprovalRequestService() ([]ApprovalRequest, error) {
	return GetAllApprovalRequest()
}

func GetApprovalRequestByIDService(id uint) (*ApprovalRequest, error) {
	return GetApprovalRequestByID(id)
}

func UpdateApprovalRequestService(data *ApprovalRequest) error {
	return UpdateApprovalRequest(data)
}

func DeleteApprovalRequestService(id uint) error {
	return DeleteApprovalRequest(id)
}
