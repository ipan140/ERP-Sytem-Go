package point_of_sale

import (
	"ERP-System/config"
)

func CreatePosSession(data *PosSession) error {
	return config.DB.Create(data).Error
}

func GetAllPosSession() ([]PosSession, error) {
	var list []PosSession
	err := config.DB.Find(&list).Error
	return list, err
}

func GetPosSessionByID(id uint) (*PosSession, error) {
	var data PosSession
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdatePosSession(data *PosSession) error {
	return config.DB.Save(data).Error
}

func DeletePosSession(id uint) error {
	return config.DB.Delete(&PosSession{}, id).Error
}

func CreatePosConfig(data *PosConfig) error { return config.DB.Create(data).Error }
func GetAllPosConfig() ([]PosConfig, error) {
	var list []PosConfig
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPosConfigByID(id uint) (*PosConfig, error) {
	var data PosConfig
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePosConfig(data *PosConfig) error { return config.DB.Save(data).Error }
func DeletePosConfig(id uint) error         { return config.DB.Delete(&PosConfig{}, id).Error }

func CreatePosOrder(data *PosOrder) error { return config.DB.Create(data).Error }
func GetAllPosOrder() ([]PosOrder, error) {
	var list []PosOrder
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPosOrderByID(id uint) (*PosOrder, error) {
	var data PosOrder
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePosOrder(data *PosOrder) error { return config.DB.Save(data).Error }
func DeletePosOrder(id uint) error        { return config.DB.Delete(&PosOrder{}, id).Error }

func CreatePosOrderLine(data *PosOrderLine) error { return config.DB.Create(data).Error }
func GetAllPosOrderLine() ([]PosOrderLine, error) {
	var list []PosOrderLine
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPosOrderLineByID(id uint) (*PosOrderLine, error) {
	var data PosOrderLine
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePosOrderLine(data *PosOrderLine) error { return config.DB.Save(data).Error }
func DeletePosOrderLine(id uint) error            { return config.DB.Delete(&PosOrderLine{}, id).Error }

func CreatePosPayment(data *PosPayment) error { return config.DB.Create(data).Error }
func GetAllPosPayment() ([]PosPayment, error) {
	var list []PosPayment
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPosPaymentByID(id uint) (*PosPayment, error) {
	var data PosPayment
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePosPayment(data *PosPayment) error { return config.DB.Save(data).Error }
func DeletePosPayment(id uint) error          { return config.DB.Delete(&PosPayment{}, id).Error }

func CreateLoyaltyProgram(data *LoyaltyProgram) error { return config.DB.Create(data).Error }
func GetAllLoyaltyProgram() ([]LoyaltyProgram, error) {
	var list []LoyaltyProgram
	err := config.DB.Find(&list).Error
	return list, err
}
func GetLoyaltyProgramByID(id uint) (*LoyaltyProgram, error) {
	var data LoyaltyProgram
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateLoyaltyProgram(data *LoyaltyProgram) error { return config.DB.Save(data).Error }
func DeleteLoyaltyProgram(id uint) error              { return config.DB.Delete(&LoyaltyProgram{}, id).Error }
