package base

import (
	"ERP-System/config"
	"time"
)

type Currency struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(10);default:''" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. USD, IDR
	Symbol    string    `gorm:"type:varchar(10)" json:"symbol"`
	Rate      float64   `gorm:"type:decimal(15,4);default:1.0" json:"rate"`
	IsBase    bool      `gorm:"default:false" json:"is_base"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Country struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Code string `gorm:"type:varchar(10);unique" json:"code"` // ID, US, SG
}

type CountryState struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CountryID uint   `json:"country_id"`
	Country *Country `gorm:"foreignKey:CountryID"` // Auto-added relation
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Code      string `gorm:"type:varchar(10)" json:"code"` // JBR, JK
}

type Partner struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	IsCompany  bool      `gorm:"default:false" json:"is_company"`
	ParentID   *uint     `json:"parent_id"`                                      // Jika kontak ini adalah pegawai dari perusahaan lain
	Parent *Partner `gorm:"foreignKey:ParentID"` // Auto-added relation
	Type       string    `gorm:"type:varchar(50);default:'contact'" json:"type"` // contact, invoice, delivery
	Email      string    `gorm:"type:varchar(255)" json:"email"`
	Phone      string    `gorm:"type:varchar(50)" json:"phone"`
	Mobile     string    `gorm:"type:varchar(50)" json:"mobile"`
	Street     string    `gorm:"type:varchar(255)" json:"street"`
	Street2    string    `gorm:"type:varchar(255)" json:"street2"`
	City       string    `gorm:"type:varchar(100)" json:"city"`
	StateID    *uint     `json:"state_id"`
	State *CountryState `gorm:"foreignKey:StateID" json:"state,omitempty"` // Odoo relation mapped
	Zip        string    `gorm:"type:varchar(20)" json:"zip"`
	CountryID  *uint     `json:"country_id"`
	Country *Country `gorm:"foreignKey:CountryID"` // Auto-added relation
	Vat        string    `gorm:"type:varchar(50)" json:"vat"` // NPWP / Tax ID
	IsCustomer bool      `gorm:"default:true" json:"is_customer"`
	IsVendor   bool      `gorm:"default:false" json:"is_vendor"`
	
	// Enterprise Credit Limit & Financial Risk Control
	CreditLimit     float64   `gorm:"type:numeric(15,2);default:50000000" json:"credit_limit"`     // Plafon limit kredit piutang (Default: Rp 50 Juta)
	IsCreditHold    bool      `gorm:"default:false" json:"is_credit_hold"`                        // Apakah akun pelanggan sedang diblokir/on-hold
	MaxOverdueDays  int       `gorm:"default:30" json:"max_overdue_days"`                         // Toleransi jatuh tempo maksimal (hari)
	TotalReceivable float64   `gorm:"type:numeric(15,2);default:0" json:"total_receivable"`       // Saldo piutang aktif

	CreatedAt  time.Time `json:"created_at"`
}


func (Currency) TableName() string {
	return "setting.currencies"
}

func (Country) TableName() string {
	return "setting.countries"
}

func (CountryState) TableName() string {
	return "setting.country_states"
}

func (Partner) TableName() string {
	return "setting.partners"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Currency{}, &Country{}, &CountryState{}, &Partner{})
}
