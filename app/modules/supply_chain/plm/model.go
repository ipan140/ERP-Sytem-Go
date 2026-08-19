package plm

import (
	"ERP-System/config"
	"time"
)

type PlmEcoType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"` // PlmEco Change, Routing Change
}

type PlmEco struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // ECO/2026/001: Ubah Kayu Jadi Plastik
	TypeID    uint      `json:"type_id"`
	ProductID uint      `json:"product_id"`
	OldBomID  *uint     `json:"old_bom_id"` // Resep lama
	NewBomID  *uint     `json:"new_bom_id"` // Resep revisi
	State     string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, progress, approved, done
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PlmEcoType{}, &PlmEco{})
}
