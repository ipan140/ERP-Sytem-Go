package budget

import (
	"ERP-System/config"
)

func GetAllBudgetsRepo() ([]DepartmentBudget, error) {
	var list []DepartmentBudget
	err := config.DB.Order("usage_percent desc").Find(&list).Error
	return list, err
}

func GetBudgetByIDRepo(id uint) (*DepartmentBudget, error) {
	var item DepartmentBudget
	err := config.DB.First(&item, id).Error
	return &item, err
}

func CreateBudgetRepo(b *DepartmentBudget) error {
	return config.DB.Create(b).Error
}

func UpdateBudgetRepo(b *DepartmentBudget) error {
	return config.DB.Save(b).Error
}

func DeleteBudgetRepo(id uint) error {
	return config.DB.Delete(&DepartmentBudget{}, id).Error
}
