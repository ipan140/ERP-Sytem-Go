package purchase

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreatePurchaseOrder(data *PurchaseOrder) error {
	return config.DB.Create(data).Error
}

func GetAllPurchaseOrder() ([]PurchaseOrder, error) {
	var list []PurchaseOrder
	err := config.DB.Preload("Partner").
		Preload("OrderLines").
		Preload("OrderLines.Product").
		Preload("OrderLines.Product.ProductTemplate").
		Find(&list).Error
	return list, err
}

func GetPaginatedPurchaseOrders(offset, limit int, search, state string, partnerID uint) ([]PurchaseOrder, int64, error) {
	query := config.DB.Model(&PurchaseOrder{}).
		Joins("LEFT JOIN setting.partners ON setting.partners.id = supply_chain.purchase_orders.partner_id")

	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(supply_chain.purchase_orders.name) LIKE ? OR LOWER(setting.partners.name) LIKE ?",
			searchPattern, searchPattern,
		)
	}

	if state != "" {
		query = query.Where("supply_chain.purchase_orders.state = ?", state)
	}

	if partnerID > 0 {
		query = query.Where("supply_chain.purchase_orders.partner_id = ?", partnerID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []PurchaseOrder
	err := query.
		Preload("Partner").
		Preload("OrderLines").
		Preload("OrderLines.Product").
		Preload("OrderLines.Product.ProductTemplate").
		Order("supply_chain.purchase_orders.id DESC").
		Offset(offset).
		Limit(limit).
		Find(&list).Error

	return list, total, err
}

func GetPurchaseOrderByID(id uint) (*PurchaseOrder, error) {
	var data PurchaseOrder
	err := config.DB.
		Preload("Partner").
		Preload("OrderLines").
		Preload("OrderLines.Product").
		Preload("OrderLines.Product.ProductTemplate").
		First(&data, id).Error
	return &data, err
}

func GetPurchaseSummary() (PurchaseSummary, error) {
	var summary PurchaseSummary

	config.DB.Model(&PurchaseOrder{}).Count(&summary.TotalPOCount)
	config.DB.Model(&PurchaseOrder{}).Where("state = ?", "to_approve").Count(&summary.ToApproveCount)
	config.DB.Model(&PurchaseOrder{}).Where("state = ?", "purchase").Count(&summary.ToReceiveCount)

	type SpendResult struct {
		Total float64
	}
	var spend SpendResult
	config.DB.Model(&PurchaseOrder{}).
		Where("state IN ('purchase', 'done')").
		Select("COALESCE(SUM(amount_total), 0) as total").
		Scan(&spend)
	summary.TotalSpentMonthly = spend.Total

	var vendorCount int64
	config.DB.Table("setting.partners").
		Where("is_vendor = true OR id IN (SELECT DISTINCT partner_id FROM supply_chain.purchase_orders)").
		Count(&vendorCount)
	summary.ActiveVendorCount = vendorCount

	return summary, nil
}

func CreatePurchaseOrderWithLines(req CreatePORequest) (*PurchaseOrder, error) {
	var po PurchaseOrder
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&PurchaseOrder{}).Count(&count)
		poName := fmt.Sprintf("PO/%s/%04d", time.Now().Format("2006"), count+1)

		orderDate := time.Now()
		if req.DateOrder != nil {
			orderDate = *req.DateOrder
		}

		var untaxed float64
		var lines []PurchaseOrderLine
		for _, l := range req.Lines {
			subtotal := l.Quantity * l.PriceUnit
			untaxed += subtotal
			lines = append(lines, PurchaseOrderLine{
				ProductID:     l.ProductID,
				Name:          l.Name,
				Quantity:      l.Quantity,
				PriceUnit:     l.PriceUnit,
				PriceSubtotal: subtotal,
			})
		}

		tax := untaxed * 0.11 // PPN 11%
		total := untaxed + tax

		po = PurchaseOrder{
			Name:          poName,
			PartnerID:     req.PartnerID,
			State:         "draft",
			AmountUntaxed: untaxed,
			AmountTax:     tax,
			AmountTotal:   total,
			DateOrder:     orderDate,
			Notes:         req.Notes,
			CreatedAt:     time.Now(),
		}

		if err := tx.Create(&po).Error; err != nil {
			return err
		}

		for i := range lines {
			lines[i].OrderID = po.ID
			if err := tx.Create(&lines[i]).Error; err != nil {
				return err
			}
		}

		po.OrderLines = lines
		return nil
	})

	return &po, err
}

func ConfirmPurchaseOrder(id uint) (*PurchaseOrder, error) {
	var po PurchaseOrder
	if err := config.DB.Preload("OrderLines").First(&po, id).Error; err != nil {
		return nil, err
	}

	if po.State != "draft" && po.State != "sent" {
		return &po, errors.New("hanya PO berstatus draft atau sent yang dapat dikonfirmasi")
	}

	// Gerbang Validasi Approval: jika total > Rp 50.000.000 -> to_approve
	if po.AmountTotal > 50000000 {
		po.State = "to_approve"
	} else {
		po.State = "purchase"
	}

	err := config.DB.Save(&po).Error
	return &po, err
}

func ApprovePurchaseOrder(id uint, approverID uint) (*PurchaseOrder, error) {
	var po PurchaseOrder
	if err := config.DB.First(&po, id).Error; err != nil {
		return nil, err
	}

	if po.State != "to_approve" {
		return &po, errors.New("hanya PO berstatus to_approve yang membutuhkan persetujuan manajer")
	}

	now := time.Now()
	po.State = "purchase"
	if approverID > 0 {
		po.ApprovedBy = &approverID
	}
	po.ApprovedAt = &now

	err := config.DB.Save(&po).Error
	return &po, err
}

func ReceivePurchaseOrderProducts(poID uint, warehouseID uint, items []ReceiveItemInput, notes string) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var po PurchaseOrder
		if err := tx.Preload("OrderLines").Preload("OrderLines.Product").First(&po, poID).Error; err != nil {
			return err
		}

		if po.State != "purchase" && po.State != "to_approve" {
			return errors.New("hanya Purchase Order yang berstatus sah (purchase) yang dapat menerima barang")
		}

		var destLoc inventory.StockLocation
		if warehouseID > 0 {
			tx.Where("warehouse_id = ?", warehouseID).First(&destLoc)
		} else {
			tx.First(&destLoc)
		}
		var destLocID *uint
		if destLoc.ID > 0 {
			destLocID = &destLoc.ID
		}

		receiptCode := fmt.Sprintf("WH/IN/%s/%04d", time.Now().Format("20060102"), po.ID)
		picking := inventory.StockPicking{
			Name:           receiptCode,
			LocationDestID: destLocID,
			PartnerID:      &po.PartnerID,
			State:          "done",
			ScheduledDate:  time.Now(),
			CreatedAt:      time.Now(),
		}
		if err := tx.Create(&picking).Error; err != nil {
			return err
		}

		receivedMap := make(map[uint]float64)
		for _, it := range items {
			receivedMap[it.LineID] = it.QtyReceived
		}

		allDone := true
		for i := range po.OrderLines {
			line := &po.OrderLines[i]
			recQty := receivedMap[line.ID]
			if recQty > 0 {
				line.QtyReceived += recQty
				if err := tx.Save(line).Error; err != nil {
					return err
				}

				if line.ProductID > 0 {
					var prod inventory.Product
					if err := tx.First(&prod, line.ProductID).Error; err == nil {
						prod.StockQty += recQty
						tx.Save(&prod)
					}

					move := inventory.StockMove{
						Name:           fmt.Sprintf("Penerimaan PO %s: %s (Qty: %.2f)", po.Name, line.Name, recQty),
						PickingID:      &picking.ID,
						ProductID:      line.ProductID,
						Quantity:       recQty,
						QuantityDone:   recQty,
						LocationDestID: destLocID,
						State:          "done",
						CreatedAt:      time.Now(),
					}
					tx.Create(&move)

					unitCost := line.PriceUnit
					valLayer := inventory.StockValuationLayer{
						ProductID:   line.ProductID,
						Quantity:    recQty,
						UnitCost:    unitCost,
						Value:       recQty * unitCost,
						Description: fmt.Sprintf("Penerimaan PO: %s (%s)", po.Name, receiptCode),
						CreatedAt:   time.Now(),
					}
					tx.Create(&valLayer)
				}
			}

			if line.QtyReceived < line.Quantity {
				allDone = false
			}
		}

		if allDone {
			po.State = "done"
		} else {
			po.State = "purchase"
		}
		return tx.Save(&po).Error
	})
}

func UpdatePurchaseOrder(data *PurchaseOrder) error {
	return config.DB.Save(data).Error
}

func DeletePurchaseOrder(id uint) error {
	return config.DB.Delete(&PurchaseOrder{}, id).Error
}

func CreatePurchaseRequisition(data *PurchaseRequisition) error { return config.DB.Create(data).Error }
func GetAllPurchaseRequisition() ([]PurchaseRequisition, error) {
	var list []PurchaseRequisition
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetPurchaseRequisitionByID(id uint) (*PurchaseRequisition, error) {
	var data PurchaseRequisition
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePurchaseRequisition(data *PurchaseRequisition) error { return config.DB.Save(data).Error }
func DeletePurchaseRequisition(id uint) error {
	return config.DB.Delete(&PurchaseRequisition{}, id).Error
}

func CreateProductSupplierInfo(data *ProductSupplierInfo) error { return config.DB.Create(data).Error }
func GetAllProductSupplierInfo() ([]ProductSupplierInfo, error) {
	var list []ProductSupplierInfo
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetProductSupplierInfoByID(id uint) (*ProductSupplierInfo, error) {
	var data ProductSupplierInfo
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateProductSupplierInfo(data *ProductSupplierInfo) error { return config.DB.Save(data).Error }
func DeleteProductSupplierInfo(id uint) error {
	return config.DB.Delete(&ProductSupplierInfo{}, id).Error
}

func CreatePurchaseOrderLine(data *PurchaseOrderLine) error { return config.DB.Create(data).Error }
func GetAllPurchaseOrderLine() ([]PurchaseOrderLine, error) {
	var list []PurchaseOrderLine
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetPurchaseOrderLineByID(id uint) (*PurchaseOrderLine, error) {
	var data PurchaseOrderLine
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePurchaseOrderLine(data *PurchaseOrderLine) error { return config.DB.Save(data).Error }
func DeletePurchaseOrderLine(id uint) error                 { return config.DB.Delete(&PurchaseOrderLine{}, id).Error }
