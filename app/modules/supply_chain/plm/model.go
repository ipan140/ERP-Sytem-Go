package plm

import (
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/app/modules/supply_chain/inventory"
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
	Type *PlmEcoType `gorm:"foreignKey:TypeID" json:"type,omitempty"` // Odoo relation mapped
	ProductID uint      `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	OldBomID  *uint     `json:"old_bom_id"`                                    // Resep lama
	OldBom *manufacturing.MrpBom `gorm:"foreignKey:OldBomID" json:"oldbom,omitempty"` // Odoo relation mapped
	NewBomID  *uint     `json:"new_bom_id"`                                    // Resep revisi
	NewBom *manufacturing.MrpBom `gorm:"foreignKey:NewBomID" json:"newbom,omitempty"` // Odoo relation mapped
	State     string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, progress, approved, done
	CreatedAt time.Time `json:"created_at"`
}


func (PlmEcoType) TableName() string {
	return "supply_chain.plm_eco_types"
}

func (PlmEco) TableName() string {
	return "supply_chain.plm_ecos"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PlmEcoType{}, &PlmEco{})
}
