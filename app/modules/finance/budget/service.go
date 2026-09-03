package budget

import (
	"math"
)

func GetAllBudgetsService() ([]DepartmentBudget, float64, float64, float64, float64, error) {
	list, err := GetAllBudgetsRepo()
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	var totalLimit, totalSpent float64
	for _, it := range list {
		totalLimit += it.AllocatedLimit
		totalSpent += it.RealizedSpent
	}
	remaining := totalLimit - totalSpent
	overallUsage := 0.0
	if totalLimit > 0 {
		overallUsage = (totalSpent / totalLimit) * 100
	}

	return list, totalLimit, totalSpent, remaining, overallUsage, nil
}

func calculateBudgetMetrics(b *DepartmentBudget) {
	b.RemainingBudget = b.AllocatedLimit - b.RealizedSpent
	if b.AllocatedLimit > 0 {
		b.UsagePercent = math.Round(((b.RealizedSpent / b.AllocatedLimit) * 100) * 100) / 100
	}
	if b.UsagePercent > 95 {
		b.Status = "Kritis (Hampir Penuh)"
	} else if b.UsagePercent > 75 {
		b.Status = "Peringatan"
	} else {
		b.Status = "Aman"
	}
}

func CreateBudgetService(b *DepartmentBudget) error {
	calculateBudgetMetrics(b)
	return CreateBudgetRepo(b)
}

func UpdateBudgetService(b *DepartmentBudget) error {
	calculateBudgetMetrics(b)
	return UpdateBudgetRepo(b)
}

func DeleteBudgetService(id uint) error {
	return DeleteBudgetRepo(id)
}
