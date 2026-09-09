package inventory

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

type InventoryDocumentPayload struct {
	Action string `json:"action"` // "import_master_data" or "print_delivery_order"
	RefID  string `json:"ref_id"` // File path for import, DO number for printing
	UserID uint   `json:"user_id"`
}

func ProcessInventoryDocumentService(action string, refID string, userID uint) error {
	if rabbitmq.Channel != nil {
		req := InventoryDocumentPayload{
			Action: action,
			RefID:  refID,
			UserID: userID,
		}
		body, _ := json.Marshal(req)
		
		var queue string
		if action == "import_master_data" {
			queue = "core_excel_import"
		} else {
			queue = "finance_report_generator" // Reusing document generator queue
		}

		err := rabbitmq.PublishEvent(rabbitmq.Channel, queue, body)
		log.Printf("📦 Event RabbitMQ: Inventory Document Task '%s' dikirim ke antrean %s!", action, queue)
		return err
	}
	return nil
}

func CreateProductService(data *Product) error {
	return CreateProduct(data)
}

func GetAllProductService() ([]Product, error) {
	return GetAllProduct()
}

func GetPaginatedProductsService(offset, limit int, search string, categoryID, warehouseID uint, stockStatus string) ([]Product, int64, error) {
	return GetPaginatedProducts(offset, limit, search, categoryID, warehouseID, stockStatus)
}

func GetInventorySummaryService() (InventorySummary, error) {
	return GetInventorySummary()
}

func ApplyStockAdjustmentService(req StockAdjustmentRequest) error {
	return ApplyStockAdjustment(req)
}

func ApplyInternalTransferService(req InternalTransferRequest) error {
	return ApplyInternalTransfer(req)
}

func GetProductByIDService(id uint) (*Product, error) {
	return GetProductByID(id)
}

func UpdateProductService(data *Product) error {
	return UpdateProduct(data)
}

func DeleteProductService(id uint) error {
	return DeleteProduct(id)
}

func CreateProductCategoryService(data *ProductCategory) error { return CreateProductCategory(data) }
func GetAllProductCategoryService() ([]ProductCategory, error) { return GetAllProductCategory() }
func GetProductCategoryByIDService(id uint) (*ProductCategory, error) {
	return GetProductCategoryByID(id)
}
func UpdateProductCategoryService(data *ProductCategory) error { return UpdateProductCategory(data) }
func DeleteProductCategoryService(id uint) error               { return DeleteProductCategory(id) }

func CreateUoMCategoryService(data *UoMCategory) error        { return CreateUoMCategory(data) }
func GetAllUoMCategoryService() ([]UoMCategory, error)        { return GetAllUoMCategory() }
func GetUoMCategoryByIDService(id uint) (*UoMCategory, error) { return GetUoMCategoryByID(id) }
func UpdateUoMCategoryService(data *UoMCategory) error        { return UpdateUoMCategory(data) }
func DeleteUoMCategoryService(id uint) error                  { return DeleteUoMCategory(id) }

func CreateUoMService(data *UoM) error        { return CreateUoM(data) }
func GetAllUoMService() ([]UoM, error)        { return GetAllUoM() }
func GetUoMByIDService(id uint) (*UoM, error) { return GetUoMByID(id) }
func UpdateUoMService(data *UoM) error        { return UpdateUoM(data) }
func DeleteUoMService(id uint) error          { return DeleteUoM(id) }

func CreateProductTemplateService(data *ProductTemplate) error { return CreateProductTemplate(data) }
func GetAllProductTemplateService() ([]ProductTemplate, error) { return GetAllProductTemplate() }
func GetProductTemplateByIDService(id uint) (*ProductTemplate, error) {
	return GetProductTemplateByID(id)
}
func UpdateProductTemplateService(data *ProductTemplate) error { return UpdateProductTemplate(data) }
func DeleteProductTemplateService(id uint) error               { return DeleteProductTemplate(id) }

func CreateProductAttributeService(data *ProductAttribute) error { return CreateProductAttribute(data) }
func GetAllProductAttributeService() ([]ProductAttribute, error) { return GetAllProductAttribute() }
func GetProductAttributeByIDService(id uint) (*ProductAttribute, error) {
	return GetProductAttributeByID(id)
}
func UpdateProductAttributeService(data *ProductAttribute) error { return UpdateProductAttribute(data) }
func DeleteProductAttributeService(id uint) error                { return DeleteProductAttribute(id) }

func CreateProductAttributeValueService(data *ProductAttributeValue) error {
	return CreateProductAttributeValue(data)
}
func GetAllProductAttributeValueService() ([]ProductAttributeValue, error) {
	return GetAllProductAttributeValue()
}
func GetProductAttributeValueByIDService(id uint) (*ProductAttributeValue, error) {
	return GetProductAttributeValueByID(id)
}
func UpdateProductAttributeValueService(data *ProductAttributeValue) error {
	return UpdateProductAttributeValue(data)
}
func DeleteProductAttributeValueService(id uint) error { return DeleteProductAttributeValue(id) }

func CreateStockWarehouseService(data *StockWarehouse) error        { return CreateStockWarehouse(data) }
func GetAllStockWarehouseService() ([]StockWarehouse, error)        { return GetAllStockWarehouse() }
func GetStockWarehouseByIDService(id uint) (*StockWarehouse, error) { return GetStockWarehouseByID(id) }
func UpdateStockWarehouseService(data *StockWarehouse) error        { return UpdateStockWarehouse(data) }
func DeleteStockWarehouseService(id uint) error                     { return DeleteStockWarehouse(id) }

func CreateStockLocationService(data *StockLocation) error        { return CreateStockLocation(data) }
func GetAllStockLocationService() ([]StockLocation, error)        { return GetAllStockLocation() }
func GetStockLocationByIDService(id uint) (*StockLocation, error) { return GetStockLocationByID(id) }
func UpdateStockLocationService(data *StockLocation) error        { return UpdateStockLocation(data) }
func DeleteStockLocationService(id uint) error                    { return DeleteStockLocation(id) }

func CreateStockPickingService(data *StockPicking) error        { return CreateStockPicking(data) }
func GetAllStockPickingService() ([]StockPicking, error)        { return GetAllStockPicking() }
func GetStockPickingByIDService(id uint) (*StockPicking, error) { return GetStockPickingByID(id) }
func UpdateStockPickingService(data *StockPicking) error        { return UpdateStockPicking(data) }
func DeleteStockPickingService(id uint) error                   { return DeleteStockPicking(id) }

func CreateStockLotService(data *StockLot) error        { return CreateStockLot(data) }
func GetAllStockLotService() ([]StockLot, error)        { return GetAllStockLot() }
func GetStockLotByIDService(id uint) (*StockLot, error) { return GetStockLotByID(id) }
func UpdateStockLotService(data *StockLot) error        { return UpdateStockLot(data) }
func DeleteStockLotService(id uint) error               { return DeleteStockLot(id) }

func CreateStockQuantService(data *StockQuant) error        { return CreateStockQuant(data) }
func GetAllStockQuantService() ([]StockQuant, error)        { return GetAllStockQuant() }
func GetStockQuantByIDService(id uint) (*StockQuant, error) { return GetStockQuantByID(id) }
func UpdateStockQuantService(data *StockQuant) error        { return UpdateStockQuant(data) }
func DeleteStockQuantService(id uint) error                 { return DeleteStockQuant(id) }

func CreateStockPutawayRuleService(data *StockPutawayRule) error { return CreateStockPutawayRule(data) }
func GetAllStockPutawayRuleService() ([]StockPutawayRule, error) { return GetAllStockPutawayRule() }
func GetStockPutawayRuleByIDService(id uint) (*StockPutawayRule, error) {
	return GetStockPutawayRuleByID(id)
}
func UpdateStockPutawayRuleService(data *StockPutawayRule) error { return UpdateStockPutawayRule(data) }
func DeleteStockPutawayRuleService(id uint) error                { return DeleteStockPutawayRule(id) }

func CreateStockValuationLayerService(data *StockValuationLayer) error {
	return CreateStockValuationLayer(data)
}
func GetAllStockValuationLayerService() ([]StockValuationLayer, error) {
	return GetAllStockValuationLayer()
}
func GetStockValuationLayerByIDService(id uint) (*StockValuationLayer, error) {
	return GetStockValuationLayerByID(id)
}
func UpdateStockValuationLayerService(data *StockValuationLayer) error {
	return UpdateStockValuationLayer(data)
}
func DeleteStockValuationLayerService(id uint) error { return DeleteStockValuationLayer(id) }
