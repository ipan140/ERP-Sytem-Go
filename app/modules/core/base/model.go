package base

import (
	"ERP-System/config"
	"time"
)

type Currency struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. USD, IDR
	Symbol    string    `gorm:"type:varchar(10)" json:"symbol"`
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
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Code      string `gorm:"type:varchar(10)" json:"code"` // JBR, JK
}

type Partner struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	IsCompany    bool      `gorm:"default:false" json:"is_company"`
	ParentID     *uint     `json:"parent_id"` // Jika kontak ini adalah pegawai dari perusahaan lain
	Type         string    `gorm:"type:varchar(50);default:'contact'" json:"type"` // contact, invoice, delivery
	Email        string    `gorm:"type:varchar(255)" json:"email"`
	Phone        string    `gorm:"type:varchar(50)" json:"phone"`
	Mobile       string    `gorm:"type:varchar(50)" json:"mobile"`
	Street       string    `gorm:"type:varchar(255)" json:"street"`
	Street2      string    `gorm:"type:varchar(255)" json:"street2"`
	City         string    `gorm:"type:varchar(100)" json:"city"`
	StateID      *uint     `json:"state_id"`
	Zip          string    `gorm:"type:varchar(20)" json:"zip"`
	CountryID    *uint     `json:"country_id"`
	Vat          string    `gorm:"type:varchar(50)" json:"vat"` // NPWP / Tax ID
	IsCustomer   bool      `gorm:"default:true" json:"is_customer"`
	IsVendor     bool      `gorm:"default:false" json:"is_vendor"`
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Currency{}, &Country{}, &CountryState{}, &Partner{})
}
