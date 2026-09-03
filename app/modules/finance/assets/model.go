package assets

import (
	"ERP-System/config"
	"time"
)

type AssetCategory struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Code              string    `gorm:"type:varchar(50);unique;not null" json:"code"` // e.g. IT, VEHICLE, MACHINE
	Name              string    `gorm:"type:varchar(100);not null" json:"name"`       // e.g. Peralatan IT & Komputer
	DefaultUsefulLife int       `gorm:"default:48" json:"default_useful_life"`        // Bulan (misal 48 bulan = 4 tahun)
	GLAccountDebit    string    `gorm:"type:varchar(100);default:'6-1008 Beban Depresiasi'" json:"gl_account_debit"`
	GLAccountCredit   string    `gorm:"type:varchar(100);default:'1-1009 Akumulasi Penyusutan'" json:"gl_account_credit"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type FixedAsset struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Code                string    `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name                string    `gorm:"type:varchar(200);not null" json:"name"`
	Category            string    `gorm:"type:varchar(100)" json:"category"` // Peralatan Kantor, Kendaraan, Mesin
	AcquisitionDate     time.Time `json:"acquisition_date"`
	AcquisitionCost     float64   `gorm:"type:numeric(15,2)" json:"acquisition_cost"`
	UsefulLifeMonths    int       `json:"useful_life_months"` // 48 bulan = 4 tahun
	ResidualValue       float64   `gorm:"type:numeric(15,2);default:0" json:"residual_value"`
	DepreciationMethod  string    `gorm:"type:varchar(50);default:'Garis Lurus'" json:"depreciation_method"`
	MonthlyDepreciation float64   `gorm:"type:numeric(15,2)" json:"monthly_depreciation"`
	AccumulatedDepr     float64   `gorm:"type:numeric(15,2)" json:"accumulated_depreciation"`
	NetBookValue        float64   `gorm:"type:numeric(15,2)" json:"net_book_value"`
	Status              string    `gorm:"type:varchar(30);default:'Aktif'" json:"status"`
	CreatedAt           time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &FixedAsset{}, &AssetCategory{})
}

func (FixedAsset) TableName() string {
	return "finance.fixed_assets"
}

func (AssetCategory) TableName() string {
	return "finance.asset_categories"
}
