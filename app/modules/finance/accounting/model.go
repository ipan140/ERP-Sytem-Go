package accounting

import (
	"ERP-System/config"
	"time"
)

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(50);not null;unique" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"` // e.g. receivable, payable, bank, income, expense
	CreatedAt time.Time `json:"created_at"`
}

type Journal struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"type:varchar(10);not null;unique" json:"code"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Type string `gorm:"type:varchar(50);not null" json:"type"` // sale, purchase, cash, bank, general
}

type JournalEntry struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. INV/2026/001
	JournalID uint      `json:"journal_id"`
	Date      time.Time `json:"date"`
	State     string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, posted
	CreatedAt time.Time `json:"created_at"`
}

type JournalItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	EntryID   uint    `json:"entry_id"`
	AccountID uint    `json:"account_id"`
	Name      string  `gorm:"type:varchar(255)" json:"name"`
	Debit     float64 `gorm:"type:numeric(15,2);default:0" json:"debit"`
	Credit    float64 `gorm:"type:numeric(15,2);default:0" json:"credit"`
}

// Rekonsiliasi Bank
type BankStatement struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. Mutasi Agustus
	Date      time.Time `json:"date"`
	State     string    `gorm:"type:varchar(20);default:'open'" json:"state"` // open, validated
	CreatedAt time.Time `json:"created_at"`
}

type BankStatementLine struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StatementID uint      `json:"statement_id"`
	Date        time.Time `json:"date"`
	PaymentRef  string    `gorm:"type:varchar(255)" json:"payment_ref"`
	Amount      float64   `gorm:"type:numeric(15,2)" json:"amount"` // Positif (Uang Masuk), Negatif (Keluar)
}

// Akuntansi Analitik (Cost Center)
type AnalyticAccount struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Proyek A, Divisi Marketing
}

type AnalyticLine struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	AccountID uint    `json:"account_id"`
	Name      string  `gorm:"type:varchar(255)" json:"name"`
	Amount    float64 `gorm:"type:numeric(15,2)" json:"amount"` // Biaya atau Pendapatan Analitik
}

// ---- FITUR ODOO MEDIUM (Ditambahkan Manual) ----

type AccountReconcileModel struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
	MatchText   string  `gorm:"type:varchar(255)" json:"match_text"` // Kata kunci mutasi bank
	AccountID   uint    `json:"account_id"`                          // Akun tujuan
	AmountType  string  `gorm:"type:varchar(20);default:'percentage'" json:"amount_type"`
	AmountValue float64 `gorm:"type:numeric(15,2);default:100" json:"amount_value"`
}

type FollowupRule struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(100);not null" json:"name"` // e.g. Peringatan 1
	DelayDays  int    `json:"delay_days"`                             // Telat berapa hari
	SendEmail  bool   `gorm:"default:true" json:"send_email"`
	SendLetter bool   `gorm:"default:false" json:"send_letter"`
}

type AccountLockDate struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	CompanyID          uint      `json:"company_id"`
	FiscalYearLockDate time.Time `json:"fiscal_year_lock_date"` // Kunci Mutlak
	TaxLockDate        time.Time `json:"tax_lock_date"`         // Kunci Pajak Saja
}

// ---- FITUR PAYMENT GATEWAY (Ditambahkan Manual) ----

type PaymentAcquirer struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`       // e.g. Midtrans, Stripe
	Provider string `gorm:"type:varchar(50);not null" json:"provider"`    // midtrans, stripe, paypal
	State    string `gorm:"type:varchar(20);default:'test'" json:"state"` // test, enabled, disabled
	APIKey   string `gorm:"type:varchar(255)" json:"api_key"`             // Secret Key dari Provider
}

type PaymentTransaction struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	AcquirerID           uint      `json:"acquirer_id"`
	Reference            string    `gorm:"type:varchar(100);unique" json:"reference"` // Nomor referensi internal
	Amount               float64   `gorm:"type:numeric(15,2);not null" json:"amount"`
	PartnerID            uint      `json:"partner_id"`                                      // Pelanggan yang membayar
	InvoiceID            *uint     `json:"invoice_id"`                                      // Tagihan yang dibayar
	State                string    `gorm:"type:varchar(20);default:'draft'" json:"state"`   // draft, pending, authorized, done, error
	GatewayTransactionID string    `gorm:"type:varchar(100)" json:"gateway_transaction_id"` // ID unik dari Midtrans/Stripe
	CreatedAt            time.Time `json:"created_at"`
}

type AccountIncoterms struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"type:varchar(3);not null;unique" json:"code"` // EXW, FOB, CIF
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Account{}, &Journal{}, &JournalEntry{}, &JournalItem{}, &BankStatement{}, &BankStatementLine{}, &AnalyticAccount{}, &AnalyticLine{}, &AccountReconcileModel{}, &FollowupRule{}, &AccountLockDate{}, &PaymentAcquirer{}, &PaymentTransaction{}, &AccountIncoterms{})
}

type Payment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	InvoiceID   uint      `json:"invoice_id"`
	Amount      float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	CreatedAt   time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Payment{})
}

func SeedDefaultAccounts() {
	// Pengecekan apakah tabel akun sudah ada isinya
	var count int64
	config.DB.Model(&Account{}).Count(&count)
	if count == 0 {
		defaultAccounts := []Account{
			{Code: "1-1001", Name: "Kas Besar", Type: "bank"},
			{Code: "1-1002", Name: "Bank BCA", Type: "bank"},
			{Code: "1-2001", Name: "Piutang Usaha", Type: "receivable"},
			{Code: "2-1001", Name: "Hutang Usaha", Type: "payable"},
			{Code: "4-1001", Name: "Pendapatan Penjualan", Type: "income"},
			{Code: "6-1001", Name: "Beban Gaji", Type: "expense"},
		}
		for _, acc := range defaultAccounts {
			config.DB.Create(&acc)
		}
	}
}

type FiscalPosition struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Ekspor Luar Negeri
	CountryID uint   `json:"country_id"`
	TaxSrcID  uint   `json:"tax_src_id"`  // Pajak Asal (e.g. PPN 11%)
	TaxDestID uint   `json:"tax_dest_id"` // Pajak Tujuan (e.g. PPN 0%)
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &FiscalPosition{})
}

// 1. Multi-Currency
type Currency struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(50);not null" json:"name"` // e.g. USD, EUR, IDR
	Symbol string `gorm:"type:varchar(10)" json:"symbol"`        // e.g. $, Rp
}

type ExchangeRate struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CurrencyID uint      `json:"currency_id"`
	Date       time.Time `json:"date"`
	Rate       float64   `gorm:"type:numeric(15,6)" json:"rate"` // Nilai tukar terhadap mata uang dasar perusahaan
}

// 2. Fixed Asset Depreciation
type Asset struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Name               string    `gorm:"type:varchar(255);not null" json:"name"`
	OriginalValue      float64   `gorm:"type:numeric(15,2)" json:"original_value"`
	SalvageValue       float64   `gorm:"type:numeric(15,2)" json:"salvage_value"`
	DepreciationMethod string    `gorm:"type:varchar(50);default:'linear'" json:"depreciation_method"` // linear, degressive
	DurationMonths     int       `json:"duration_months"`
	StartDate          time.Time `json:"start_date"`
	State              string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, open, close
}

type AssetDepreciationLine struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	AssetID          uint      `json:"asset_id"`
	DepreciationDate time.Time `json:"depreciation_date"`
	Amount           float64   `gorm:"type:numeric(15,2)" json:"amount"`
	IsPosted         bool      `gorm:"default:false" json:"is_posted"`
}

// 3. Deferred Revenues
type DeferredRevenue struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	TotalAmount    float64   `gorm:"type:numeric(15,2)" json:"total_amount"`
	DurationMonths int       `json:"duration_months"`
	StartDate      time.Time `json:"start_date"`
	State          string    `gorm:"type:varchar(20);default:'draft'" json:"state"`
}

type DeferredRevenueLine struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	DeferredRevenueID uint      `json:"deferred_revenue_id"`
	Date              time.Time `json:"date"`
	Amount            float64   `gorm:"type:numeric(15,2)" json:"amount"`
	IsPosted          bool      `gorm:"default:false" json:"is_posted"`
}

// 4. Budgeting
type Budget struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	State     string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, validated, done
}

type BudgetLine struct {
	ID                uint    `gorm:"primaryKey" json:"id"`
	BudgetID          uint    `json:"budget_id"`
	AnalyticAccountID uint    `json:"analytic_account_id"` // Pusat Biaya
	PlannedAmount     float64 `gorm:"type:numeric(15,2)" json:"planned_amount"`
}

func (Account) TableName() string {
	return "finance.accounts"
}
func (Journal) TableName() string {
	return "finance.journals"
}
func (JournalEntry) TableName() string {
	return "finance.journal_entrys"
}
func (JournalItem) TableName() string {
	return "finance.journal_items"
}
func (BankStatement) TableName() string {
	return "finance.bank_statements"
}
func (BankStatementLine) TableName() string {
	return "finance.bank_statement_lines"
}
func (AnalyticAccount) TableName() string {
	return "finance.analytic_accounts"
}
func (AnalyticLine) TableName() string {
	return "finance.analytic_lines"
}
func (AccountReconcileModel) TableName() string {
	return "finance.account_reconcile_models"
}
func (FollowupRule) TableName() string {
	return "finance.followup_rules"
}
func (AccountLockDate) TableName() string {
	return "finance.account_lock_dates"
}
func (PaymentAcquirer) TableName() string {
	return "finance.payment_acquirers"
}
func (PaymentTransaction) TableName() string {
	return "finance.payment_transactions"
}
func (AccountIncoterms) TableName() string {
	return "finance.account_incotermses"
}
func (Payment) TableName() string {
	return "finance.payments"
}
func (FiscalPosition) TableName() string {
	return "finance.fiscal_positions"
}
func (Currency) TableName() string {
	return "setting.currencies"
}
func (ExchangeRate) TableName() string {
	return "finance.exchange_rates"
}
func (Asset) TableName() string {
	return "finance.assets"
}
func (AssetDepreciationLine) TableName() string {
	return "finance.asset_depreciation_lines"
}
func (DeferredRevenue) TableName() string {
	return "finance.deferred_revenues"
}
func (DeferredRevenueLine) TableName() string {
	return "finance.deferred_revenue_lines"
}
func (Budget) TableName() string {
	return "finance.budgets"
}
func (BudgetLine) TableName() string {
	return "finance.budget_lines"
}