package quality

import (
	"ERP-System/config"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetPaginatedQualityChecks(offset, limit int, search, result string, productID uint) ([]QualityCheck, int64, error) {
	var list []QualityCheck
	var total int64

	query := config.DB.Model(&QualityCheck{}).
		Preload("Product.ProductTemplate").
		Preload("Point").
		Preload("Picking").
		Preload("Production")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("supply_chain.quality_checks.name ILIKE ? OR product_id IN (SELECT p.id FROM supply_chain.products p JOIN supply_chain.product_templates pt ON p.product_template_id = pt.id WHERE pt.name ILIKE ? OR p.default_code ILIKE ?)", s, s, s)
	}

	if result != "" && result != "all" {
		query = query.Where("result = ?", result)
	}

	if productID > 0 {
		query = query.Where("product_id = ?", productID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetQualitySummary() (QualitySummary, error) {
	var summary QualitySummary

	config.DB.Model(&QualityCheck{}).Count(&summary.TotalChecksCount)
	config.DB.Model(&QualityCheck{}).Where("result = 'pending'").Count(&summary.PendingChecksCount)
	config.DB.Model(&QualityCheck{}).Where("result = 'pass'").Count(&summary.PassedChecksCount)
	config.DB.Model(&QualityCheck{}).Where("result = 'fail'").Count(&summary.FailedChecksCount)
	config.DB.Model(&QualityPoint{}).Count(&summary.TotalPointsCount)

	return summary, nil
}

func CreateQualityCheckWithSequence(req CreateQualityCheckRequest) (*QualityCheck, error) {
	var qc QualityCheck
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&QualityCheck{}).Count(&count)
		qcName := fmt.Sprintf("QC/%s/%04d", time.Now().Format("2006"), count+1)

		qc = QualityCheck{
			Name:         qcName,
			PointID:      req.PointID,
			ProductID:    req.ProductID,
			PickingID:    req.PickingID,
			ProductionID: req.ProductionID,
			MeasureValue: req.MeasureValue,
			Notes:        req.Notes,
			Result:       "pending",
		}

		return tx.Create(&qc).Error
	})

	if err != nil {
		return nil, err
	}

	return GetQualityCheckByID(qc.ID)
}

func ProcessQualityCheck(id uint, req ProcessQCRequest, inspectorID *uint) (*QualityCheck, error) {
	var qc QualityCheck
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Point").First(&qc, id).Error; err != nil {
			return err
		}

		res := req.Result
		if res != "pass" && res != "fail" {
			return errors.New("hasil inspeksi QC harus bernilai pass atau fail")
		}

		now := time.Now()
		qc.Result = res
		qc.MeasureValue = req.MeasureValue
		qc.Notes = req.Notes
		qc.InspectorID = inspectorID
		qc.InspectedAt = &now

		return tx.Save(&qc).Error
	})

	if err != nil {
		return nil, err
	}

	return GetQualityCheckByID(id)
}

func GetPaginatedQualityPoints(offset, limit int, search string) ([]QualityPoint, int64, error) {
	var list []QualityPoint
	var total int64

	query := config.DB.Model(&QualityPoint{}).Preload("Product.ProductTemplate")
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR product_id IN (SELECT p.id FROM supply_chain.products p JOIN supply_chain.product_templates pt ON p.product_template_id = pt.id WHERE pt.name ILIKE ?)", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func CreateQualityPoint(data *QualityPoint) error {
	return config.DB.Create(data).Error
}

func GetAllQualityPoints() ([]QualityPoint, error) {
	var list []QualityPoint
	err := config.DB.Preload("Product.ProductTemplate").Find(&list).Error
	return list, err
}

func DeleteQualityPoint(id uint) error {
	return config.DB.Delete(&QualityPoint{}, id).Error
}

func CreateQualityCheck(data *QualityCheck) error {
	return config.DB.Create(data).Error
}

func GetAllQualityCheck() ([]QualityCheck, error) {
	var list []QualityCheck
	err := config.DB.
		Preload("Product.ProductTemplate").
		Preload("Point").
		Preload("Picking").
		Preload("Production").
		Find(&list).Error
	return list, err
}

func GetQualityCheckByID(id uint) (*QualityCheck, error) {
	var data QualityCheck
	err := config.DB.
		Preload("Product.ProductTemplate").
		Preload("Point").
		Preload("Picking").
		Preload("Production").
		First(&data, id).Error
	return &data, err
}

func UpdateQualityCheck(data *QualityCheck) error {
	return config.DB.Save(data).Error
}

func DeleteQualityCheck(id uint) error {
	return config.DB.Delete(&QualityCheck{}, id).Error
}

