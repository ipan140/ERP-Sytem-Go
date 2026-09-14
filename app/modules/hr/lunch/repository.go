package lunch

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateLunchOrder(data *LunchOrder) error {
	return config.DB.Create(data).Error
}

func GetAllLunchOrder() ([]LunchOrder, error) {
	var list []LunchOrder
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedLunchOrders(offset, limit int, search, employeeID, state string) ([]LunchOrder, int64, error) {
	var list []LunchOrder
	var total int64

	query := config.DB.Model(&LunchOrder{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.lunch_orders.employee_id = ?", employeeID)
	}

	if state != "" && state != "all" {
		query = query.Where("hrd.lunch_orders.state = ?", state)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.lunch_orders.employee_id").
			Where("hrd.employees.name ILIKE ? OR hrd.lunch_orders.state ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.lunch_orders.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetPaginatedLunchCashmoves(offset, limit int, search, employeeID string) ([]LunchCashmove, int64, error) {
	var list []LunchCashmove
	var total int64

	query := config.DB.Model(&LunchCashmove{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.lunch_cashmoves.employee_id = ?", employeeID)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.lunch_cashmoves.employee_id").
			Where("hrd.employees.name ILIKE ? OR hrd.lunch_cashmoves.description ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.lunch_cashmoves.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetLunchOrderByID(id uint) (*LunchOrder, error) {
	var data LunchOrder
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateLunchOrder(data *LunchOrder) error {
	return config.DB.Save(data).Error
}

func DeleteLunchOrder(id uint) error {
	return config.DB.Delete(&LunchOrder{}, id).Error
}

func CreateLunchCashmove(data *LunchCashmove) error { return config.DB.Create(data).Error }
func GetAllLunchCashmove() ([]LunchCashmove, error) {
	var list []LunchCashmove
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetLunchCashmoveByID(id uint) (*LunchCashmove, error) {
	var data LunchCashmove
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateLunchCashmove(data *LunchCashmove) error { return config.DB.Save(data).Error }
func DeleteLunchCashmove(id uint) error             { return config.DB.Delete(&LunchCashmove{}, id).Error }
