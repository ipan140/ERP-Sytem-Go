package lunch

import (
	"ERP-System/config"
)

func CreateLunchOrder(data *LunchOrder) error {
	return config.DB.Create(data).Error
}

func GetAllLunchOrder() ([]LunchOrder, error) {
	var list []LunchOrder
	err := config.DB.Find(&list).Error
	return list, err
}

func GetLunchOrderByID(id uint) (*LunchOrder, error) {
	var data LunchOrder
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateLunchOrder(data *LunchOrder) error {
	return config.DB.Save(data).Error
}

func DeleteLunchOrder(id uint) error {
	return config.DB.Delete(&LunchOrder{}, id).Error
}
