package purchase

import (
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
)

func CreatePurchaseOrderService(data *PurchaseOrder) error {
	return CreatePurchaseOrder(data)
}

func GetAllPurchaseOrderService() ([]PurchaseOrder, error) {
	return GetAllPurchaseOrder()
}

func GetPurchaseOrderByIDService(id uint) (*PurchaseOrder, error) {
	return GetPurchaseOrderByID(id)
}

func UpdatePurchaseOrderService(data *PurchaseOrder) error {
	return UpdatePurchaseOrder(data)
}

func DeletePurchaseOrderService(id uint) error {
	return DeletePurchaseOrder(id)
}

func CreatePurchaseRequisitionService(data *PurchaseRequisition) error {
	return CreatePurchaseRequisition(data)
}
func GetAllPurchaseRequisitionService() ([]PurchaseRequisition, error) {
	return GetAllPurchaseRequisition()
}
func GetPurchaseRequisitionByIDService(id uint) (*PurchaseRequisition, error) {
	return GetPurchaseRequisitionByID(id)
}
func UpdatePurchaseRequisitionService(data *PurchaseRequisition) error {
	return UpdatePurchaseRequisition(data)
}
func DeletePurchaseRequisitionService(id uint) error { return DeletePurchaseRequisition(id) }

func CreateProductSupplierInfoService(data *ProductSupplierInfo) error {
	return CreateProductSupplierInfo(data)
}
func GetAllProductSupplierInfoService() ([]ProductSupplierInfo, error) {
	return GetAllProductSupplierInfo()
}
func GetProductSupplierInfoByIDService(id uint) (*ProductSupplierInfo, error) {
	return GetProductSupplierInfoByID(id)
}
func UpdateProductSupplierInfoService(data *ProductSupplierInfo) error {
	return UpdateProductSupplierInfo(data)
}
func DeleteProductSupplierInfoService(id uint) error { return DeleteProductSupplierInfo(id) }

func CreatePurchaseOrderLineService(data *PurchaseOrderLine) error {
	return CreatePurchaseOrderLine(data)
}
func GetAllPurchaseOrderLineService() ([]PurchaseOrderLine, error) { return GetAllPurchaseOrderLine() }
func GetPurchaseOrderLineByIDService(id uint) (*PurchaseOrderLine, error) {
	return GetPurchaseOrderLineByID(id)
}
func UpdatePurchaseOrderLineService(data *PurchaseOrderLine) error {
	return UpdatePurchaseOrderLine(data)
}
func DeletePurchaseOrderLineService(id uint) error { return DeletePurchaseOrderLine(id) }

func GeneratePurchaseOrderPDFService(poID uint, userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"document_type": "purchase_order", "document_id": poID, "format": "pdf", "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("📄 Event RabbitMQ: Generate PDF Purchase Order %d dikirim ke antrean!", poID)
	}
	return nil
}
