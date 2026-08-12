package auth

import (
	"ERP-System/config"
	"time"
)

type Company struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      string    `gorm:"type:varchar(50);default:'staff'" json:"role"`
	CompanyID uint      `json:"company_id"`
	Company   Company   `gorm:"foreignKey:CompanyID" json:"company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// init otomatis dipanggil oleh Golang saat package ini di-import.
// Ini adalah trik modular Odoo agar modul meregistrasi modelnya sendiri.
func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Company{}, &User{})
}
