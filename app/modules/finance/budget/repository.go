package budget

import (
	"ERP-System/config"
)

func GetAllBudgetsRepo() ([]DepartmentBudget, error) {
	var list []DepartmentBudget
	err := config.DB.Order("usage_percent desc").Find(&list).Error
	return list, err
}

func GetPaginatedBudgetsRepo(offset int, limit int, search string) ([]DepartmentBudget, int64, float64, float64, float64, float64, error) {
	var list []DepartmentBudget
	var total int64

	var summary struct {
		TotalLimit float64
		TotalSpent float64
	}
	config.DB.Model(&DepartmentBudget{}).Select("COALESCE(SUM(allocated_limit), 0) as total_limit, COALESCE(SUM(realized_spent), 0) as total_spent").Scan(&summary)

	remaining := summary.TotalLimit - summary.TotalSpent
	overallUsage := 0.0
	if summary.TotalLimit > 0 {
		overallUsage = (summary.TotalSpent / summary.TotalLimit) * 100
	}

	query := config.DB.Model(&DepartmentBudget{})
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("department_name ILIKE ? OR fiscal_period ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, 0, 0, 0, err
	}

	err := query.Order("usage_percent desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, summary.TotalLimit, summary.TotalSpent, remaining, overallUsage, err
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
