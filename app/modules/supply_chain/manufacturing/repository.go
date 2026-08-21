package manufacturing

import (
	"ERP-System/config"
)

func CreateMrpProduction(data *MrpProduction) error {
	return config.DB.Create(data).Error
}

func GetAllMrpProduction() ([]MrpProduction, error) {
	var list []MrpProduction
	err := config.DB.Find(&list).Error
	return list, err
}

func GetMrpProductionByID(id uint) (*MrpProduction, error) {
	var data MrpProduction
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateMrpProduction(data *MrpProduction) error {
	return config.DB.Save(data).Error
}

func DeleteMrpProduction(id uint) error {
	return config.DB.Delete(&MrpProduction{}, id).Error
}

func CreateMrpWorkcenter(data *MrpWorkcenter) error { return config.DB.Create(data).Error }
func GetAllMrpWorkcenter() ([]MrpWorkcenter, error) {
	var list []MrpWorkcenter
	err := config.DB.Find(&list).Error
	return list, err
}
func GetMrpWorkcenterByID(id uint) (*MrpWorkcenter, error) {
	var data MrpWorkcenter
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateMrpWorkcenter(data *MrpWorkcenter) error { return config.DB.Save(data).Error }
func DeleteMrpWorkcenter(id uint) error             { return config.DB.Delete(&MrpWorkcenter{}, id).Error }

func CreateMrpBom(data *MrpBom) error { return config.DB.Create(data).Error }
func GetAllMrpBom() ([]MrpBom, error) {
	var list []MrpBom
	err := config.DB.Find(&list).Error
	return list, err
}
func GetMrpBomByID(id uint) (*MrpBom, error) {
	var data MrpBom
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateMrpBom(data *MrpBom) error { return config.DB.Save(data).Error }
func DeleteMrpBom(id uint) error      { return config.DB.Delete(&MrpBom{}, id).Error }

func CreateMrpBomLine(data *MrpBomLine) error { return config.DB.Create(data).Error }
func GetAllMrpBomLine() ([]MrpBomLine, error) {
	var list []MrpBomLine
	err := config.DB.Find(&list).Error
	return list, err
}
func GetMrpBomLineByID(id uint) (*MrpBomLine, error) {
	var data MrpBomLine
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateMrpBomLine(data *MrpBomLine) error { return config.DB.Save(data).Error }
func DeleteMrpBomLine(id uint) error          { return config.DB.Delete(&MrpBomLine{}, id).Error }

func CreateMrpBomByproduct(data *MrpBomByproduct) error { return config.DB.Create(data).Error }
func GetAllMrpBomByproduct() ([]MrpBomByproduct, error) {
	var list []MrpBomByproduct
	err := config.DB.Find(&list).Error
	return list, err
}
func GetMrpBomByproductByID(id uint) (*MrpBomByproduct, error) {
	var data MrpBomByproduct
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateMrpBomByproduct(data *MrpBomByproduct) error { return config.DB.Save(data).Error }
func DeleteMrpBomByproduct(id uint) error               { return config.DB.Delete(&MrpBomByproduct{}, id).Error }

func CreateMrpWorkorder(data *MrpWorkorder) error { return config.DB.Create(data).Error }
func GetAllMrpWorkorder() ([]MrpWorkorder, error) {
	var list []MrpWorkorder
	err := config.DB.Find(&list).Error
	return list, err
}
func GetMrpWorkorderByID(id uint) (*MrpWorkorder, error) {
	var data MrpWorkorder
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateMrpWorkorder(data *MrpWorkorder) error { return config.DB.Save(data).Error }
func DeleteMrpWorkorder(id uint) error            { return config.DB.Delete(&MrpWorkorder{}, id).Error }
