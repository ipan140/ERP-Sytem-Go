package barcode

import (
	"ERP-System/config"
	"time"
)

type BarcodeNomenclature struct {
	ID         uint          `gorm:"primaryKey" json:"id"`
	Name       string        `gorm:"type:varchar(255);not null" json:"name"`                // e.g. Default GS1 Nomenclature
	UpcEanConv string        `gorm:"type:varchar(50);default:'always'" json:"upc_ean_conv"` // always, never, ean13_to_upca
	Rules      []BarcodeRule `gorm:"foreignKey:NomenclatureID" json:"rules,omitempty"`
	CreatedAt  time.Time     `json:"created_at"`
}

type BarcodeRule struct {
	ID             uint                 `gorm:"primaryKey" json:"id"`
	Name           string               `gorm:"type:varchar(255);not null" json:"name"` // e.g. Product Barcodes, Weight Barcode, Location
	NomenclatureID uint                 `json:"nomenclature_id"`
	Nomenclature   *BarcodeNomenclature `gorm:"foreignKey:NomenclatureID" json:"nomenclature,omitempty"`
	Sequence       int                  `gorm:"default:10" json:"sequence"`
	Type           string               `gorm:"type:varchar(50);not null" json:"type"`          // product, lot, location, package, weight
	Encoding       string               `gorm:"type:varchar(50);default:'any'" json:"encoding"` // any, ean13, ean8, upca, gs1-128
	Pattern        string               `gorm:"type:varchar(255);not null" json:"pattern"`      // Regex pattern
	CreatedAt      time.Time            `json:"created_at"`
}

func (BarcodeNomenclature) TableName() string {
	return "supply_chain.barcode_nomenclatures"
}

func (BarcodeRule) TableName() string {
	return "supply_chain.barcode_rules"
}

// Enterprise DTOs
type BarcodeScanRequest struct {
	Barcode string `json:"barcode"`
}

type BarcodeScanResponse struct {
	Type    string      `json:"type"` // product, location, picking, mo, qc
	Found   bool        `json:"found"`
	Barcode string      `json:"barcode"`
	Title   string      `json:"title"`
	Data    interface{} `json:"data"`
}

type BarcodeSummary struct {
	TotalNomenclatures int64 `json:"total_nomenclatures"`
	TotalRules         int64 `json:"total_rules"`
	TotalBarcodedSKUs  int64 `json:"total_barcoded_skus"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &BarcodeNomenclature{}, &BarcodeRule{})
}
