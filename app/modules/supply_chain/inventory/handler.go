package inventory

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateProduct godoc
// @Summary Create a new Product
// @Description Create a new Product in the system
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory [post]
// @Security BearerAuth
func CreateProductHandler(c echo.Context) error {
	var data Product
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateProductService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllProduct godoc
// @Summary Get all Product
// @Description Retrieve a list of all Product
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory [get]
// @Security BearerAuth
func GetAllProductHandler(c echo.Context) error {
	data, err := GetAllProductService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetProductByID godoc
// @Summary Get a Product by ID
// @Description Retrieve a specific Product by its ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/{id} [get]
// @Security BearerAuth
func GetProductByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateProduct godoc
// @Summary Update a Product
// @Description Update an existing Product
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/{id} [put]
// @Security BearerAuth
func UpdateProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateProductService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteProduct godoc
// @Summary Delete a Product
// @Description Delete a Product by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/{id} [delete]
// @Security BearerAuth
func DeleteProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProductService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create ProductCategory
// @Description Create a new ProductCategory
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productcategory [post]
// @Security BearerAuth
func CreateProductCategoryHandler(c echo.Context) error {
	var data ProductCategory
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateProductCategoryService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all ProductCategory
// @Description Retrieve a list of all ProductCategory
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productcategory [get]
// @Security BearerAuth
func GetAllProductCategoryHandler(c echo.Context) error {
	data, err := GetAllProductCategoryService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetProductCategoryByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductCategoryByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update ProductCategory
// @Description Update an existing ProductCategory
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "ProductCategory ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productcategory/{id} [put]
// @Security BearerAuth
func UpdateProductCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductCategoryByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateProductCategoryService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete ProductCategory
// @Description Delete ProductCategory by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "ProductCategory ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productcategory/{id} [delete]
// @Security BearerAuth
func DeleteProductCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProductCategoryService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create UoMCategory
// @Description Create a new UoMCategory
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uomcategory [post]
// @Security BearerAuth
func CreateUoMCategoryHandler(c echo.Context) error {
	var data UoMCategory
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateUoMCategoryService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all UoMCategory
// @Description Retrieve a list of all UoMCategory
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uomcategory [get]
// @Security BearerAuth
func GetAllUoMCategoryHandler(c echo.Context) error {
	data, err := GetAllUoMCategoryService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetUoMCategoryByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUoMCategoryByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update UoMCategory
// @Description Update an existing UoMCategory
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "UoMCategory ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uomcategory/{id} [put]
// @Security BearerAuth
func UpdateUoMCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUoMCategoryByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateUoMCategoryService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete UoMCategory
// @Description Delete UoMCategory by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "UoMCategory ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uomcategory/{id} [delete]
// @Security BearerAuth
func DeleteUoMCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteUoMCategoryService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create UoM
// @Description Create a new UoM
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uom [post]
// @Security BearerAuth
func CreateUoMHandler(c echo.Context) error {
	var data UoM
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateUoMService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all UoM
// @Description Retrieve a list of all UoM
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uom [get]
// @Security BearerAuth
func GetAllUoMHandler(c echo.Context) error {
	data, err := GetAllUoMService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetUoMByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUoMByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update UoM
// @Description Update an existing UoM
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "UoM ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uom/{id} [put]
// @Security BearerAuth
func UpdateUoMHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUoMByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateUoMService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete UoM
// @Description Delete UoM by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "UoM ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/uom/{id} [delete]
// @Security BearerAuth
func DeleteUoMHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteUoMService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create ProductTemplate
// @Description Create a new ProductTemplate
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/producttemplate [post]
// @Security BearerAuth
func CreateProductTemplateHandler(c echo.Context) error {
	var data ProductTemplate
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateProductTemplateService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all ProductTemplate
// @Description Retrieve a list of all ProductTemplate
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/producttemplate [get]
// @Security BearerAuth
func GetAllProductTemplateHandler(c echo.Context) error {
	data, err := GetAllProductTemplateService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetProductTemplateByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductTemplateByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update ProductTemplate
// @Description Update an existing ProductTemplate
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "ProductTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/producttemplate/{id} [put]
// @Security BearerAuth
func UpdateProductTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductTemplateByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateProductTemplateService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete ProductTemplate
// @Description Delete ProductTemplate by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "ProductTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/producttemplate/{id} [delete]
// @Security BearerAuth
func DeleteProductTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProductTemplateService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create ProductAttribute
// @Description Create a new ProductAttribute
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattribute [post]
// @Security BearerAuth
func CreateProductAttributeHandler(c echo.Context) error {
	var data ProductAttribute
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateProductAttributeService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all ProductAttribute
// @Description Retrieve a list of all ProductAttribute
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattribute [get]
// @Security BearerAuth
func GetAllProductAttributeHandler(c echo.Context) error {
	data, err := GetAllProductAttributeService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetProductAttributeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductAttributeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update ProductAttribute
// @Description Update an existing ProductAttribute
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "ProductAttribute ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattribute/{id} [put]
// @Security BearerAuth
func UpdateProductAttributeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductAttributeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateProductAttributeService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete ProductAttribute
// @Description Delete ProductAttribute by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "ProductAttribute ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattribute/{id} [delete]
// @Security BearerAuth
func DeleteProductAttributeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProductAttributeService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create ProductAttributeValue
// @Description Create a new ProductAttributeValue
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattributevalue [post]
// @Security BearerAuth
func CreateProductAttributeValueHandler(c echo.Context) error {
	var data ProductAttributeValue
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateProductAttributeValueService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all ProductAttributeValue
// @Description Retrieve a list of all ProductAttributeValue
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattributevalue [get]
// @Security BearerAuth
func GetAllProductAttributeValueHandler(c echo.Context) error {
	data, err := GetAllProductAttributeValueService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetProductAttributeValueByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductAttributeValueByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update ProductAttributeValue
// @Description Update an existing ProductAttributeValue
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "ProductAttributeValue ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattributevalue/{id} [put]
// @Security BearerAuth
func UpdateProductAttributeValueHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductAttributeValueByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateProductAttributeValueService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete ProductAttributeValue
// @Description Delete ProductAttributeValue by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "ProductAttributeValue ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/productattributevalue/{id} [delete]
// @Security BearerAuth
func DeleteProductAttributeValueHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProductAttributeValueService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create StockWarehouse
// @Description Create a new StockWarehouse
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockwarehouse [post]
// @Security BearerAuth
func CreateStockWarehouseHandler(c echo.Context) error {
	var data StockWarehouse
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStockWarehouseService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all StockWarehouse
// @Description Retrieve a list of all StockWarehouse
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockwarehouse [get]
// @Security BearerAuth
func GetAllStockWarehouseHandler(c echo.Context) error {
	data, err := GetAllStockWarehouseService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStockWarehouseByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockWarehouseByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update StockWarehouse
// @Description Update an existing StockWarehouse
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "StockWarehouse ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockwarehouse/{id} [put]
// @Security BearerAuth
func UpdateStockWarehouseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockWarehouseByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStockWarehouseService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete StockWarehouse
// @Description Delete StockWarehouse by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "StockWarehouse ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockwarehouse/{id} [delete]
// @Security BearerAuth
func DeleteStockWarehouseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStockWarehouseService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create StockLocation
// @Description Create a new StockLocation
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklocation [post]
// @Security BearerAuth
func CreateStockLocationHandler(c echo.Context) error {
	var data StockLocation
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStockLocationService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all StockLocation
// @Description Retrieve a list of all StockLocation
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklocation [get]
// @Security BearerAuth
func GetAllStockLocationHandler(c echo.Context) error {
	data, err := GetAllStockLocationService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStockLocationByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockLocationByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update StockLocation
// @Description Update an existing StockLocation
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "StockLocation ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklocation/{id} [put]
// @Security BearerAuth
func UpdateStockLocationHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockLocationByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStockLocationService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete StockLocation
// @Description Delete StockLocation by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "StockLocation ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklocation/{id} [delete]
// @Security BearerAuth
func DeleteStockLocationHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStockLocationService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create StockPicking
// @Description Create a new StockPicking
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockpicking [post]
// @Security BearerAuth
func CreateStockPickingHandler(c echo.Context) error {
	var data StockPicking
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStockPickingService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all StockPicking
// @Description Retrieve a list of all StockPicking
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockpicking [get]
// @Security BearerAuth
func GetAllStockPickingHandler(c echo.Context) error {
	data, err := GetAllStockPickingService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStockPickingByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockPickingByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update StockPicking
// @Description Update an existing StockPicking
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "StockPicking ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockpicking/{id} [put]
// @Security BearerAuth
func UpdateStockPickingHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockPickingByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStockPickingService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete StockPicking
// @Description Delete StockPicking by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "StockPicking ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockpicking/{id} [delete]
// @Security BearerAuth
func DeleteStockPickingHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStockPickingService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create StockLot
// @Description Create a new StockLot
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklot [post]
// @Security BearerAuth
func CreateStockLotHandler(c echo.Context) error {
	var data StockLot
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStockLotService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all StockLot
// @Description Retrieve a list of all StockLot
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklot [get]
// @Security BearerAuth
func GetAllStockLotHandler(c echo.Context) error {
	data, err := GetAllStockLotService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStockLotByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockLotByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update StockLot
// @Description Update an existing StockLot
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "StockLot ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklot/{id} [put]
// @Security BearerAuth
func UpdateStockLotHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockLotByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStockLotService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete StockLot
// @Description Delete StockLot by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "StockLot ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stocklot/{id} [delete]
// @Security BearerAuth
func DeleteStockLotHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStockLotService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create StockQuant
// @Description Create a new StockQuant
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockquant [post]
// @Security BearerAuth
func CreateStockQuantHandler(c echo.Context) error {
	var data StockQuant
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStockQuantService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all StockQuant
// @Description Retrieve a list of all StockQuant
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockquant [get]
// @Security BearerAuth
func GetAllStockQuantHandler(c echo.Context) error {
	data, err := GetAllStockQuantService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStockQuantByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockQuantByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update StockQuant
// @Description Update an existing StockQuant
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "StockQuant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockquant/{id} [put]
// @Security BearerAuth
func UpdateStockQuantHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockQuantByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStockQuantService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete StockQuant
// @Description Delete StockQuant by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "StockQuant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockquant/{id} [delete]
// @Security BearerAuth
func DeleteStockQuantHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStockQuantService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create StockPutawayRule
// @Description Create a new StockPutawayRule
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockputawayrule [post]
// @Security BearerAuth
func CreateStockPutawayRuleHandler(c echo.Context) error {
	var data StockPutawayRule
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStockPutawayRuleService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all StockPutawayRule
// @Description Retrieve a list of all StockPutawayRule
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockputawayrule [get]
// @Security BearerAuth
func GetAllStockPutawayRuleHandler(c echo.Context) error {
	data, err := GetAllStockPutawayRuleService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStockPutawayRuleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockPutawayRuleByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update StockPutawayRule
// @Description Update an existing StockPutawayRule
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "StockPutawayRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockputawayrule/{id} [put]
// @Security BearerAuth
func UpdateStockPutawayRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockPutawayRuleByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStockPutawayRuleService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete StockPutawayRule
// @Description Delete StockPutawayRule by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "StockPutawayRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockputawayrule/{id} [delete]
// @Security BearerAuth
func DeleteStockPutawayRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStockPutawayRuleService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create StockValuationLayer
// @Description Create a new StockValuationLayer
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockvaluationlayer [post]
// @Security BearerAuth
func CreateStockValuationLayerHandler(c echo.Context) error {
	var data StockValuationLayer
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStockValuationLayerService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all StockValuationLayer
// @Description Retrieve a list of all StockValuationLayer
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockvaluationlayer [get]
// @Security BearerAuth
func GetAllStockValuationLayerHandler(c echo.Context) error {
	data, err := GetAllStockValuationLayerService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStockValuationLayerByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockValuationLayerByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update StockValuationLayer
// @Description Update an existing StockValuationLayer
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "StockValuationLayer ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockvaluationlayer/{id} [put]
// @Security BearerAuth
func UpdateStockValuationLayerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStockValuationLayerByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStockValuationLayerService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete StockValuationLayer
// @Description Delete StockValuationLayer by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "StockValuationLayer ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/stockvaluationlayer/{id} [delete]
// @Security BearerAuth
func DeleteStockValuationLayerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStockValuationLayerService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
