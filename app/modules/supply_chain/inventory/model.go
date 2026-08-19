package inventory

import (
	"ERP-System/config"
	"time"
)

type ProductCategory struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	ParentID      *uint  `json:"parent_id"`
	CostingMethod string `gorm:"type:varchar(50);default:'standard'" json:"costing_method"` // standard, average, fifo
}

type UoMCategory struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"` // Weight, Length, Unit
}

type UoM struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	CategoryID uint    `json:"category_id"`
	Name       string  `gorm:"type:varchar(255);not null" json:"name"` // kg, cm, pcs
	Type       string  `gorm:"type:varchar(50);default:'reference'" json:"type"` // reference, bigger, smaller
	Factor     float64 `gorm:"type:numeric(15,4);default:1" json:"factor"` // Ratio to reference UoM
}

type ProductTemplate struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Type          string    `gorm:"type:varchar(50);default:'product'" json:"type"` // product (storable), consu (consumable), service
	CategoryID    uint      `json:"category_id"`
	UoMID         uint      `json:"uom_id"`
	UoMPoID       uint      `json:"uom_po_id"` // UoM for Purchase
	ListPrice     float64   `gorm:"type:numeric(15,2);default:0" json:"list_price"` // Default Sale Price
	StandardPrice float64   `gorm:"type:numeric(15,2);default:0" json:"standard_price"` // Cost
	Tracking      string    `gorm:"type:varchar(50);default:'none'" json:"tracking"` // none, lot, serial
	CreatedAt     time.Time `json:"created_at"`
}

type ProductAttribute struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Color, Size
}

type ProductAttributeValue struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	AttributeID uint   `json:"attribute_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Red, XL
}

type Product struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ProductTemplateID uint      `json:"product_template_id"`
	DefaultCode       string    `gorm:"type:varchar(100);unique" json:"default_code"` // SKU
	Barcode           string    `gorm:"type:varchar(100)" json:"barcode"`
	Volume            float64   `gorm:"type:numeric(15,3);default:0" json:"volume"`
	Weight            float64   `gorm:"type:numeric(15,3);default:0" json:"weight"`
	StockQty          float64   `gorm:"type:numeric(15,2);default:0" json:"stock_qty"`
	CreatedAt         time.Time `json:"created_at"`
}

type StockWarehouse struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // Main Warehouse
	Code      string    `gorm:"type:varchar(10);not null" json:"code"`  // WH
	CompanyID uint      `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
}

type StockLocation struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"` // Stock, Shelf 1
	Barcode     string `gorm:"type:varchar(100);unique" json:"barcode"`
	ParentID    *uint  `json:"parent_id"`
	Usage       string `gorm:"type:varchar(50);default:'internal'" json:"usage"` // supplier, view, internal, customer, inventory, production
	WarehouseID *uint  `json:"warehouse_id"`
}

type StockPicking struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"` // WH/OUT/0001
	LocationID     uint      `json:"location_id"`                            // Source
	LocationDestID uint      `json:"location_dest_id"`                       // Destination
	PartnerID      *uint     `json:"partner_id"`                             // Customer/Vendor
	State          string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, waiting, confirmed, assigned, done, cancel
	ScheduledDate  time.Time `json:"scheduled_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type StockMove struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	PickingID      *uint     `json:"picking_id"`
	ProductID      uint      `json:"product_id"`
	Quantity       float64   `gorm:"type:numeric(15,2);not null;default:0" json:"quantity"`
	QuantityDone   float64   `gorm:"type:numeric(15,2);default:0" json:"quantity_done"`
	LocationID     uint      `json:"location_id"`
	LocationDestID uint      `json:"location_dest_id"`
	State          string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, waiting, confirmed, assigned, done, cancel
	CreatedAt      time.Time `json:"created_at"`
}

type StockLot struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"` // Lot/Serial Number
	ProductID      uint      `json:"product_id"`
	CompanyID      uint      `json:"company_id"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type StockQuant struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProductID  uint    `json:"product_id"`
	LocationID uint    `json:"location_id"`
	LotID      *uint   `json:"lot_id"`
	Quantity   float64 `gorm:"type:numeric(15,2);default:0" json:"quantity"` // Real-time available
}

type StockPutawayRule struct {
	ID             uint `gorm:"primaryKey" json:"id"`
	ProductID      uint `json:"product_id"`
	CategoryID     uint `json:"category_id"`
	LocationInID   uint `json:"location_in_id"`   // When arriving here...
	LocationOutID  uint `json:"location_out_id"`  // Put it here automatically
}

type StockValuationLayer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ProductID   uint      `json:"product_id"`
	Quantity    float64   `gorm:"type:numeric(15,2);default:0" json:"quantity"`
	UnitCost    float64   `gorm:"type:numeric(15,2);default:0" json:"unit_cost"`
	Value       float64   `gorm:"type:numeric(15,2);default:0" json:"value"` // Qty * UnitCost
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type StockInventory struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"` // INV/2026/001
	LocationID  uint      `json:"location_id"`
	Date        time.Time `json:"date"`
	State       string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, confirm, done
	CompanyID   uint      `json:"company_id"`
}

type StockInventoryLine struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	InventoryID    uint    `json:"inventory_id"`
	ProductID      uint    `json:"product_id"`
	LocationID     uint    `json:"location_id"`
	TheoreticalQty float64 `gorm:"type:numeric(15,2);default:0" json:"theoretical_qty"`
	ProductQty     float64 `gorm:"type:numeric(15,2);default:0" json:"product_qty"` // Actual counted
}

type StockScrap struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"` // SP/0001
	ProductID   uint      `json:"product_id"`
	ScrapQty    float64   `gorm:"type:numeric(15,2);not null;default:1" json:"scrap_qty"`
	LocationID  uint      `json:"location_id"`       // Dari mana asalnya
	ScrapLocID  uint      `json:"scrap_location_id"` // Lokasi pembuangan (Virtual)
	State       string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done
	CreatedAt   time.Time `json:"created_at"`
}

// ---- FITUR ODOO ENTERPRISE/DEWA (Ditambahkan Manual) ----

type StockRoute struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // Dropship, Make to Order (MTO)
	CompanyID uint      `json:"company_id"`
	Sequence  int       `gorm:"default:10" json:"sequence"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type StockRule struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	RouteID        uint      `json:"route_id"`
	Action         string    `gorm:"type:varchar(50);not null" json:"action"` // pull, push, buy, manufacture
	LocationSrcID  *uint     `json:"location_src_id"`
	LocationDestID *uint     `json:"location_dest_id"`
	Sequence       int       `gorm:"default:10" json:"sequence"`
	CreatedAt      time.Time `json:"created_at"`
}

type StockLandedCost struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100)" json:"name"` // LC/2026/001
	Date        time.Time `json:"date"`
	PickingID   uint      `json:"picking_id"` // Dokumen penerimaan barang dari vendor
	AmountTotal float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	State       string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done
	CreatedAt   time.Time `json:"created_at"`
}

type StockLandedCostLine struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	CostID      uint    `json:"cost_id"`
	ProductID   uint    `json:"product_id"` // Biaya (misal: Cukai, Ongkir Kapal)
	SplitMethod string  `gorm:"type:varchar(50);default:'equal'" json:"split_method"` // equal, by_quantity, by_weight, by_volume
	PriceUnit   float64 `gorm:"type:numeric(15,2);not null;default:0" json:"price_unit"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ProductCategory{}, &UoMCategory{}, &UoM{}, &ProductTemplate{}, &ProductAttribute{}, &ProductAttributeValue{}, &Product{}, &StockWarehouse{}, &StockLocation{}, &StockPicking{}, &StockMove{}, &StockLot{}, &StockQuant{}, &StockPutawayRule{}, &StockValuationLayer{}, &StockInventory{}, &StockInventoryLine{}, &StockScrap{}, &StockRoute{}, &StockRule{}, &StockLandedCost{}, &StockLandedCostLine{})
}
