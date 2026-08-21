package sales_core

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSaleOrder godoc
// @Summary Create a new SaleOrder
// @Description Create a new SaleOrder in the system
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/sales_core [post]
// @Security BearerAuth
func CreateSaleOrderHandler(c echo.Context) error {
	var data SaleOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSaleOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSaleOrder godoc
// @Summary Get all SaleOrder
// @Description Retrieve a list of all SaleOrder
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core [get]
// @Security BearerAuth
func GetAllSaleOrderHandler(c echo.Context) error {
	data, err := GetAllSaleOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSaleOrderByID godoc
// @Summary Get a SaleOrder by ID
// @Description Retrieve a specific SaleOrder by its ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [get]
// @Security BearerAuth
func GetSaleOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSaleOrder godoc
// @Summary Update a SaleOrder
// @Description Update an existing SaleOrder
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [put]
// @Security BearerAuth
func UpdateSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSaleOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSaleOrder godoc
// @Summary Delete a SaleOrder
// @Description Delete a SaleOrder by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [delete]
// @Security BearerAuth
func DeleteSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSaleOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create Pricelist
// @Description Create a new Pricelist
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelist [post]
// @Security BearerAuth
func CreatePricelistHandler(c echo.Context) error {
	var data Pricelist
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePricelistService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Pricelist
// @Description Retrieve a list of all Pricelist
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelist [get]
// @Security BearerAuth
func GetAllPricelistHandler(c echo.Context) error {
	data, err := GetAllPricelistService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPricelistByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Pricelist
// @Description Update an existing Pricelist
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "Pricelist ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelist/{id} [put]
// @Security BearerAuth
func UpdatePricelistHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePricelistService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Pricelist
// @Description Delete Pricelist by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "Pricelist ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelist/{id} [delete]
// @Security BearerAuth
func DeletePricelistHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePricelistService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create PricelistItem
// @Description Create a new PricelistItem
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelistitem [post]
// @Security BearerAuth
func CreatePricelistItemHandler(c echo.Context) error {
	var data PricelistItem
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePricelistItemService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PricelistItem
// @Description Retrieve a list of all PricelistItem
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelistitem [get]
// @Security BearerAuth
func GetAllPricelistItemHandler(c echo.Context) error {
	data, err := GetAllPricelistItemService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPricelistItemByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistItemByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PricelistItem
// @Description Update an existing PricelistItem
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "PricelistItem ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelistitem/{id} [put]
// @Security BearerAuth
func UpdatePricelistItemHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistItemByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePricelistItemService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PricelistItem
// @Description Delete PricelistItem by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "PricelistItem ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelistitem/{id} [delete]
// @Security BearerAuth
func DeletePricelistItemHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePricelistItemService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create QuotationTemplate
// @Description Create a new QuotationTemplate
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/sales_core/quotationtemplate [post]
// @Security BearerAuth
func CreateQuotationTemplateHandler(c echo.Context) error {
	var data QuotationTemplate
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateQuotationTemplateService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all QuotationTemplate
// @Description Retrieve a list of all QuotationTemplate
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/quotationtemplate [get]
// @Security BearerAuth
func GetAllQuotationTemplateHandler(c echo.Context) error {
	data, err := GetAllQuotationTemplateService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetQuotationTemplateByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQuotationTemplateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update QuotationTemplate
// @Description Update an existing QuotationTemplate
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "QuotationTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/quotationtemplate/{id} [put]
// @Security BearerAuth
func UpdateQuotationTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQuotationTemplateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateQuotationTemplateService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete QuotationTemplate
// @Description Delete QuotationTemplate by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "QuotationTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/quotationtemplate/{id} [delete]
// @Security BearerAuth
func DeleteQuotationTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteQuotationTemplateService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create DeliveryMethod
// @Description Create a new DeliveryMethod
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/sales_core/deliverymethod [post]
// @Security BearerAuth
func CreateDeliveryMethodHandler(c echo.Context) error {
	var data DeliveryMethod
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateDeliveryMethodService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all DeliveryMethod
// @Description Retrieve a list of all DeliveryMethod
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/deliverymethod [get]
// @Security BearerAuth
func GetAllDeliveryMethodHandler(c echo.Context) error {
	data, err := GetAllDeliveryMethodService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetDeliveryMethodByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDeliveryMethodByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update DeliveryMethod
// @Description Update an existing DeliveryMethod
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "DeliveryMethod ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/deliverymethod/{id} [put]
// @Security BearerAuth
func UpdateDeliveryMethodHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDeliveryMethodByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateDeliveryMethodService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete DeliveryMethod
// @Description Delete DeliveryMethod by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "DeliveryMethod ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/deliverymethod/{id} [delete]
// @Security BearerAuth
func DeleteDeliveryMethodHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteDeliveryMethodService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create SaleOrderLine
// @Description Create a new SaleOrderLine
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/sales_core/saleorderline [post]
// @Security BearerAuth
func CreateSaleOrderLineHandler(c echo.Context) error {
	var data SaleOrderLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSaleOrderLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all SaleOrderLine
// @Description Retrieve a list of all SaleOrderLine
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/saleorderline [get]
// @Security BearerAuth
func GetAllSaleOrderLineHandler(c echo.Context) error {
	data, err := GetAllSaleOrderLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSaleOrderLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update SaleOrderLine
// @Description Update an existing SaleOrderLine
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/saleorderline/{id} [put]
// @Security BearerAuth
func UpdateSaleOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSaleOrderLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete SaleOrderLine
// @Description Delete SaleOrderLine by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/saleorderline/{id} [delete]
// @Security BearerAuth
func DeleteSaleOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSaleOrderLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
