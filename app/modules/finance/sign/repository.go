package sign

import (
	"ERP-System/config"
)

func CreateSignatureRequest(data *SignatureRequest) error {
	return config.DB.Create(data).Error
}

func GetAllSignatureRequest() ([]SignatureRequest, error) {
	var list []SignatureRequest
	err := config.DB.Find(&list).Error
	return list, err
}

func GetSignatureRequestByID(id uint) (*SignatureRequest, error) {
	var data SignatureRequest
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateSignatureRequest(data *SignatureRequest) error {
	return config.DB.Save(data).Error
}

func DeleteSignatureRequest(id uint) error {
	return config.DB.Delete(&SignatureRequest{}, id).Error
}
