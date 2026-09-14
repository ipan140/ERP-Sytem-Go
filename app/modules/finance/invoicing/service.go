package invoicing

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"ERP-System/app/modules/finance/accounting"
	"ERP-System/config"
	"ERP-System/pkg/rabbitmq"
)

type DocumentPayload struct {
	DocumentID   uint   `json:"document_id"`
	DocumentType string `json:"document_type"` // e.g., "invoice", "payslip", "quotation"
	Format       string `json:"format"`        // "pdf" or "excel"
	UserID       uint   `json:"user_id"`
}

func GenerateInvoicePDFService(invoiceID uint, userID uint) error {
	if rabbitmq.Channel != nil {
		req := DocumentPayload{
			DocumentID:   invoiceID,
			DocumentType: "invoice",
			Format:       "pdf",
			UserID:       userID,
		}
		body, _ := json.Marshal(req)
		err := rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("📄 Event RabbitMQ: Generate PDF Invoice %d dikirim ke antrean!", invoiceID)
		return err
	}
	return nil
}

func CreateInvoiceService(data *Invoice) error {
	return CreateInvoice(data)
}

func GetAllInvoiceService() ([]Invoice, error) {
	return GetAllInvoice()
}

func GetPaginatedInvoiceService(offset int, limit int, search string, status string) ([]Invoice, int64, error) {
	return GetPaginatedInvoices(offset, limit, search, status)
}

func GetInvoiceByIDService(id uint) (*Invoice, error) {
	return GetInvoiceByID(id)
}

func UpdateInvoiceService(data *Invoice) error {
	return UpdateInvoice(data)
}

func DeleteInvoiceService(id uint) error {
	return DeleteInvoice(id)
}

// PostInvoiceService mensimulasikan logika Odoo saat Faktur dikonfirmasi.
// Ini akan mengubah status faktur dan otomatis membuat Jurnal Entri di Akuntansi.
func PostInvoiceService(invoiceID uint) error {
	invoice, err := GetInvoiceByID(invoiceID)
	if err != nil {
		return err
	}
	if invoice.State == "posted" {
		return errors.New("faktur sudah diposting sebelumnya")
	}

	// Validasi Penganggaran (Budgeting Validator)
	// Jika faktur ini terkait dengan AnalyticAccount tertentu, kita cek apakah biayanya sudah melebihi Budget
	// (Simulasi Sederhana)
	var budgetLine accounting.BudgetLine
	// Misal faktur ini dibebankan ke analytic_account_id = 1 (Proyek A)
	err = config.DB.Where("analytic_account_id = ?", 1).First(&budgetLine).Error
	if err == nil {
		// Jika budget ditemukan, kita cek apakah AmountTotal melebihi sisa budget (simulasi logik sederhana: planned_amount vs invoice amount)
		if invoice.AmountTotal > budgetLine.PlannedAmount {
			return errors.New("Validasi Gagal: Faktur ini melebihi ambang batas Budget yang ditetapkan (Budgeting Alert)")
		}
	}

	// 1. Ubah status Faktur
	invoice.State = "posted"
	if err := UpdateInvoice(invoice); err != nil {
		return err
	}

	// 2. Memicu Pembuatan Jurnal di Modul Akuntansi (Auto-Journaling)
	// Kita asumsikan ada Jurnal Penjualan (JournalID = 1)
	entry := &accounting.JournalEntry{
		Name:      invoice.Name, // Nama jurnal sama dengan nama faktur
		JournalID: 1,
		Date:      time.Now(),
		State:     "posted",
	}

	// Buat Entry Induk
	if err := accounting.CreateJournalEntryService(entry); err != nil {
		return err
	}

	// Buat Baris Jurnal (Double-Entry: Debit Piutang, Kredit Pendapatan)
	// Asumsi AccountID 1 = Piutang, AccountID 2 = Pendapatan
	debitItem := &accounting.JournalItem{
		EntryID:   entry.ID,
		AccountID: 1,
		Name:      "Piutang Pelanggan - " + invoice.Name,
		Debit:     invoice.AmountTotal,
		Credit:    0,
	}

	creditItem := &accounting.JournalItem{
		EntryID:   entry.ID,
		AccountID: 2,
		Name:      "Pendapatan Penjualan - " + invoice.Name,
		Debit:     0,
		Credit:    invoice.AmountTotal,
	}

	accounting.CreateJournalItemService(debitItem)
	accounting.CreateJournalItemService(creditItem)

	return nil
}

// RefundInvoiceService membuat Credit Note (Faktur Minus) dari Faktur yang sudah ada
func RefundInvoiceService(invoiceID uint) error {
	invoice, err := GetInvoiceByID(invoiceID)
	if err != nil {
		return err
	}
	if invoice.State != "posted" {
		return errors.New("hanya faktur yang sudah diposting yang bisa di-refund")
	}

	// Buat Jurnal Pembalik di modul Akuntansi (Negative Amounts)
	entry := &accounting.JournalEntry{
		Name:      "REFUND - " + invoice.Name,
		JournalID: 1,
		Date:      time.Now(),
		State:     "posted",
	}

	if err := accounting.CreateJournalEntryService(entry); err != nil {
		return err
	}

	// Balik posisi jurnal awal (Debit jadi Kredit, Kredit jadi Debit)
	// Atau mencatat minus. Di Odoo biasanya membalikkan akun.
	debitItem := &accounting.JournalItem{
		EntryID:   entry.ID,
		AccountID: 1,
		Name:      "Piutang Pelanggan (Reversal) - " + invoice.Name,
		Debit:     0,
		Credit:    invoice.AmountTotal,
	}

	creditItem := &accounting.JournalItem{
		EntryID:   entry.ID,
		AccountID: 2,
		Name:      "Pendapatan Penjualan (Reversal) - " + invoice.Name,
		Debit:     invoice.AmountTotal,
		Credit:    0,
	}

	accounting.CreateJournalItemService(debitItem)
	accounting.CreateJournalItemService(creditItem)

	// Set faktur menjadi cancelled
	invoice.State = "cancelled"
	UpdateInvoice(invoice)

	return nil
}

// RunDunningProcess menyapu faktur yang telat bayar dan menaikkan level peringatannya.
func RunDunningProcess() error {
	var invoices []Invoice
	// Cari faktur yang diposting, belum lunas, dan melewati batas waktu
	config.DB.Where("state = ? AND residual_amount > 0 AND due_date < ?", "posted", time.Now()).Find(&invoices)

	for _, inv := range invoices {
		inv.FollowUpLevel += 1

		// Simulasi ngirim email ancaman
		// fmt.Printf("Mengirim Peringatan Level %d ke Klien untuk Faktur %s\n", inv.FollowUpLevel, inv.Name)

		config.DB.Save(&inv)
	}
	return nil
}

// AutoSwapTax mensimulasikan Fiscal Position. Jika klien dari luar negeri, pajaknya di-Nol-kan.
func AutoSwapTax(invoiceLine *InvoiceLine, isForeignCustomer bool) {
	if isForeignCustomer {
		// Asumsi ID pajak 0% adalah 2
		zeroTax := uint(2)
		invoiceLine.TaxID = &zeroTax
	}
}

func CreatePaymentTermLineService(data *PaymentTermLine) error { return CreatePaymentTermLine(data) }
func GetAllPaymentTermLineService() ([]PaymentTermLine, error) { return GetAllPaymentTermLine() }
func GetPaymentTermLineByIDService(id uint) (*PaymentTermLine, error) {
	return GetPaymentTermLineByID(id)
}
func UpdatePaymentTermLineService(data *PaymentTermLine) error { return UpdatePaymentTermLine(data) }
func DeletePaymentTermLineService(id uint) error               { return DeletePaymentTermLine(id) }

func CreateTaxRepartitionLineService(data *TaxRepartitionLine) error {
	return CreateTaxRepartitionLine(data)
}
func GetAllTaxRepartitionLineService() ([]TaxRepartitionLine, error) {
	return GetAllTaxRepartitionLine()
}
func GetTaxRepartitionLineByIDService(id uint) (*TaxRepartitionLine, error) {
	return GetTaxRepartitionLineByID(id)
}
func UpdateTaxRepartitionLineService(data *TaxRepartitionLine) error {
	return UpdateTaxRepartitionLine(data)
}
func DeleteTaxRepartitionLineService(id uint) error { return DeleteTaxRepartitionLine(id) }
