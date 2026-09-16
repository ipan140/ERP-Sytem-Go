package documents

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateWorkspace(data *Workspace) error {
	return config.DB.Create(data).Error
}

func GetAllWorkspace() ([]Workspace, error) {
	var list []Workspace
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedWorkspaces(offset, limit int, search string) ([]Workspace, int64, error) {
	var list []Workspace
	var total int64
	query := config.DB.Model(&Workspace{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ?", s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetWorkspaceByID(id uint) (*Workspace, error) {
	var data Workspace
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateWorkspace(data *Workspace) error {
	return config.DB.Save(data).Error
}

func DeleteWorkspace(id uint) error {
	return config.DB.Delete(&Workspace{}, id).Error
}
