package accounting

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

// --- Account (COA) Repository ---
func CreateAccount(data *Account) error {
	return config.DB.Create(data).Error
}

func GetAllAccounts() ([]Account, error) {
	var list []Account
	err := config.DB.Order("code asc").Find(&list).Error
	return list, err
}

func GetPaginatedAccounts(offset int, limit int, search string, category string) ([]Account, int64, error) {
	var list []Account
	var total int64

	query := config.DB.Model(&Account{})

	if category != "" && category != "all" && category != "All" && category != "Semua" {
		query = query.Where("type = ? OR category ILIKE ?", category, "%"+category+"%")
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("code ILIKE ? OR name ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("code asc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetAccountByID(id uint) (*Account, error) {
	var data Account
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateAccount(data *Account) error {
	return config.DB.Save(data).Error
}

func DeleteAccount(id uint) error {
	return config.DB.Delete(&Account{}, id).Error
}

func SeedStandardIndonesianCOA() error {
	standardAccounts := []Account{
		// 1 - ASET (HARTA)
		{Code: "1-1000", Name: "Kas & Setara Kas", Type: "asset", Category: "Aset Lancar", Balance: 0, IsActive: true},
		{Code: "1-1001", Name: "Kas Kecil (Petty Cash)", Type: "asset", Category: "Aset Lancar", Balance: 15000000, IsActive: true},
		{Code: "1-1002", Name: "Rekening Operasional BCA", Type: "asset", Category: "Aset Lancar", Balance: 450000000, IsActive: true},
		{Code: "1-1003", Name: "Rekening Penggajian Mandiri", Type: "asset", Category: "Aset Lancar", Balance: 250000000, IsActive: true},
		{Code: "1-1200", Name: "Piutang Usaha (Trade Receivables)", Type: "asset", Category: "Aset Lancar", Balance: 125000000, IsActive: true},
		{Code: "1-1300", Name: "Persediaan Barang Dagang", Type: "asset", Category: "Aset Lancar", Balance: 80000000, IsActive: true},
		{Code: "1-1400", Name: "Piutang Kasbon Karyawan", Type: "asset", Category: "Aset Lancar", Balance: 12000000, IsActive: true},
		{Code: "1-2000", Name: "Aset Tetap (Peralatan & Kantor)", Type: "asset", Category: "Aset Tidak Lancar", Balance: 350000000, IsActive: true},
		{Code: "1-2001", Name: "Akumulasi Penyusutan Aset Tetap", Type: "asset", Category: "Aset Tidak Lancar", Balance: -35000000, IsActive: true},

		// 2 - KEWAJIBAN (HUTANG)
		{Code: "2-1000", Name: "Hutang Usaha (Trade Payables)", Type: "liability", Category: "Hutang Jangka Pendek", Balance: 75000000, IsActive: true},
		{Code: "2-1100", Name: "Hutang Gaji Karyawan", Type: "liability", Category: "Hutang Jangka Pendek", Balance: 0, IsActive: true},
		{Code: "2-1200", Name: "Hutang Pajak PPh 21 Karyawan", Type: "liability", Category: "Hutang Jangka Pendek", Balance: 4500000, IsActive: true},
		{Code: "2-1300", Name: "Hutang Pajak PPN Keluaran", Type: "liability", Category: "Hutang Jangka Pendek", Balance: 12500000, IsActive: true},

		// 3 - EKUITAS (MODAL)
		{Code: "3-1000", Name: "Modal Disetor", Type: "equity", Category: "Ekuitas", Balance: 1000000000, IsActive: true},
		{Code: "3-2000", Name: "Laba Ditahan (Retained Earnings)", Type: "equity", Category: "Ekuitas", Balance: 170000000, IsActive: true},

		// 4 - PENDAPATAN
		{Code: "4-1000", Name: "Pendapatan Penjualan Produk", Type: "income", Category: "Pendapatan Usaha", Balance: 650000000, IsActive: true},
		{Code: "4-2000", Name: "Pendapatan Jasa & Layanan", Type: "income", Category: "Pendapatan Usaha", Balance: 180000000, IsActive: true},

		// 5 - HARGA POKOK PENJUALAN (HPP)
		{Code: "5-1000", Name: "Harga Pokok Penjualan (HPP)", Type: "expense", Category: "Beban Pokok", Balance: 320000000, IsActive: true},

		// 6 - BEBAN OPERASIONAL & UMUM
		{Code: "6-1000", Name: "Beban Gaji & Upah Karyawan", Type: "expense", Category: "Beban Personalia", Balance: 145000000, IsActive: true},
		{Code: "6-1001", Name: "Beban THR & Bonus Karyawan", Type: "expense", Category: "Beban Personalia", Balance: 25000000, IsActive: true},
		{Code: "6-1002", Name: "Beban Lembur Karyawan", Type: "expense", Category: "Beban Personalia", Balance: 8500000, IsActive: true},
		{Code: "6-2000", Name: "Beban Klaim & Reimbursement", Type: "expense", Category: "Beban Operasional", Balance: 14200000, IsActive: true},
		{Code: "6-2100", Name: "Beban Sewa Gedung & Kantor", Type: "expense", Category: "Beban Operasional", Balance: 30000000, IsActive: true},
		{Code: "6-2200", Name: "Beban Utilitas (Listrik, Air, Internet)", Type: "expense", Category: "Beban Operasional", Balance: 7500000, IsActive: true},
		{Code: "6-3000", Name: "Beban Penyusutan Aset", Type: "expense", Category: "Beban Non-Kas", Balance: 5000000, IsActive: true},
	}

	for _, acc := range standardAccounts {
		var existing Account
		err := config.DB.Where("code = ?", acc.Code).First(&existing).Error
		if err != nil { // Not found, create it
			config.DB.Create(&acc)
		} else if existing.Balance == 0 {
		    // Update balance to make the dummy data look good if it was 0
		    config.DB.Model(&existing).Update("balance", acc.Balance)
		}
	}
	return nil
}

func CreateJournalEntry(data *JournalEntry) error {
	return config.DB.Create(data).Error
}

func GetAllJournalEntry() ([]JournalEntry, error) {
	var list []JournalEntry
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedJournalEntries(offset int, limit int, search string, status string) ([]JournalEntry, int64, error) {
	var list []JournalEntry
	var total int64

	query := config.DB.Model(&JournalEntry{}).Preload(clause.Associations)

	if status != "" && status != "all" && status != "All" && status != "Semua" {
		query = query.Where("state = ?", status)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ?", s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetJournalEntryByID(id uint) (*JournalEntry, error) {
	var data JournalEntry
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateJournalEntry(data *JournalEntry) error {
	return config.DB.Save(data).Error
}

func DeleteJournalEntry(id uint) error {
	return config.DB.Delete(&JournalEntry{}, id).Error
}

func CreateJournalItem(data *JournalItem) error {
	return config.DB.Create(data).Error
}

func CreateAccountReconcileModel(data *AccountReconcileModel) error {
	return config.DB.Create(data).Error
}
func GetAllAccountReconcileModel() ([]AccountReconcileModel, error) {
	var list []AccountReconcileModel
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetAccountReconcileModelByID(id uint) (*AccountReconcileModel, error) {
	var data AccountReconcileModel
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateAccountReconcileModel(data *AccountReconcileModel) error {
	return config.DB.Save(data).Error
}
func DeleteAccountReconcileModel(id uint) error {
	return config.DB.Delete(&AccountReconcileModel{}, id).Error
}

func CreateFollowupRule(data *FollowupRule) error { return config.DB.Create(data).Error }
func GetAllFollowupRule() ([]FollowupRule, error) {
	var list []FollowupRule
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetFollowupRuleByID(id uint) (*FollowupRule, error) {
	var data FollowupRule
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateFollowupRule(data *FollowupRule) error { return config.DB.Save(data).Error }
func DeleteFollowupRule(id uint) error            { return config.DB.Delete(&FollowupRule{}, id).Error }

func CreateAccountLockDate(data *AccountLockDate) error { return config.DB.Create(data).Error }
func GetAllAccountLockDate() ([]AccountLockDate, error) {
	var list []AccountLockDate
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetAccountLockDateByID(id uint) (*AccountLockDate, error) {
	var data AccountLockDate
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateAccountLockDate(data *AccountLockDate) error { return config.DB.Save(data).Error }
func DeleteAccountLockDate(id uint) error               { return config.DB.Delete(&AccountLockDate{}, id).Error }

func CreatePaymentAcquirer(data *PaymentAcquirer) error { return config.DB.Create(data).Error }
func GetAllPaymentAcquirer() ([]PaymentAcquirer, error) {
	var list []PaymentAcquirer
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetPaymentAcquirerByID(id uint) (*PaymentAcquirer, error) {
	var data PaymentAcquirer
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePaymentAcquirer(data *PaymentAcquirer) error { return config.DB.Save(data).Error }
func DeletePaymentAcquirer(id uint) error               { return config.DB.Delete(&PaymentAcquirer{}, id).Error }

func CreatePaymentTransaction(data *PaymentTransaction) error { return config.DB.Create(data).Error }
func GetAllPaymentTransaction() ([]PaymentTransaction, error) {
	var list []PaymentTransaction
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetPaymentTransactionByID(id uint) (*PaymentTransaction, error) {
	var data PaymentTransaction
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePaymentTransaction(data *PaymentTransaction) error { return config.DB.Save(data).Error }
func DeletePaymentTransaction(id uint) error {
	return config.DB.Delete(&PaymentTransaction{}, id).Error
}

func CreateAccountIncoterms(data *AccountIncoterms) error { return config.DB.Create(data).Error }
func GetAllAccountIncoterms() ([]AccountIncoterms, error) {
	var list []AccountIncoterms
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetAccountIncotermsByID(id uint) (*AccountIncoterms, error) {
	var data AccountIncoterms
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateAccountIncoterms(data *AccountIncoterms) error { return config.DB.Save(data).Error }
func DeleteAccountIncoterms(id uint) error                { return config.DB.Delete(&AccountIncoterms{}, id).Error }
