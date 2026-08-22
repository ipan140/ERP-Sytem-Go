package sales_core

import (
	"ERP-System/app/modules/sales"
	"ERP-System/app/modules/services/project"
	"ERP-System/config"
	"fmt"
)

func CreateSaleOrderService(data *SaleOrder) error        { return CreateSaleOrder(data) }
func GetAllSaleOrderService() ([]SaleOrder, error)        { return GetAllSaleOrder() }
func GetSaleOrderByIDService(id uint) (*SaleOrder, error) { return GetSaleOrderByID(id) }
func UpdateSaleOrderService(data *SaleOrder) error {
	if data.State == "sale" {
		project.AutoCreateProjectFromSales(data.ID, data.PartnerID, data.Name)
		
		// [RabbitMQ] - Fase 4: Publish Event agar modul Finance dan Supply Chain dapat menangkapnya
		orderIDStr := fmt.Sprintf("%d", data.ID)
		customerIDStr := fmt.Sprintf("%d", data.PartnerID)
		_ = sales.PublishOrderCompletedEvent(orderIDStr, data.AmountTotal, customerIDStr)
	}
	return UpdateSaleOrder(data)
}
func DeleteSaleOrderService(id uint) error { return DeleteSaleOrder(id) }

func CreatePricelistService(data *Pricelist) error        { return CreatePricelist(data) }
func GetAllPricelistService() ([]Pricelist, error)        { return GetAllPricelist() }
func GetPricelistByIDService(id uint) (*Pricelist, error) { return GetPricelistByID(id) }
func UpdatePricelistService(data *Pricelist) error        { return UpdatePricelist(data) }
func DeletePricelistService(id uint) error                { return DeletePricelist(id) }

func CreatePricelistItemService(data *PricelistItem) error        { return CreatePricelistItem(data) }
func GetAllPricelistItemService() ([]PricelistItem, error)        { return GetAllPricelistItem() }
func GetPricelistItemByIDService(id uint) (*PricelistItem, error) { return GetPricelistItemByID(id) }
func UpdatePricelistItemService(data *PricelistItem) error        { return UpdatePricelistItem(data) }
func DeletePricelistItemService(id uint) error                    { return DeletePricelistItem(id) }

func CreateQuotationTemplateService(data *QuotationTemplate) error {
	return CreateQuotationTemplate(data)
}
func GetAllQuotationTemplateService() ([]QuotationTemplate, error) { return GetAllQuotationTemplate() }
func GetQuotationTemplateByIDService(id uint) (*QuotationTemplate, error) {
	return GetQuotationTemplateByID(id)
}
func UpdateQuotationTemplateService(data *QuotationTemplate) error {
	return UpdateQuotationTemplate(data)
}
func DeleteQuotationTemplateService(id uint) error { return DeleteQuotationTemplate(id) }

func CreateDeliveryMethodService(data *DeliveryMethod) error        { return CreateDeliveryMethod(data) }
func GetAllDeliveryMethodService() ([]DeliveryMethod, error)        { return GetAllDeliveryMethod() }
func GetDeliveryMethodByIDService(id uint) (*DeliveryMethod, error) { return GetDeliveryMethodByID(id) }
func UpdateDeliveryMethodService(data *DeliveryMethod) error        { return UpdateDeliveryMethod(data) }
func DeleteDeliveryMethodService(id uint) error                     { return DeleteDeliveryMethod(id) }

func CreateSaleOrderLineService(data *SaleOrderLine) error {
	data.SubTotal = data.Quantity * data.UnitPrice * (1 - (data.Discount / 100))
	if err := CreateSaleOrderLine(data); err != nil {
		return err
	}
	return RecalculateSaleOrder(data.OrderID)
}
func GetAllSaleOrderLineService() ([]SaleOrderLine, error)        { return GetAllSaleOrderLine() }
func GetSaleOrderLineByIDService(id uint) (*SaleOrderLine, error) { return GetSaleOrderLineByID(id) }
func UpdateSaleOrderLineService(data *SaleOrderLine) error {
	data.SubTotal = data.Quantity * data.UnitPrice * (1 - (data.Discount / 100))
	if err := UpdateSaleOrderLine(data); err != nil {
		return err
	}
	return RecalculateSaleOrder(data.OrderID)
}
func DeleteSaleOrderLineService(id uint) error {
	line, err := GetSaleOrderLineByID(id)
	if err != nil {
		return err
	}
	orderID := line.OrderID
	if err := DeleteSaleOrderLine(id); err != nil {
		return err
	}
	return RecalculateSaleOrder(orderID)
}

func RecalculateSaleOrder(orderID uint) error {
	var lines []SaleOrderLine
	if err := config.DB.Where("order_id = ?", orderID).Find(&lines).Error; err != nil {
		return err
	}

	var amountUntaxed float64
	for _, line := range lines {
		amountUntaxed += line.SubTotal
	}

	amountTax := amountUntaxed * 0.11 // Asumsi PPN 11%
	amountTotal := amountUntaxed + amountTax

	return config.DB.Model(&SaleOrder{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"amount_untaxed": amountUntaxed,
		"amount_tax":     amountTax,
		"amount_total":   amountTotal,
	}).Error
}
