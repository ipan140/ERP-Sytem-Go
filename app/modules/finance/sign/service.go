package sign

func CreateSignatureRequestService(data *SignatureRequest) error {
	return CreateSignatureRequest(data)
}

func GetAllSignatureRequestService() ([]SignatureRequest, error) {
	return GetAllSignatureRequest()
}

func GetSignatureRequestByIDService(id uint) (*SignatureRequest, error) {
	return GetSignatureRequestByID(id)
}

func UpdateSignatureRequestService(data *SignatureRequest) error {
	return UpdateSignatureRequest(data)
}

func DeleteSignatureRequestService(id uint) error {
	return DeleteSignatureRequest(id)
}
