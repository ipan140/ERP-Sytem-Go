package inventory

import (
	"ERP-System/app/modules/auth"
	"ERP-System/app/modules/core/base"
	"ERP-System/config"
	"time"
)

type ProductCategory struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	ParentID      *uint  `json:"parent_id"`
	Parent *ProductCategory `gorm:"foreignKey:ParentID"` // Auto-added relation
	CostingMethod string `gorm:"type:varchar(50);default:'standard'" json:"costing_method"` // standard, average, fifo
}

type UoMCategory struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"` // Weight, Length, Unit
}

type UoM struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	CategoryID uint    `json:"category_id"`
	Category *ProductCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"` // Odoo relation mapped
	Name       string  `gorm:"type:varchar(255);not null" json:"name"`           // kg, cm, pcs
	Type       string  `gorm:"type:varchar(50);default:'reference'" json:"type"` // reference, bigger, smaller
	Factor     float64 `gorm:"type:numeric(15,4);default:1" json:"factor"`       // Ratio to reference UoM
}

type ProductTemplate struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Type          string    `gorm:"type:varchar(50);default:'product'" json:"type"` // product (storable), consu (consumable), service
	CategoryID    uint      `json:"category_id"`
	Category *ProductCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"` // Odoo relation mapped
	UoMID         uint      `json:"uom_id"`
	UoM *UoM `gorm:"foreignKey:UoMID"` // Auto-added relation
	UoMPoID       uint      `json:"uom_po_id"`                                          // UoM for Purchase
	UoMPo *UoM `gorm:"foreignKey:UoMPoID" json:"uompo,omitempty"` // Odoo relation mapped
	ListPrice     float64   `gorm:"type:numeric(15,2);default:0" json:"list_price"`     // Default Sale Price
	StandardPrice float64   `gorm:"type:numeric(15,2);default:0" json:"standard_price"` // Cost
	Tracking      string    `gorm:"type:varchar(50);default:'none'" json:"tracking"`    // none, lot, serial
	CreatedAt     time.Time `json:"created_at"`
}

type ProductAttribute struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Color, Size
}

type ProductAttributeValue struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	AttributeID uint   `json:"attribute_id"`
	Attribute *ProductAttribute `gorm:"foreignKey:AttributeID" json:"attribute,omitempty"` // Odoo relation mapped
	Name        string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Red, XL
}

type Product struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ProductTemplateID uint      `json:"product_template_id"`
	ProductTemplate *ProductTemplate `gorm:"foreignKey:ProductTemplateID"` // Auto-added relation
	DefaultCode       string    `gorm:"type:varchar(100);unique" json:"default_code"` // SKU
	Barcode           string    `gorm:"type:varchar(100)" json:"barcode"`
	Volume            float64   `gorm:"type:numeric(15,3);default:0" json:"volume"`
	Weight            float64   `gorm:"type:numeric(15,3);default:0" json:"weight"`
	StockQty          float64   `gorm:"type:numeric(15,2);default:0" json:"stock_qty"`
	ReservedQty       float64   `gorm:"type:numeric(15,2);default:0" json:"reserved_qty"` // FASE 3: Stok terpesan (dikunci oleh Sales Order terkonfirmasi)
	CreatedAt         time.Time `json:"created_at"`
}

type StockWarehouse struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // Main Warehouse
	Code      string    `gorm:"type:varchar(10);not null" json:"code"`  // WH
	CompanyID uint      `json:"company_id"`
	Company *auth.Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"` // Cross-module relation
	CreatedAt time.Time `json:"created_at"`
}

type StockLocation struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"` // Stock, Shelf 1
	Barcode     string `gorm:"type:varchar(100);unique" json:"barcode"`
	ParentID    *uint  `json:"parent_id"`
	Parent *StockLocation `gorm:"foreignKey:ParentID"` // Auto-added relation
	Usage       string `gorm:"type:varchar(50);default:'internal'" json:"usage"` // supplier, view, internal, customer, inventory, production
	WarehouseID *uint  `json:"warehouse_id"`
	Warehouse *StockWarehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"` // Odoo relation mapped
}

type StockPicking struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"`        // WH/OUT/0001
	LocationID     uint      `json:"location_id"`                                   // Source
	Location *StockLocation `gorm:"foreignKey:LocationID" json:"location,omitempty"` // Odoo relation mapped
	LocationDestID uint      `json:"location_dest_id"`                              // Destination
	LocationDest *StockLocation `gorm:"foreignKey:LocationDestID" json:"locationdest,omitempty"` // Odoo relation mapped
	PartnerID      *uint     `json:"partner_id"`                                    // Customer/Vendor
	Partner *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"` // Cross-module relation
	State          string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, waiting, confirmed, assigned, done, cancel
	ScheduledDate  time.Time `json:"scheduled_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type StockMove struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	PickingID      *uint     `json:"picking_id"`
	Picking *StockPicking `gorm:"foreignKey:PickingID" json:"picking,omitempty"` // Odoo relation mapped
	ProductID      uint      `json:"product_id"`
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	Quantity       float64   `gorm:"type:numeric(15,2);not null;default:0" json:"quantity"`
	QuantityDone   float64   `gorm:"type:numeric(15,2);default:0" json:"quantity_done"`
	LocationID     uint      `json:"location_id"`
	Location *StockLocation `gorm:"foreignKey:LocationID" json:"location,omitempty"` // Odoo relation mapped
	LocationDestID uint      `json:"location_dest_id"`
	LocationDest *StockLocation `gorm:"foreignKey:LocationDestID" json:"locationdest,omitempty"` // Odoo relation mapped
	State          string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, waiting, confirmed, assigned, done, cancel
	CreatedAt      time.Time `json:"created_at"`
}

type StockLot struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"` // Lot/Serial Number
	ProductID      uint      `json:"product_id"`
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	CompanyID      uint      `json:"company_id"`
	Company *auth.Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"` // Cross-module relation
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type StockQuant struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProductID  uint    `json:"product_id"`
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	LocationID uint    `json:"location_id"`
	Location *StockLocation `gorm:"foreignKey:LocationID" json:"location,omitempty"` // Odoo relation mapped
	LotID      *uint   `json:"lot_id"`
	Lot *StockLot `gorm:"foreignKey:LotID" json:"lot,omitempty"` // Odoo relation mapped
	Quantity   float64 `gorm:"type:numeric(15,2);default:0" json:"quantity"` // Real-time available
}

type StockPutawayRule struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	ProductID     uint `json:"product_id"`
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	CategoryID    uint `json:"category_id"`
	Category *ProductCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"` // Odoo relation mapped
	LocationInID  uint `json:"location_in_id"`  // When arriving here...
	LocationIn *StockLocation `gorm:"foreignKey:LocationInID" json:"locationin,omitempty"` // Odoo relation mapped
	LocationOutID uint `json:"location_out_id"` // Put it here automatically
	LocationOut *StockLocation `gorm:"foreignKey:LocationOutID" json:"locationout,omitempty"` // Odoo relation mapped
}

type StockValuationLayer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ProductID   uint      `json:"product_id"`
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	Quantity    float64   `gorm:"type:numeric(15,2);default:0" json:"quantity"`
	UnitCost    float64   `gorm:"type:numeric(15,2);default:0" json:"unit_cost"`
	Value       float64   `gorm:"type:numeric(15,2);default:0" json:"value"` // Qty * UnitCost
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type StockInventory struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"` // INV/2026/001
	LocationID uint      `json:"location_id"`
	Location *StockLocation `gorm:"foreignKey:LocationID" json:"location,omitempty"` // Odoo relation mapped
	Date       time.Time `json:"date"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, confirm, done
	CompanyID  uint      `json:"company_id"`
	Company *auth.Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"` // Cross-module relation
}

type StockInventoryLine struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	InventoryID    uint    `json:"inventory_id"`
	Inventory *StockInventory `gorm:"foreignKey:InventoryID" json:"inventory,omitempty"` // Odoo relation mapped
	ProductID      uint    `json:"product_id"`
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	LocationID     uint    `json:"location_id"`
	Location *StockLocation `gorm:"foreignKey:LocationID" json:"location,omitempty"` // Odoo relation mapped
	TheoreticalQty float64 `gorm:"type:numeric(15,2);default:0" json:"theoretical_qty"`
	ProductQty     float64 `gorm:"type:numeric(15,2);default:0" json:"product_qty"` // Actual counted
}

type StockScrap struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"name"` // SP/0001
	ProductID  uint      `json:"product_id"`
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	ScrapQty   float64   `gorm:"type:numeric(15,2);not null;default:1" json:"scrap_qty"`
	LocationID uint      `json:"location_id"`                                   // Dari mana asalnya
	Location *StockLocation `gorm:"foreignKey:LocationID" json:"location,omitempty"` // Odoo relation mapped
	ScrapLocID uint      `json:"scrap_location_id"`                             // Lokasi pembuangan (Virtual)
	ScrapLoc *StockLocation `gorm:"foreignKey:ScrapLocID" json:"scraploc,omitempty"` // Odoo relation mapped
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done
	CreatedAt  time.Time `json:"created_at"`
}

// ---- FITUR ODOO ENTERPRISE/DEWA (Ditambahkan Manual) ----

type StockRoute struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // Dropship, Make to Order (MTO)
	CompanyID uint      `json:"company_id"`
	Company *auth.Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"` // Cross-module relation
	Sequence  int       `gorm:"default:10" json:"sequence"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type StockRule struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	RouteID        uint      `json:"route_id"`
	Route *StockRoute `gorm:"foreignKey:RouteID" json:"route,omitempty"` // Odoo relation mapped
	Action         string    `gorm:"type:varchar(50);not null" json:"action"` // pull, push, buy, manufacture
	LocationSrcID  *uint     `json:"location_src_id"`
	LocationSrc *StockLocation `gorm:"foreignKey:LocationSrcID" json:"locationsrc,omitempty"` // Odoo relation mapped
	LocationDestID *uint     `json:"location_dest_id"`
	LocationDest *StockLocation `gorm:"foreignKey:LocationDestID" json:"locationdest,omitempty"` // Odoo relation mapped
	Sequence       int       `gorm:"default:10" json:"sequence"`
	CreatedAt      time.Time `json:"created_at"`
}

type StockLandedCost struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100)" json:"name"` // LC/2026/001
	Date        time.Time `json:"date"`
	PickingID   uint      `json:"picking_id"` // Dokumen penerimaan barang dari vendor
	Picking *StockPicking `gorm:"foreignKey:PickingID" json:"picking,omitempty"` // Odoo relation mapped
	AmountTotal float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	State       string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done
	CreatedAt   time.Time `json:"created_at"`
}

type StockLandedCostLine struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	CostID      uint    `json:"cost_id"`
	Cost *StockLandedCost `gorm:"foreignKey:CostID" json:"cost,omitempty"` // Odoo relation mapped
	ProductID   uint    `json:"product_id"`                                           // Biaya (misal: Cukai, Ongkir Kapal)
	Product *Product `gorm:"foreignKey:ProductID"` // Auto-added relation
	SplitMethod string  `gorm:"type:varchar(50);default:'equal'" json:"split_method"` // equal, by_quantity, by_weight, by_volume
	PriceUnit   float64 `gorm:"type:numeric(15,2);not null;default:0" json:"price_unit"`
}


func (ProductCategory) TableName() string {
	return "supply_chain.product_categories"
}

func (UoMCategory) TableName() string {
	return "supply_chain.uo_m_categories"
}

func (UoM) TableName() string {
	return "supply_chain.uo_ms"
}

func (ProductTemplate) TableName() string {
	return "supply_chain.product_templates"
}

func (ProductAttribute) TableName() string {
	return "supply_chain.product_attributes"
}

func (ProductAttributeValue) TableName() string {
	return "supply_chain.product_attribute_values"
}

func (Product) TableName() string {
	return "supply_chain.products"
}

func (StockWarehouse) TableName() string {
	return "supply_chain.stock_warehouses"
}

func (StockLocation) TableName() string {
	return "supply_chain.stock_locations"
}

func (StockPicking) TableName() string {
	return "supply_chain.stock_pickings"
}

func (StockMove) TableName() string {
	return "supply_chain.stock_moves"
}

func (StockLot) TableName() string {
	return "supply_chain.stock_lots"
}

func (StockQuant) TableName() string {
	return "supply_chain.stock_quants"
}

func (StockPutawayRule) TableName() string {
	return "supply_chain.stock_putaway_rules"
}

func (StockValuationLayer) TableName() string {
	return "supply_chain.stock_valuation_layers"
}

func (StockInventory) TableName() string {
	return "supply_chain.stock_inventories"
}

func (StockInventoryLine) TableName() string {
	return "supply_chain.stock_inventory_lines"
}

func (StockScrap) TableName() string {
	return "supply_chain.stock_scraps"
}

func (StockRoute) TableName() string {
	return "supply_chain.stock_routes"
}

func (StockRule) TableName() string {
	return "supply_chain.stock_rules"
}

func (StockLandedCost) TableName() string {
	return "supply_chain.stock_landed_costs"
}

func (StockLandedCostLine) TableName() string {
	return "supply_chain.stock_landed_cost_lines"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ProductCategory{}, &UoMCategory{}, &UoM{}, &ProductTemplate{}, &ProductAttribute{}, &ProductAttributeValue{}, &Product{}, &StockWarehouse{}, &StockLocation{}, &StockPicking{}, &StockMove{}, &StockLot{}, &StockQuant{}, &StockPutawayRule{}, &StockValuationLayer{}, &StockInventory{}, &StockInventoryLine{}, &StockScrap{}, &StockRoute{}, &StockRule{}, &StockLandedCost{}, &StockLandedCostLine{})
}
