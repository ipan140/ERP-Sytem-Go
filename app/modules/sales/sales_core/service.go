package sales_core

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/finance/invoicing"
	"ERP-System/app/modules/sales"
	"ERP-System/app/modules/services/project"
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"ERP-System/pkg/rabbitmq"

	"gorm.io/gorm"
)

type QuotationPayload struct {
	QuotationID uint   `json:"quotation_id"`
	Format      string `json:"format"`
	UserID      uint   `json:"user_id"`
}

func GenerateQuotationPDFService(quotationID uint, userID uint) error {
	if rabbitmq.Channel != nil {
		req := QuotationPayload{
			QuotationID: quotationID,
			Format:      "pdf",
			UserID:      userID,
		}
		body, _ := json.Marshal(req)
		err := rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body) // Reusing doc generator
		log.Printf("📄 Event RabbitMQ: Generate PDF Quotation %d dikirim ke antrean!", quotationID)
		return err
	}
	return nil
}

func CreateSaleOrderService(data *SaleOrder) error {
	if data.Name == "" {
		data.Name = fmt.Sprintf("SO/%s/%04d", time.Now().Format("2006"), time.Now().Unix()%10000)
	}
	if data.State == "" {
		data.State = "draft"
	}
	if data.DateOrder.IsZero() {
		data.DateOrder = time.Now()
	}
	if data.PricelistName == "" {
		data.PricelistName = "Standard Retail"
	}

	// Hitung DueDate berdasarkan PaymentTerm
	if data.PaymentTerm == "" {
		data.PaymentTerm = "Net 30"
	}
	switch data.PaymentTerm {
	case "COD":
		data.DueDate = data.DateOrder
	case "Net 14":
		data.DueDate = data.DateOrder.AddDate(0, 0, 14)
	case "Net 60":
		data.DueDate = data.DateOrder.AddDate(0, 0, 60)
	case "DP 30%":
		data.DueDate = data.DateOrder.AddDate(0, 0, 7)
	default: // Net 30
		data.DueDate = data.DateOrder.AddDate(0, 0, 30)
	}

	// PPN & Pajak Dinamis
	if data.TaxRate <= 0 && data.TaxType == "" {
		data.TaxRate = 11.0
		data.TaxType = "PPN 11%"
	} else if data.TaxType == "PPN 12%" {
		data.TaxRate = 12.0
	} else if data.TaxType == "Non-PPN" {
		data.TaxRate = 0.0
	} else if data.TaxType == "PPh 23" {
		data.TaxRate = 2.0
	}

	// Calculate subtotal for lines if provided & calculate max discount
	var untaxed float64
	var maxDisc float64
	for i := range data.OrderLines {
		if data.OrderLines[i].Discount > maxDisc {
			maxDisc = data.OrderLines[i].Discount
		}
		data.OrderLines[i].SubTotal = data.OrderLines[i].Quantity * data.OrderLines[i].UnitPrice * (1 - (data.OrderLines[i].Discount / 100))
		untaxed += data.OrderLines[i].SubTotal
	}
	data.MaxDiscount = maxDisc
	// FASE 2: Hierarki Persetujuan Diskon Bertingkat (Tiered Approval)
	if maxDisc > 20.0 || untaxed > 250000000.0 {
		data.NeedsApproval = true
		data.ApprovalTier = "Director"
		if data.ApprovalStatus == "" || data.ApprovalStatus == "None" {
			data.ApprovalStatus = "Pending_Director"
		}
	} else if maxDisc > 10.0 {
		data.NeedsApproval = true
		data.ApprovalTier = "ASM"
		if data.ApprovalStatus == "" || data.ApprovalStatus == "None" {
			data.ApprovalStatus = "Pending_ASM"
		}
	} else {
		data.NeedsApproval = false
		data.ApprovalTier = "Auto"
		data.ApprovalStatus = "None"
	}

	if untaxed > 0 {
		data.AmountUntaxed = untaxed
		data.AmountTax = untaxed * (data.TaxRate / 100.0)
		data.AmountTotal = untaxed + data.AmountTax
	}

	// FASE 1: Verifikasi Limit Kredit Pelanggan
	if data.PartnerID > 0 {
		var partner base.Partner
		if err := config.DB.First(&partner, data.PartnerID).Error; err == nil {
			data.CreditLimit = partner.CreditLimit
			data.CurrentReceivable = partner.TotalReceivable
			if partner.IsCreditHold {
				data.CreditStatus = "Hold"
			} else if partner.CreditLimit > 0 && (partner.TotalReceivable+data.AmountTotal) > partner.CreditLimit {
				data.CreditStatus = "Exceeded"
			} else if partner.CreditLimit > 0 && (partner.TotalReceivable+data.AmountTotal) > (partner.CreditLimit*0.8) {
				data.CreditStatus = "Warning"
			} else {
				data.CreditStatus = "OK"
			}
		}
	}

	// FASE 4: Hitung Komisi Sales otomatis (default 3% atau custom rate)
	if data.CommissionRate <= 0 {
		data.CommissionRate = 3.0
	}
	data.CommissionAmount = data.AmountUntaxed * (data.CommissionRate / 100.0)
	if data.CommissionStatus == "" {
		data.CommissionStatus = "Unpaid"
	}

	return CreateSaleOrder(data)
}

func GetAllSaleOrderService() ([]SaleOrder, error)        { return GetAllSaleOrder() }
func GetPaginatedSaleOrderService(offset int, limit int, search string, branch string, status string) ([]SaleOrder, int64, error) {
	return GetPaginatedSaleOrders(offset, limit, search, branch, status)
}
func GetSaleOrderByIDService(id uint) (*SaleOrder, error) { return GetSaleOrderByID(id) }

func UpdateSaleOrderService(data *SaleOrder) error {
	// Recalculate totals if order lines are passed
	if len(data.OrderLines) > 0 {
		var untaxed float64
		var maxDisc float64
		for i := range data.OrderLines {
			if data.OrderLines[i].Discount > maxDisc {
				maxDisc = data.OrderLines[i].Discount
			}
			data.OrderLines[i].SubTotal = data.OrderLines[i].Quantity * data.OrderLines[i].UnitPrice * (1 - (data.OrderLines[i].Discount / 100))
			untaxed += data.OrderLines[i].SubTotal
		}
		data.MaxDiscount = maxDisc
		if (maxDisc > 20.0 || untaxed > 250000000.0) && data.ApprovalStatus != "Approved" {
			data.NeedsApproval = true
			data.ApprovalTier = "Director"
			data.ApprovalStatus = "Pending_Director"
		} else if maxDisc > 10.0 && data.ApprovalStatus != "Approved" {
			data.NeedsApproval = true
			data.ApprovalTier = "ASM"
			data.ApprovalStatus = "Pending_ASM"
		} else if maxDisc <= 10.0 && untaxed <= 250000000.0 {
			data.NeedsApproval = false
			data.ApprovalTier = "Auto"
			data.ApprovalStatus = "None"
		}

		if data.TaxRate <= 0 {
			data.TaxRate = 11.0
		}
		data.AmountUntaxed = untaxed
		data.AmountTax = untaxed * (data.TaxRate / 100.0)
		data.AmountTotal = untaxed + data.AmountTax
	}

	if data.CommissionRate <= 0 {
		data.CommissionRate = 3.0
	}
	data.CommissionAmount = data.AmountUntaxed * (data.CommissionRate / 100.0)

	if data.State == "sale" {
		project.AutoCreateProjectFromSales(data.ID, data.PartnerID, data.Name)
		
		// [RabbitMQ] - Publish Event agar modul Finance dan Supply Chain dapat menangkapnya
		orderIDStr := fmt.Sprintf("%d", data.ID)
		customerIDStr := fmt.Sprintf("%d", data.PartnerID)
		_ = sales.PublishOrderCompletedEvent(orderIDStr, data.AmountTotal, customerIDStr)
	}
	return UpdateSaleOrder(data)
}

// ConfirmSaleOrderService validates stock availability, discount approval, and transitions status to 'sale' (Sales Order Confirmed)
func ConfirmSaleOrderService(id uint) (*SaleOrder, error) {
	order, err := GetSaleOrderByID(id)
	if err != nil {
		return nil, err
	}

	// [FASE 2 Guard] Persetujuan Diskon Bertingkat (ASM atau National Director)
	if order.NeedsApproval && order.ApprovalStatus != "Approved" {
		roleRequired := "Area Sales Manager (ASM)"
		if order.ApprovalTier == "Director" {
			roleRequired = "National Sales Director (Diskon >20% atau Nilai >Rp 250 Jt)"
		}
		return nil, fmt.Errorf("pesanan memerlukan persetujuan dari %s (Diskon %.1f%%, Status: %s)",
			roleRequired, order.MaxDiscount, order.ApprovalStatus)
	}

	// [FASE 1 Guard] Pengecekan Plafon Limit Kredit & Akun On-Hold
	if (order.CreditStatus == "Exceeded" || order.CreditStatus == "Hold") && !order.IsCreditBypassed {
		return nil, fmt.Errorf("pesanan DITOLAK: Saldo piutang pelanggan (Rp %.0f) melebihi limit kredit (Rp %.0f) atau status akun On-Hold. Wajib meminta persetujuan Finance (Bypass Credit Hold)!",
			order.CurrentReceivable+order.AmountTotal, order.CreditLimit)
	}

	// Cek ketersediaan stok fisik gudang jika produk terdaftar
	for _, line := range order.OrderLines {
		if line.ProductID > 0 {
			var prod inventory.Product
			if err := config.DB.First(&prod, line.ProductID).Error; err == nil {
				if prod.StockQty < line.Quantity {
					// Kurang stok - kita berikan peringatan atau log
					log.Printf("⚠️ Peringatan: Stok produk ID %d (%s) tersisa %.2f, pesanan membutuhkan %.2f",
						prod.ID, prod.DefaultCode, prod.StockQty, line.Quantity)
				}
			}
		}
	}

	order.State = "sale"
	if err := UpdateSaleOrderService(order); err != nil {
		return nil, err
	}

	// [FASE 3 ERP Integrasi] Otomatisasi reservasi stok & Dokumen Pengeluaran Barang / Surat Jalan (Delivery Order)
	deliveryPicking := inventory.StockPicking{
		Name:          fmt.Sprintf("WH/OUT/%s/%04d", time.Now().Format("2006"), order.ID),
		State:         "confirmed", // Menunggu disiapkan / dikirim di gudang
		ScheduledDate: time.Now().AddDate(0, 0, 3),
	}
	if order.PartnerID > 0 {
		pid := order.PartnerID
		deliveryPicking.PartnerID = &pid
	}
	_ = config.DB.Create(&deliveryPicking).Error

	// Kunci stok (Reserved Stock) & catat baris perpindahan fisik (StockMove)
	for _, line := range order.OrderLines {
		if line.ProductID > 0 {
			// Update kuota reserved di tabel produk
			config.DB.Model(&inventory.Product{}).Where("id = ?", line.ProductID).
				Update("reserved_qty", gorm.Expr("reserved_qty + ?", line.Quantity))

			// Buat baris stock move
			move := inventory.StockMove{
				Name:         fmt.Sprintf("Kirim: %s", line.Description),
				PickingID:    &deliveryPicking.ID,
				ProductID:    line.ProductID,
				Quantity:     line.Quantity,
				QuantityDone: line.DeliveredQty,
				State:        "confirmed",
			}
			_ = config.DB.Create(&move)
		}
	}

	log.Printf("📦 Stock Reserved & DO Created: Surat Jalan %s berhasil dibuat untuk Sales Order %s dengan status Confirmed",
		deliveryPicking.Name, order.Name)

	return order, nil
}

// CreateInvoiceFromSaleOrderService: One-Click Invoicing dari Sales Order ke Modul Finance Invoicing
func CreateInvoiceFromSaleOrderService(id uint) (*invoicing.Invoice, error) {
	order, err := GetSaleOrderByID(id)
	if err != nil {
		return nil, err
	}

	// Buat Faktur di Modul Finance
	inv := invoicing.Invoice{
		Name:           fmt.Sprintf("INV/%s/%04d", time.Now().Format("2006"), order.ID),
		PartnerID:      order.PartnerID,
		InvoiceDate:    time.Now(),
		DueDate:        time.Now().AddDate(0, 0, 30), // Jatuh tempo 30 hari (Net 30)
		State:          "draft",                      // Draft Faktur siap divalidasi Accounting
		AmountUntaxed:  order.AmountUntaxed,
		AmountTax:      order.AmountTax,
		AmountTotal:    order.AmountTotal,
		ResidualAmount: order.AmountTotal,
		CreatedAt:      time.Now(),
	}

	if err := config.DB.Create(&inv).Error; err != nil {
		return nil, err
	}

	// Salin baris produk ke finance.invoice_lines
	for _, l := range order.OrderLines {
		invLine := invoicing.InvoiceLine{
			InvoiceID:   inv.ID,
			Description: l.Description,
			Quantity:    l.Quantity,
			UnitPrice:   l.UnitPrice,
			SubTotal:    l.SubTotal,
		}
		_ = config.DB.Create(&invLine)
	}

	// Log integrasi
	log.Printf("💰 One-Click Invoicing: Invoice %s berhasil dibuat untuk Sales Order %s (Total: Rp %.2f)",
		inv.Name, order.Name, inv.AmountTotal)

	return &inv, nil
}

// PaymentLinkResult merepresentasikan respon link pembayaran
type PaymentLinkResult struct {
	OrderID     uint    `json:"order_id"`
	OrderName   string  `json:"order_name"`
	PaymentURL  string  `json:"payment_url"`
	GrossAmount float64 `json:"gross_amount"`
	ExpiredAt   string  `json:"expired_at"`
	Provider    string  `json:"provider"`
	SnapToken   string  `json:"snap_token"`
}

// GeneratePaymentLinkService: Membuat direct payment link (Midtrans / QRIS / Virtual Account)
func GeneratePaymentLinkService(id uint) (*PaymentLinkResult, error) {
	order, err := GetSaleOrderByID(id)
	if err != nil {
		return nil, err
	}

	// Buat simulated/direct payment link terintegrasi Midtrans / PG Gateway
	token := fmt.Sprintf("SNAP-%s-%d-%d", time.Now().Format("20060102"), order.ID, time.Now().Unix()%100000)
	paymentURL := fmt.Sprintf("https://app.sandbox.midtrans.com/snap/v2/vtweb/%s", token)

	// Tandai pada pesanan bahwa link pembayaran sudah dibuat/dikirim
	order.IsPaymentLinkSent = true
	_ = config.DB.Model(&SaleOrder{}).Where("id = ?", order.ID).Update("is_payment_link_sent", true)

	result := &PaymentLinkResult{
		OrderID:     order.ID,
		OrderName:   order.Name,
		PaymentURL:  paymentURL,
		GrossAmount: order.AmountTotal,
		ExpiredAt:   time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05"),
		Provider:    "Midtrans Enterprise / QRIS Instant",
		SnapToken:   token,
	}

	log.Printf("💳 Direct Payment Link generated for Order %s: %s", order.Name, paymentURL)
	return result, nil
}

func UpdateSaleOrderStatusService(id uint, newState string) (*SaleOrder, error) {
	order, err := GetSaleOrderByID(id)
	if err != nil {
		return nil, err
	}
	order.State = newState
	if newState == "sale" {
		return ConfirmSaleOrderService(id)
	}
	if err := UpdateSaleOrder(order); err != nil {
		return nil, err
	}
	return order, nil
}

// ApproveDiscountSaleOrderService handles manager approval / rejection for high discounts (>15%)
func ApproveDiscountSaleOrderService(id uint, status string, approver string) (*SaleOrder, error) {
	order, err := GetSaleOrderByID(id)
	if err != nil {
		return nil, err
	}
	if status != "Approved" && status != "Rejected" {
		status = "Approved"
	}
	order.ApprovalStatus = status
	if approver == "" {
		approver = "Sales Manager (Authorized)"
	}
	order.ApprovedBy = approver
	if status == "Approved" {
		order.NeedsApproval = false
	}
	if err := UpdateSaleOrder(order); err != nil {
		return nil, err
	}
	log.Printf("🛡️ Discount Approval for Order %s: %s by %s (Max Discount: %.1f%%)",
		order.Name, status, approver, order.MaxDiscount)
	return order, nil
}

// SignSaleOrderService handles customer digital e-signature on Quotation/Sales Order
func SignSaleOrderService(id uint, signerName string) (*SaleOrder, error) {
	order, err := GetSaleOrderByID(id)
	if err != nil {
		return nil, err
	}
	order.IsSigned = true
	if signerName == "" {
		signerName = order.CustomerName
	}
	order.CustomerName = signerName
	if err := UpdateSaleOrder(order); err != nil {
		return nil, err
	}
	log.Printf("✍️ Digital E-Signature verified for Order %s by %s", order.Name, signerName)
	return order, nil
}

func DeleteSaleOrderService(id uint) error { return DeleteSaleOrder(id) }

func CreatePricelistService(data *Pricelist) error        { return CreatePricelist(data) }
func GetAllPricelistService() ([]Pricelist, error)        { return GetAllPricelist() }
func GetPricelistByIDService(id uint) (*Pricelist, error) { return GetPricelistByID(id) }
func UpdatePricelistService(data *Pricelist) error        { return UpdatePricelist(data) }
func DeletePricelistService(id uint) error                { return DeletePricelist(id) }

func CreatePricelistItemService(data *PricelistItem) error        { return CreatePricelistItem(data) }
func GetAllPricelistItemService() ([]PricelistItem, error)        { return GetAllPricelistItem() }
func GetPricelistItemByIDService(id uint) (*PricelistItem, error) { return GetPricelistItemByID(id) }
func UpdatePricelistItemService(data *PricelistItem) error        { return UpdatePricelistItem(data) }
func DeletePricelistItemService(id uint) error                    { return DeletePricelistItem(id) }

func CreateQuotationTemplateService(data *QuotationTemplate) error {
	return CreateQuotationTemplate(data)
}
func GetAllQuotationTemplateService() ([]QuotationTemplate, error) { return GetAllQuotationTemplate() }
func GetQuotationTemplateByIDService(id uint) (*QuotationTemplate, error) {
	return GetQuotationTemplateByID(id)
}
func UpdateQuotationTemplateService(data *QuotationTemplate) error {
	return UpdateQuotationTemplate(data)
}
func DeleteQuotationTemplateService(id uint) error { return DeleteQuotationTemplate(id) }

func CreateDeliveryMethodService(data *DeliveryMethod) error        { return CreateDeliveryMethod(data) }
func GetAllDeliveryMethodService() ([]DeliveryMethod, error)        { return GetAllDeliveryMethod() }
func GetDeliveryMethodByIDService(id uint) (*DeliveryMethod, error) { return GetDeliveryMethodByID(id) }
func UpdateDeliveryMethodService(data *DeliveryMethod) error        { return UpdateDeliveryMethod(data) }
func DeleteDeliveryMethodService(id uint) error                     { return DeleteDeliveryMethod(id) }

func CreateSaleOrderLineService(data *SaleOrderLine) error {
	data.SubTotal = data.Quantity * data.UnitPrice * (1 - (data.Discount / 100))
	if err := CreateSaleOrderLine(data); err != nil {
		return err
	}
	return RecalculateSaleOrder(data.OrderID)
}
func GetAllSaleOrderLineService() ([]SaleOrderLine, error)        { return GetAllSaleOrderLine() }
func GetSaleOrderLineByIDService(id uint) (*SaleOrderLine, error) { return GetSaleOrderLineByID(id) }
func UpdateSaleOrderLineService(data *SaleOrderLine) error {
	data.SubTotal = data.Quantity * data.UnitPrice * (1 - (data.Discount / 100))
	if err := UpdateSaleOrderLine(data); err != nil {
		return err
	}
	return RecalculateSaleOrder(data.OrderID)
}
func DeleteSaleOrderLineService(id uint) error {
	line, err := GetSaleOrderLineByID(id)
	if err != nil {
		return err
	}
	orderID := line.OrderID
	if err := DeleteSaleOrderLine(id); err != nil {
		return err
	}
	return RecalculateSaleOrder(orderID)
}

func RecalculateSaleOrder(orderID uint) error {
	var lines []SaleOrderLine
	if err := config.DB.Where("order_id = ?", orderID).Find(&lines).Error; err != nil {
		return err
	}

	var amountUntaxed float64
	for _, line := range lines {
		amountUntaxed += line.SubTotal
	}

	amountTax := amountUntaxed * 0.11 // Asumsi PPN 11%
	amountTotal := amountUntaxed + amountTax

	return config.DB.Model(&SaleOrder{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"amount_untaxed": amountUntaxed,
		"amount_tax":     amountTax,
		"amount_total":   amountTotal,
	}).Error
}

type SalesLeaderboardItem struct {
	SalespersonName   string  `json:"salesperson_name"`
	TotalOrders       int64   `json:"total_orders"`
	TotalRevenue      float64 `json:"total_revenue"`
	TotalCommission   float64 `json:"total_commission"`
	ConfirmedDeals    int64   `json:"confirmed_deals"`
	TargetRevenue     float64 `json:"target_revenue"`
	AchievementPct    float64 `json:"achievement_pct"`
}

// GetSalesLeaderboardService mengagregasi data kinerja dan komisi setiap salesperson
func GetSalesLeaderboardService() ([]SalesLeaderboardItem, error) {
	var results []SalesLeaderboardItem

	type QueryResult struct {
		SalespersonName string  `gorm:"column:salesperson_name"`
		TotalOrders     int64   `gorm:"column:total_orders"`
		TotalRevenue    float64 `gorm:"column:total_revenue"`
		TotalCommission float64 `gorm:"column:total_commission"`
		ConfirmedDeals  int64   `gorm:"column:confirmed_deals"`
	}

	var rows []QueryResult
	err := config.DB.Model(&SaleOrder{}).
		Select("salesperson_name, COUNT(id) as total_orders, COALESCE(SUM(amount_untaxed), 0) as total_revenue, COALESCE(SUM(commission_amount), 0) as total_commission, COUNT(CASE WHEN state = 'sale' OR state = 'done' THEN 1 END) as confirmed_deals").
		Group("salesperson_name").
		Order("total_revenue DESC").
		Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	targetRevenue := 100000000.0 // Default KPI Target Rp 100jt per sales
	for _, r := range rows {
		name := r.SalespersonName
		if name == "" {
			name = "Sales Team"
		}
		achieve := 0.0
		if targetRevenue > 0 {
			achieve = (r.TotalRevenue / targetRevenue) * 100.0
		}
		results = append(results, SalesLeaderboardItem{
			SalespersonName: name,
			TotalOrders:     r.TotalOrders,
			TotalRevenue:    r.TotalRevenue,
			TotalCommission: r.TotalCommission,
			ConfirmedDeals:  r.ConfirmedDeals,
			TargetRevenue:   targetRevenue,
			AchievementPct:  achieve,
		})
	}

	return results, nil
}

// BypassCreditHoldService memberikan wewenang kepada Finance Manager untuk melepaskan limit kredit tertahan
func BypassCreditHoldService(orderID uint, managerName string) (*SaleOrder, error) {
	order, err := GetSaleOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	if managerName == "" {
		managerName = "Finance Manager"
	}
	order.IsCreditBypassed = true
	order.BypassedBy = managerName
	if err := config.DB.Model(&SaleOrder{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"is_credit_bypassed": true,
		"bypassed_by":         managerName,
	}).Error; err != nil {
		return nil, err
	}
	log.Printf("🛡️ Credit Limit Bypassed: Sales Order %s disetujui untuk bypass limit oleh %s", order.Name, managerName)
	return order, nil
}

// ExportEFakturCSVService menghasilkan format CSV E-Faktur DJP resmi siap impor
func ExportEFakturCSVService(orderID uint) (string, error) {
	order, err := GetSaleOrderByID(orderID)
	if err != nil {
		return "", err
	}

	npwp := "00.000.000.0-000.000"
	if order.Partner != nil && order.Partner.Vat != "" {
		npwp = order.Partner.Vat
	}

	custName := order.CustomerName
	if custName == "" && order.Partner != nil {
		custName = order.Partner.Name
	}

	nsfp := order.Nsfp
	if nsfp == "" {
		nsfp = fmt.Sprintf("010.002-26.%08d", order.ID)
	}

	// Baris Header Resmi E-Faktur DJP Faktur Keluaran (FK)
	headerFK := "FK,KD_JENIS_TRANSAKSI,FG_PENGGANTI,NOMOR_FAKTUR,MASA_PAJAK,TAHUN_PAJAK,TANGGAL_FAKTUR,NPWP,NAMA,ALAMAT_LENGKAP,JUMLAH_DPP,JUMLAH_PPN,JUMLAH_PPNBM,ID_KETERANGAN_TAMBAHAN,FG_UANG_MUKA,UANG_MUKA_DPP,UANG_MUKA_PPN,JUMLAH_PPNBM_DIBEBASKAN,REFERENSI\n"
	headerOF := "OF,KODE_OBJEK,NAMA,HARGA_SATUAN,JUMLAH_BARANG,HARGA_TOTAL,DISKON,DPP,PPN,TARIF_PPNBM,PPNBM\n"

	now := order.DateOrder.Format("02/01/2006")
	masaPajak := order.DateOrder.Format("01")
	tahunPajak := order.DateOrder.Format("2006")

	rowFK := fmt.Sprintf("FK,01,0,%s,%s,%s,%s,%s,%s,Indonesia,%.2f,%.2f,0,0,0,0,0,0,%s\n",
		nsfp, masaPajak, tahunPajak, now, npwp, custName, order.AmountUntaxed, order.AmountTax, order.Name)

	var rowsOF string
	for _, l := range order.OrderLines {
		dpp := l.SubTotal
		ppn := dpp * (order.TaxRate / 100.0)
		itemCode := fmt.Sprintf("BRG-%d", l.ProductID)
		rowsOF += fmt.Sprintf("OF,%s,%s,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f,0,0\n",
			itemCode, l.Description, l.UnitPrice, l.Quantity, l.Quantity*l.UnitPrice, (l.Quantity*l.UnitPrice)-dpp, dpp, ppn)
	}

	return headerFK + rowFK + headerOF + rowsOF, nil
}

// FASE 3: Ambil daftar Surat Jalan / Delivery Order yang terhubung ke Sales Order
func GetDeliveryOrdersBySaleOrderIDService(orderID uint) ([]inventory.StockPicking, error) {
	var pickings []inventory.StockPicking
	pickingPattern := fmt.Sprintf("%%%04d", orderID)
	err := config.DB.Where("name LIKE ?", pickingPattern).Find(&pickings).Error
	return pickings, err
}

// DeliverSaleOrderService memproses penyerahan barang / Surat Jalan (Partial atau Full Delivery)
func DeliverSaleOrderService(orderID uint, lineID uint, deliveredQty float64) (*SaleOrder, error) {
	order, err := GetSaleOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	for i := range order.OrderLines {
		if order.OrderLines[i].ID == lineID || (lineID == 0 && order.OrderLines[i].DeliveredQty < order.OrderLines[i].Quantity) {
			qtyToDeliver := deliveredQty
			if qtyToDeliver <= 0 {
				qtyToDeliver = order.OrderLines[i].Quantity - order.OrderLines[i].DeliveredQty
			}

			// Kurangi stok fisik nyata dan lepaskan reservasi
			if order.OrderLines[i].ProductID > 0 {
				config.DB.Model(&inventory.Product{}).Where("id = ?", order.OrderLines[i].ProductID).
					Updates(map[string]interface{}{
						"stock_qty":    gorm.Expr("stock_qty - ?", qtyToDeliver),
						"reserved_qty": gorm.Expr("GREATEST(0, reserved_qty - ?)", qtyToDeliver),
					})
			}

			order.OrderLines[i].DeliveredQty += qtyToDeliver
			config.DB.Model(&SaleOrderLine{}).Where("id = ?", order.OrderLines[i].ID).
				Update("delivered_qty", order.OrderLines[i].DeliveredQty)
		}
	}

	// Update status Surat Jalan menjadi done jika semua item terkirim
	allDone := true
	for _, l := range order.OrderLines {
		if l.DeliveredQty < l.Quantity {
			allDone = false
			break
		}
	}

	pickingPattern := fmt.Sprintf("%%%04d", orderID)
	if allDone {
		config.DB.Model(&inventory.StockPicking{}).Where("name LIKE ?", pickingPattern).
			Update("state", "done")
	}

	return order, nil
}


