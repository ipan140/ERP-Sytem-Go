package crm

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateLead(data *Lead) error {
	return config.DB.Create(data).Error
}

func GetAllLead() ([]Lead, error) {
	var list []Lead
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedLeads(offset int, limit int, search string, stageID uint, salespersonID uint, territory string) ([]Lead, int64, error) {
	var list []Lead
	var total int64

	query := config.DB.Model(&Lead{}).Preload(clause.Associations)

	if stageID > 0 {
		query = query.Where("stage_id = ?", stageID)
	}
	if salespersonID > 0 {
		query = query.Where("salesperson_id = ?", salespersonID)
	}
	if territory != "" && territory != "All" && territory != "Semua" {
		query = query.Where("territory = ?", territory)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ? OR phone ILIKE ? OR affiliate_name ILIKE ?", s, s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetLeadByID(id uint) (*Lead, error) {
	var data Lead
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateLead(data *Lead) error {
	return config.DB.Save(data).Error
}

func DeleteLead(id uint) error {
	return config.DB.Delete(&Lead{}, id).Error
}

func CreateSalesTeam(data *SalesTeam) error { return config.DB.Create(data).Error }
func GetAllSalesTeam() ([]SalesTeam, error) {
	var list []SalesTeam
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetSalesTeamByID(id uint) (*SalesTeam, error) {
	var data SalesTeam
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateSalesTeam(data *SalesTeam) error { return config.DB.Save(data).Error }
func DeleteSalesTeam(id uint) error         { return config.DB.Delete(&SalesTeam{}, id).Error }

func CreateStage(data *Stage) error { return config.DB.Create(data).Error }
func GetAllStage() ([]Stage, error) {
	var list []Stage
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStageByID(id uint) (*Stage, error) {
	var data Stage
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStage(data *Stage) error { return config.DB.Save(data).Error }
func DeleteStage(id uint) error     { return config.DB.Delete(&Stage{}, id).Error }

func CreateActivity(data *Activity) error { return config.DB.Create(data).Error }
func GetAllActivity() ([]Activity, error) {
	var list []Activity
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetActivityByID(id uint) (*Activity, error) {
	var data Activity
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateActivity(data *Activity) error { return config.DB.Save(data).Error }
func DeleteActivity(id uint) error        { return config.DB.Delete(&Activity{}, id).Error }

func CreateSalesCommission(data *SalesCommission) error { return config.DB.Create(data).Error }
func GetAllSalesCommission() ([]SalesCommission, error) {
	var list []SalesCommission
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetSalesCommissionByID(id uint) (*SalesCommission, error) {
	var data SalesCommission
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateSalesCommission(data *SalesCommission) error { return config.DB.Save(data).Error }
func DeleteSalesCommission(id uint) error               { return config.DB.Delete(&SalesCommission{}, id).Error }
