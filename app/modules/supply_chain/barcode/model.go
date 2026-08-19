package barcode

import (
	"ERP-System/config"
	"time"
)

type BarcodeNomenclature struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Default Nomenclature, GS1 Nomenclature
	UpcEanConv string    `gorm:"type:varchar(50);default:'always'" json:"upc_ean_conv"` // always, never, ean13_to_upca
	CreatedAt  time.Time `json:"created_at"`
}

type BarcodeRule struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Weight Barcode
	NomenclatureID uint      `json:"nomenclature_id"`
	Sequence       int       `gorm:"default:10" json:"sequence"`
	Type           string    `gorm:"type:varchar(50);not null" json:"type"` // product, lot, location, package, weight, discount, client
	Encoding       string    `gorm:"type:varchar(50);default:'any'" json:"encoding"` // any, ean13, ean8, upca, gs1-128
	Pattern        string    `gorm:"type:varchar(255);not null" json:"pattern"` // Regex pattern (e.g. 21.....{NNDDD})
	CreatedAt      time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &BarcodeNomenclature{}, &BarcodeRule{})
}
