package sign

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateSignatureRequest(data *SignatureRequest) error {
	return config.DB.Create(data).Error
}

func GetAllSignatureRequest() ([]SignatureRequest, error) {
	var list []SignatureRequest
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedSignatureRequests(offset int, limit int, search, status string) ([]SignatureRequest, int64, error) {
	var list []SignatureRequest
	var total int64

	query := config.DB.Model(&SignatureRequest{})

	if status != "" && status != "all" && status != "All" {
		query = query.Where("status = ?", status)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR document_title ILIKE ? OR signer_name ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetSignatureRequestByID(id uint) (*SignatureRequest, error) {
	var data SignatureRequest
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateSignatureRequest(data *SignatureRequest) error {
	return config.DB.Save(data).Error
}

func DeleteSignatureRequest(id uint) error {
	return config.DB.Delete(&SignatureRequest{}, id).Error
}
