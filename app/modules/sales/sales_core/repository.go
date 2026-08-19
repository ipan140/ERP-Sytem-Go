package sales_core

import (
	"ERP-System/config"
)

func CreateSaleOrder(data *SaleOrder) error {
	return config.DB.Create(data).Error
}

func GetAllSaleOrder() ([]SaleOrder, error) {
	var list []SaleOrder
	err := config.DB.Find(&list).Error
	return list, err
}

func GetSaleOrderByID(id uint) (*SaleOrder, error) {
	var data SaleOrder
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateSaleOrder(data *SaleOrder) error {
	return config.DB.Save(data).Error
}

func DeleteSaleOrder(id uint) error {
	return config.DB.Delete(&SaleOrder{}, id).Error
}

func CreatePricelist(data *Pricelist) error { return config.DB.Create(data).Error }
func GetAllPricelist() ([]Pricelist, error) { var list []Pricelist; err := config.DB.Find(&list).Error; return list, err }
func GetPricelistByID(id uint) (*Pricelist, error) { var data Pricelist; err := config.DB.First(&data, id).Error; return &data, err }
func UpdatePricelist(data *Pricelist) error { return config.DB.Save(data).Error }
func DeletePricelist(id uint) error { return config.DB.Delete(&Pricelist{}, id).Error }

func CreatePricelistItem(data *PricelistItem) error { return config.DB.Create(data).Error }
func GetAllPricelistItem() ([]PricelistItem, error) { var list []PricelistItem; err := config.DB.Find(&list).Error; return list, err }
func GetPricelistItemByID(id uint) (*PricelistItem, error) { var data PricelistItem; err := config.DB.First(&data, id).Error; return &data, err }
func UpdatePricelistItem(data *PricelistItem) error { return config.DB.Save(data).Error }
func DeletePricelistItem(id uint) error { return config.DB.Delete(&PricelistItem{}, id).Error }

func CreateQuotationTemplate(data *QuotationTemplate) error { return config.DB.Create(data).Error }
func GetAllQuotationTemplate() ([]QuotationTemplate, error) { var list []QuotationTemplate; err := config.DB.Find(&list).Error; return list, err }
func GetQuotationTemplateByID(id uint) (*QuotationTemplate, error) { var data QuotationTemplate; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateQuotationTemplate(data *QuotationTemplate) error { return config.DB.Save(data).Error }
func DeleteQuotationTemplate(id uint) error { return config.DB.Delete(&QuotationTemplate{}, id).Error }

func CreateDeliveryMethod(data *DeliveryMethod) error { return config.DB.Create(data).Error }
func GetAllDeliveryMethod() ([]DeliveryMethod, error) { var list []DeliveryMethod; err := config.DB.Find(&list).Error; return list, err }
func GetDeliveryMethodByID(id uint) (*DeliveryMethod, error) { var data DeliveryMethod; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateDeliveryMethod(data *DeliveryMethod) error { return config.DB.Save(data).Error }
func DeleteDeliveryMethod(id uint) error { return config.DB.Delete(&DeliveryMethod{}, id).Error }

func CreateSaleOrderLine(data *SaleOrderLine) error { return config.DB.Create(data).Error }
func GetAllSaleOrderLine() ([]SaleOrderLine, error) { var list []SaleOrderLine; err := config.DB.Find(&list).Error; return list, err }
func GetSaleOrderLineByID(id uint) (*SaleOrderLine, error) { var data SaleOrderLine; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateSaleOrderLine(data *SaleOrderLine) error { return config.DB.Save(data).Error }
func DeleteSaleOrderLine(id uint) error { return config.DB.Delete(&SaleOrderLine{}, id).Error }
