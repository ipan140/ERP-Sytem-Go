package plm

import (
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/config"
	"time"
)

type PlmEcoType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"` // ECO BoM Change, Routing Change, Material Substitution
}

type PlmEco struct {
	ID            uint                        `gorm:"primaryKey" json:"id"`
	Code          string                      `gorm:"type:varchar(50);uniqueIndex" json:"code"` // ECO/2026/0001
	Name          string                      `gorm:"type:varchar(255);not null" json:"name"` // Judul Perintah Perubahan
	TypeID        uint                        `json:"type_id"`
	Type          *PlmEcoType                 `gorm:"foreignKey:TypeID" json:"type,omitempty"`
	ProductID     uint                        `json:"product_id"`
	Product       *inventory.Product          `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	OldBomID      *uint                       `json:"old_bom_id"` // Resep formula lama
	OldBom        *manufacturing.MrpBom       `gorm:"foreignKey:OldBomID" json:"old_bom,omitempty"`
	NewBomID      *uint                       `json:"new_bom_id"` // Resep formula revisi baru
	NewBom        *manufacturing.MrpBom       `gorm:"foreignKey:NewBomID" json:"new_bom,omitempty"`
	Reason        string                      `gorm:"type:text" json:"reason"` // Alasan teknis / spesifikasi
	EffectiveDate *time.Time                  `json:"effective_date"` // Tanggal efektif berlaku
	State         string                      `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, progress, approved, done, cancel
	ApproverID    *uint                       `json:"approver_id"`
	ApprovedAt    *time.Time                  `json:"approved_at"`
	CreatedAt     time.Time                   `json:"created_at"`
}

func (PlmEcoType) TableName() string {
	return "supply_chain.plm_eco_types"
}

func (PlmEco) TableName() string {
	return "supply_chain.plm_ecos"
}

type PlmSummary struct {
	TotalEco      int64 `json:"total_eco"`
	DraftCount    int64 `json:"draft_count"`
	ProgressCount int64 `json:"progress_count"`
	ApprovedCount int64 `json:"approved_count"`
	DoneCount     int64 `json:"done_count"`
}

type CreateEcoRequest struct {
	Name          string  `json:"name"`
	TypeID        uint    `json:"type_id"`
	ProductID     uint    `json:"product_id"`
	OldBomID      *uint   `json:"old_bom_id"`
	NewBomID      *uint   `json:"new_bom_id"`
	Reason        string  `json:"reason"`
	EffectiveDate *string `json:"effective_date"`
}

type UpdateEcoStateRequest struct {
	State string  `json:"state"` // progress, approved, done, cancel
	Notes *string `json:"notes"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PlmEcoType{}, &PlmEco{})
}
