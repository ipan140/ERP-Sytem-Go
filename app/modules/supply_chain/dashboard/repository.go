package dashboard

import (
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/app/modules/supply_chain/maintenance"
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/app/modules/supply_chain/purchase"
	"ERP-System/app/modules/supply_chain/quality"
	"ERP-System/config"
	"fmt"
	"math"
	"time"
)

func GetScmDashboardSummary() (*ScmDashboardSummary, error) {
	var summary ScmDashboardSummary

	// 1. Inventory Metrics
	config.DB.Model(&inventory.Product{}).Count(&summary.TotalSKUCount)
	config.DB.Model(&inventory.Product{}).Where("stock_qty <= 10").Count(&summary.LowStockCount)

	type ValResult struct {
		TotalValuation float64
	}
	var valRes ValResult
	config.DB.Raw(`
		SELECT COALESCE(SUM(p.stock_qty * COALESCE(NULLIF(pt.standard_price, 0), pt.list_price, 0)), 0) as total_valuation
		FROM supply_chain.products p
		JOIN supply_chain.product_templates pt ON p.product_template_id = pt.id
		WHERE p.stock_qty > 0
	`).Scan(&valRes)
	summary.TotalInventoryValuation = valRes.TotalValuation

	// 2. Purchase Metrics
	config.DB.Model(&purchase.PurchaseOrder{}).
		Where("state IN ('purchase', 'to_receive')").
		Count(&summary.TotalActivePOs)

	var spendRes struct {
		TotalSpend float64
	}
	config.DB.Raw(`
		SELECT COALESCE(SUM(amount_total), 0) as total_spend
		FROM supply_chain.purchase_orders
		WHERE state IN ('purchase', 'done')
		  AND date_order >= ?
	`, time.Now().AddDate(0, -1, 0)).Scan(&spendRes)
	summary.MonthlyPurchaseSpend = spendRes.TotalSpend

	// 3. Manufacturing Metrics
	config.DB.Model(&manufacturing.MrpProduction{}).
		Where("state IN ('confirmed', 'progress')").
		Count(&summary.ActiveManufacturingMOs)

	config.DB.Model(&manufacturing.MrpProduction{}).
		Where("state = 'done' AND created_at >= ?", time.Now().AddDate(0, -1, 0)).
		Count(&summary.CompletedMOsMonthly)

	// 4. Quality Control Metrics
	var passCount, failCount, totalChecks int64
	config.DB.Model(&quality.QualityCheck{}).Count(&totalChecks)
	config.DB.Model(&quality.QualityCheck{}).Where("result = 'pass'").Count(&passCount)
	config.DB.Model(&quality.QualityCheck{}).Where("result = 'fail'").Count(&failCount)
	summary.TotalQCChecks = totalChecks
	if passCount+failCount > 0 {
		summary.QualityPassRate = math.Round((float64(passCount)/float64(passCount+failCount)*100)*10) / 10
	} else {
		summary.QualityPassRate = 100.0
	}

	// 5. Maintenance Metrics
	config.DB.Model(&maintenance.MaintenanceEquipment{}).Count(&summary.TotalEquipments)
	config.DB.Model(&maintenance.MaintenanceRequest{}).
		Where("state IN ('todo', 'progress')").
		Count(&summary.ActiveMaintenanceTickets)

	// 6. Actionable Alerts
	alerts := []ScmAlertItem{}

	// Check low stock
	if summary.LowStockCount > 0 {
		alerts = append(alerts, ScmAlertItem{
			Type:     "stock",
			Severity: "warning",
			Title:    fmt.Sprintf("%d SKU Stok Menipis", summary.LowStockCount),
			Message:  "Terdapat produk dengan persediaan di bawah batas minimum (10 unit). Segera terbitkan RFQ Pembelian.",
			RefCode:  "STOCK-ALERT",
		})
	}

	// Check broken machines
	if summary.ActiveMaintenanceTickets > 0 {
		alerts = append(alerts, ScmAlertItem{
			Type:     "maintenance",
			Severity: "danger",
			Title:    fmt.Sprintf("%d Tiket Perawatan Mesin Aktif", summary.ActiveMaintenanceTickets),
			Message:  "Terdapat mesin/aset dalam status antrean perbaikan atau pengerjaan teknisi.",
			RefCode:  "MAINT-ALERT",
		})
	}

	// Check active manufacturing
	if summary.ActiveManufacturingMOs > 0 {
		alerts = append(alerts, ScmAlertItem{
			Type:     "manufacturing",
			Severity: "info",
			Title:    fmt.Sprintf("%d SPK Manufaktur Berjalan", summary.ActiveManufacturingMOs),
			Message:  "Lini produksi aktif sedang memproses barang jadi.",
			RefCode:  "MO-RUNNING",
		})
	}

	summary.Alerts = alerts
	return &summary, nil
}

func GetScmCalendarEvents() ([]ScmCalendarEvent, error) {
	var events []ScmCalendarEvent

	// 1. Purchase Orders (Expected Deliveries)
	var pos []purchase.PurchaseOrder
	config.DB.Where("state IN ('purchase', 'to_receive', 'done')").
		Order("id desc").Limit(30).Find(&pos)
	for _, po := range pos {
		date := po.DateOrder
		if date.IsZero() {
			date = po.CreatedAt
		}
		color := "#3B82F6" // Blue
		if po.State == "done" {
			color = "#10B981" // Green
		}
		events = append(events, ScmCalendarEvent{
			ID:          fmt.Sprintf("po-%d", po.ID),
			Title:       fmt.Sprintf("PO: %s (%s)", po.Name, po.State),
			Start:       date,
			End:         date.Add(2 * time.Hour),
			Type:        "purchase",
			Status:      po.State,
			Color:       color,
			Description: fmt.Sprintf("Pembelian Bahan/Barang senilai Rp %.0f", po.AmountTotal),
		})
	}

	// 2. Manufacturing Orders (Production Runs)
	var mos []manufacturing.MrpProduction
	config.DB.Order("id desc").Limit(30).Find(&mos)
	for _, mo := range mos {
		date := mo.DatePlannedStart
		if date.IsZero() {
			date = mo.CreatedAt
		}
		color := "#F59E0B" // Amber
		if mo.State == "done" {
			color = "#059669"
		}
		events = append(events, ScmCalendarEvent{
			ID:          fmt.Sprintf("mo-%d", mo.ID),
			Title:       fmt.Sprintf("SPK: %s (%d Unit)", mo.Name, int(mo.ProductQty)),
			Start:       date,
			End:         date.Add(8 * time.Hour),
			Type:        "manufacturing",
			Status:      mo.State,
			Color:       color,
			Description: fmt.Sprintf("Perintah Kerja Manufaktur status %s", mo.State),
		})
	}

	// 3. Maintenance Requests (Scheduled Services)
	var mrs []maintenance.MaintenanceRequest
	config.DB.Preload("Equipment").Order("id desc").Limit(30).Find(&mrs)
	for _, mr := range mrs {
		date := mr.CreatedAt
		if mr.ScheduleDate != nil && !mr.ScheduleDate.IsZero() {
			date = *mr.ScheduleDate
		}
		color := "#EF4444" // Red for corrective
		if mr.Type == "preventive" {
			color = "#8B5CF6" // Purple for preventive
		}
		if mr.State == "done" {
			color = "#10B981"
		}
		eqName := "Mesin Pabrik"
		if mr.Equipment != nil && mr.Equipment.Name != "" {
			eqName = mr.Equipment.Name
		}
		events = append(events, ScmCalendarEvent{
			ID:          fmt.Sprintf("maint-%d", mr.ID),
			Title:       fmt.Sprintf("Maint: %s (%s)", eqName, mr.Type),
			Start:       date,
			End:         date.Add(4 * time.Hour),
			Type:        "maintenance",
			Status:      mr.State,
			Color:       color,
			Description: fmt.Sprintf("%s - Prioritas %s", mr.Name, mr.Priority),
		})
	}

	return events, nil
}
