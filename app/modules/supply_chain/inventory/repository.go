package inventory

import (
	"ERP-System/config"
)

func CreateProduct(data *Product) error {
	return config.DB.Create(data).Error
}

func GetAllProduct() ([]Product, error) {
	var list []Product
	err := config.DB.Find(&list).Error
	return list, err
}

func GetProductByID(id uint) (*Product, error) {
	var data Product
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateProduct(data *Product) error {
	return config.DB.Save(data).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&Product{}, id).Error
}

func CreateProductCategory(data *ProductCategory) error { return config.DB.Create(data).Error }
func GetAllProductCategory() ([]ProductCategory, error) {
	var list []ProductCategory
	err := config.DB.Find(&list).Error
	return list, err
}
func GetProductCategoryByID(id uint) (*ProductCategory, error) {
	var data ProductCategory
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateProductCategory(data *ProductCategory) error { return config.DB.Save(data).Error }
func DeleteProductCategory(id uint) error               { return config.DB.Delete(&ProductCategory{}, id).Error }

func CreateUoMCategory(data *UoMCategory) error { return config.DB.Create(data).Error }
func GetAllUoMCategory() ([]UoMCategory, error) {
	var list []UoMCategory
	err := config.DB.Find(&list).Error
	return list, err
}
func GetUoMCategoryByID(id uint) (*UoMCategory, error) {
	var data UoMCategory
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateUoMCategory(data *UoMCategory) error { return config.DB.Save(data).Error }
func DeleteUoMCategory(id uint) error           { return config.DB.Delete(&UoMCategory{}, id).Error }

func CreateUoM(data *UoM) error { return config.DB.Create(data).Error }
func GetAllUoM() ([]UoM, error) { var list []UoM; err := config.DB.Find(&list).Error; return list, err }
func GetUoMByID(id uint) (*UoM, error) {
	var data UoM
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateUoM(data *UoM) error { return config.DB.Save(data).Error }
func DeleteUoM(id uint) error   { return config.DB.Delete(&UoM{}, id).Error }

func CreateProductTemplate(data *ProductTemplate) error { return config.DB.Create(data).Error }
func GetAllProductTemplate() ([]ProductTemplate, error) {
	var list []ProductTemplate
	err := config.DB.Find(&list).Error
	return list, err
}
func GetProductTemplateByID(id uint) (*ProductTemplate, error) {
	var data ProductTemplate
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateProductTemplate(data *ProductTemplate) error { return config.DB.Save(data).Error }
func DeleteProductTemplate(id uint) error               { return config.DB.Delete(&ProductTemplate{}, id).Error }

func CreateProductAttribute(data *ProductAttribute) error { return config.DB.Create(data).Error }
func GetAllProductAttribute() ([]ProductAttribute, error) {
	var list []ProductAttribute
	err := config.DB.Find(&list).Error
	return list, err
}
func GetProductAttributeByID(id uint) (*ProductAttribute, error) {
	var data ProductAttribute
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateProductAttribute(data *ProductAttribute) error { return config.DB.Save(data).Error }
func DeleteProductAttribute(id uint) error                { return config.DB.Delete(&ProductAttribute{}, id).Error }

func CreateProductAttributeValue(data *ProductAttributeValue) error {
	return config.DB.Create(data).Error
}
func GetAllProductAttributeValue() ([]ProductAttributeValue, error) {
	var list []ProductAttributeValue
	err := config.DB.Find(&list).Error
	return list, err
}
func GetProductAttributeValueByID(id uint) (*ProductAttributeValue, error) {
	var data ProductAttributeValue
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateProductAttributeValue(data *ProductAttributeValue) error {
	return config.DB.Save(data).Error
}
func DeleteProductAttributeValue(id uint) error {
	return config.DB.Delete(&ProductAttributeValue{}, id).Error
}

func CreateStockWarehouse(data *StockWarehouse) error { return config.DB.Create(data).Error }
func GetAllStockWarehouse() ([]StockWarehouse, error) {
	var list []StockWarehouse
	err := config.DB.Find(&list).Error
	return list, err
}
func GetStockWarehouseByID(id uint) (*StockWarehouse, error) {
	var data StockWarehouse
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateStockWarehouse(data *StockWarehouse) error { return config.DB.Save(data).Error }
func DeleteStockWarehouse(id uint) error              { return config.DB.Delete(&StockWarehouse{}, id).Error }

func CreateStockLocation(data *StockLocation) error { return config.DB.Create(data).Error }
func GetAllStockLocation() ([]StockLocation, error) {
	var list []StockLocation
	err := config.DB.Find(&list).Error
	return list, err
}
func GetStockLocationByID(id uint) (*StockLocation, error) {
	var data StockLocation
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateStockLocation(data *StockLocation) error { return config.DB.Save(data).Error }
func DeleteStockLocation(id uint) error             { return config.DB.Delete(&StockLocation{}, id).Error }

func CreateStockPicking(data *StockPicking) error { return config.DB.Create(data).Error }
func GetAllStockPicking() ([]StockPicking, error) {
	var list []StockPicking
	err := config.DB.Find(&list).Error
	return list, err
}
func GetStockPickingByID(id uint) (*StockPicking, error) {
	var data StockPicking
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateStockPicking(data *StockPicking) error { return config.DB.Save(data).Error }
func DeleteStockPicking(id uint) error            { return config.DB.Delete(&StockPicking{}, id).Error }

func CreateStockLot(data *StockLot) error { return config.DB.Create(data).Error }
func GetAllStockLot() ([]StockLot, error) {
	var list []StockLot
	err := config.DB.Find(&list).Error
	return list, err
}
func GetStockLotByID(id uint) (*StockLot, error) {
	var data StockLot
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateStockLot(data *StockLot) error { return config.DB.Save(data).Error }
func DeleteStockLot(id uint) error        { return config.DB.Delete(&StockLot{}, id).Error }

func CreateStockQuant(data *StockQuant) error { return config.DB.Create(data).Error }
func GetAllStockQuant() ([]StockQuant, error) {
	var list []StockQuant
	err := config.DB.Find(&list).Error
	return list, err
}
func GetStockQuantByID(id uint) (*StockQuant, error) {
	var data StockQuant
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateStockQuant(data *StockQuant) error { return config.DB.Save(data).Error }
func DeleteStockQuant(id uint) error          { return config.DB.Delete(&StockQuant{}, id).Error }

func CreateStockPutawayRule(data *StockPutawayRule) error { return config.DB.Create(data).Error }
func GetAllStockPutawayRule() ([]StockPutawayRule, error) {
	var list []StockPutawayRule
	err := config.DB.Find(&list).Error
	return list, err
}
func GetStockPutawayRuleByID(id uint) (*StockPutawayRule, error) {
	var data StockPutawayRule
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateStockPutawayRule(data *StockPutawayRule) error { return config.DB.Save(data).Error }
func DeleteStockPutawayRule(id uint) error                { return config.DB.Delete(&StockPutawayRule{}, id).Error }

func CreateStockValuationLayer(data *StockValuationLayer) error { return config.DB.Create(data).Error }
func GetAllStockValuationLayer() ([]StockValuationLayer, error) {
	var list []StockValuationLayer
	err := config.DB.Find(&list).Error
	return list, err
}
func GetStockValuationLayerByID(id uint) (*StockValuationLayer, error) {
	var data StockValuationLayer
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateStockValuationLayer(data *StockValuationLayer) error { return config.DB.Save(data).Error }
func DeleteStockValuationLayer(id uint) error {
	return config.DB.Delete(&StockValuationLayer{}, id).Error
}
