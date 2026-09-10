package dashboard

import "time"

type ScmAlertItem struct {
	Type     string `json:"type"`     // stock, purchase, manufacturing, maintenance, quality
	Severity string `json:"severity"` // info, warning, danger
	Title    string `json:"title"`
	Message  string `json:"message"`
	RefCode  string `json:"ref_code"`
}

type ScmDashboardSummary struct {
	TotalInventoryValuation float64        `json:"total_inventory_valuation"`
	TotalSKUCount           int64          `json:"total_sku_count"`
	LowStockCount           int64          `json:"low_stock_count"`
	TotalActivePOs          int64          `json:"total_active_pos"`
	MonthlyPurchaseSpend    float64        `json:"monthly_purchase_spend"`
	ActiveManufacturingMOs  int64          `json:"active_manufacturing_mos"`
	CompletedMOsMonthly     int64          `json:"completed_mos_monthly"`
	QualityPassRate         float64        `json:"quality_pass_rate"`
	TotalQCChecks           int64          `json:"total_qc_checks"`
	TotalEquipments         int64          `json:"total_equipments"`
	ActiveMaintenanceTickets int64         `json:"active_maintenance_tickets"`
	Alerts                  []ScmAlertItem `json:"alerts"`
}

type ScmCalendarEvent struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Type        string    `json:"type"` // purchase, manufacturing, maintenance
	Status      string    `json:"status"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
}
