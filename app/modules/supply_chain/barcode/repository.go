package barcode

import (
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/app/modules/supply_chain/quality"
	"ERP-System/config"
	"strings"

	"gorm.io/gorm/clause"
)

func ScanBarcode(rawCode string) (*BarcodeScanResponse, error) {
	code := strings.TrimSpace(rawCode)
	if code == "" {
		return &BarcodeScanResponse{
			Found:   false,
			Barcode: code,
			Type:    "unknown",
			Title:   "Kode Barcode Kosong",
		}, nil
	}

	// 1. Search in Products (by barcode or default_code / SKU)
	var prod inventory.Product
	err := config.DB.Preload("ProductTemplate").
		Preload("ProductTemplate.Category").
		Preload("ProductTemplate.UoM").
		Where("barcode = ? OR default_code = ? OR id::text = ?", code, code, code).
		First(&prod).Error
	if err == nil && prod.ID > 0 {
		title := prod.DefaultCode
		if prod.ProductTemplate != nil && prod.ProductTemplate.Name != "" {
			title = prod.ProductTemplate.Name
		}
		return &BarcodeScanResponse{
			Type:    "product",
			Found:   true,
			Barcode: code,
			Title:   title,
			Data:    prod,
		}, nil
	}

	// 2. Search in Stock Pickings / Penerimaan Gudang
	var picking inventory.StockPicking
	err = config.DB.Preload(clause.Associations).
		Where("name = ? OR id::text = ?", code, code).
		First(&picking).Error
	if err == nil && picking.ID > 0 {
		return &BarcodeScanResponse{
			Type:    "picking",
			Found:   true,
			Barcode: code,
			Title:   "Dokumen Transfer / Picking " + picking.Name,
			Data:    picking,
		}, nil
	}

	// 3. Search in Manufacturing Orders (SPK MO)
	var mo manufacturing.MrpProduction
	err = config.DB.Preload(clause.Associations).
		Where("name = ? OR id::text = ?", code, code).
		First(&mo).Error
	if err == nil && mo.ID > 0 {
		return &BarcodeScanResponse{
			Type:    "mo",
			Found:   true,
			Barcode: code,
			Title:   "Perintah Produksi " + mo.Name,
			Data:    mo,
		}, nil
	}

	// 4. Search in Quality Checks
	var qc quality.QualityCheck
	err = config.DB.Preload(clause.Associations).
		Where("name = ? OR id::text = ?", code, code).
		First(&qc).Error
	if err == nil && qc.ID > 0 {
		return &BarcodeScanResponse{
			Type:    "qc",
			Found:   true,
			Barcode: code,
			Title:   "Lembar Pemeriksaan Mutu " + qc.Name,
			Data:    qc,
		}, nil
	}

	// 5. Search in Locations
	var loc inventory.StockLocation
	err = config.DB.Where("name = ? OR barcode = ? OR id::text = ?", code, code, code).
		First(&loc).Error
	if err == nil && loc.ID > 0 {
		return &BarcodeScanResponse{
			Type:    "location",
			Found:   true,
			Barcode: code,
			Title:   "Lokasi Rak Gudang " + loc.Name,
			Data:    loc,
		}, nil
	}

	return &BarcodeScanResponse{
		Type:    "unknown",
		Found:   false,
		Barcode: code,
		Title:   "Objek Tidak Ditemukan",
		Data:    nil,
	}, nil
}

func GetBarcodeSummary() (*BarcodeSummary, error) {
	var summary BarcodeSummary
	config.DB.Model(&BarcodeNomenclature{}).Count(&summary.TotalNomenclatures)
	config.DB.Model(&BarcodeRule{}).Count(&summary.TotalRules)
	config.DB.Model(&inventory.Product{}).Where("barcode != '' OR default_code != ''").Count(&summary.TotalBarcodedSKUs)
	return &summary, nil
}

func GetAllBarcodeConfig() ([]BarcodeNomenclature, error) {
	var list []BarcodeNomenclature
	err := config.DB.Preload("Rules").Find(&list).Error
	if len(list) == 0 {
		// Seed standard GS1 nomenclature
		defaultNom := BarcodeNomenclature{
			Name:       "Default GS1 & EAN-13 Nomenclature",
			UpcEanConv: "always",
			Rules: []BarcodeRule{
				{Name: "Barcode Produk (EAN-13)", Sequence: 10, Type: "product", Encoding: "ean13", Pattern: "21.....{NNDDD}"},
				{Name: "Nomor Lot / Serial (GS1)", Sequence: 20, Type: "lot", Encoding: "gs1-128", Pattern: "10{NNNNNN}"},
				{Name: "Lokasi Rak Gudang", Sequence: 30, Type: "location", Encoding: "any", Pattern: "LOC-{NNNN}"},
				{Name: "Kemasan / Pallet (Package)", Sequence: 40, Type: "package", Encoding: "any", Pattern: "PKG-{NNNN}"},
			},
		}
		config.DB.Create(&defaultNom)
		config.DB.Preload("Rules").Find(&list)
	}
	return list, err
}

func GetBarcodeConfigByID(id uint) (*BarcodeNomenclature, error) {
	var data BarcodeNomenclature
	err := config.DB.Preload("Rules").First(&data, id).Error
	return &data, err
}

func CreateBarcodeConfig(data *BarcodeNomenclature) error {
	return config.DB.Create(data).Error
}

func UpdateBarcodeConfig(data *BarcodeNomenclature) error {
	return config.DB.Save(data).Error
}

func DeleteBarcodeConfig(id uint) error {
	return config.DB.Delete(&BarcodeNomenclature{}, id).Error
}
