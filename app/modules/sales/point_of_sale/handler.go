package point_of_sale

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreatePosSession godoc
// @Summary Create a new PosSession
// @Description Create a new PosSession in the system
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Success 201 {object} PosSession
// @Param request body PosSession true "Payload"
// @Router /api/sales/point_of_sale [post]
// @Security BearerAuth
func CreatePosSessionHandler(c echo.Context) error {
	var data PosSession
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreatePosSessionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllPosSession godoc
// @Summary Get all PosSession
// @Description Retrieve a list of all PosSession
// @Tags sales-point_of_sale
// @Produce json
// @Success 200 {object} []PosSession
// @Router /api/sales/point_of_sale [get]
// @Security BearerAuth
func GetAllPosSessionHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllPosSessionService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")

	data, total, err := GetPaginatedPosSessionService(offset, limit, search, state)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetPosSessionByID godoc
// @Summary Get a PosSession by ID
// @Description Retrieve a specific PosSession by its ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "PosSession ID"
// @Success 200 {object} PosSession
// @Router /api/sales/point_of_sale/{id} [get]
// @Security BearerAuth
func GetPosSessionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosSessionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdatePosSession godoc
// @Summary Update a PosSession
// @Description Update an existing PosSession
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Param id path int true "PosSession ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/{id} [put]
// @Security BearerAuth
func UpdatePosSessionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosSessionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdatePosSessionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeletePosSession godoc
// @Summary Delete a PosSession
// @Description Delete a PosSession by ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "PosSession ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/{id} [delete]
// @Security BearerAuth
func DeletePosSessionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePosSessionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create PosConfig
// @Description Create a new PosConfig
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Success 201 {object} PosConfig
// @Param request body PosConfig true "Payload"
// @Router /api/sales/point_of_sale/posconfig [post]
// @Security BearerAuth
func CreatePosConfigHandler(c echo.Context) error {
	var data PosConfig
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePosConfigService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PosConfig
// @Description Retrieve a list of all PosConfig
// @Tags sales-point_of_sale
// @Produce json
// @Success 200 {object} PosConfig
// @Router /api/sales/point_of_sale/posconfig [get]
// @Security BearerAuth
func GetAllPosConfigHandler(c echo.Context) error {
	data, err := GetAllPosConfigService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPosConfigByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosConfigByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PosConfig
// @Description Update an existing PosConfig
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Param id path int true "PosConfig ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/posconfig/{id} [put]
// @Security BearerAuth
func UpdatePosConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosConfigByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePosConfigService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PosConfig
// @Description Delete PosConfig by ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "PosConfig ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/posconfig/{id} [delete]
// @Security BearerAuth
func DeletePosConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePosConfigService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create PosOrder
// @Description Create a new PosOrder
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Success 201 {object} PosOrder
// @Param request body PosOrder true "Payload"
// @Router /api/sales/point_of_sale/posorder [post]
// @Security BearerAuth
func CreatePosOrderHandler(c echo.Context) error {
	var data PosOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePosOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// CheckoutPosOrderHandler godoc
// @Summary POS Cashier Checkout
// @Description Process a POS cart order, record payment, and automatically deduct warehouse stock
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Param request body PosCheckoutPayload true "Checkout Payload"
// @Success 201 {object} PosOrder
// @Router /api/sales/point_of_sale/checkout [post]
// @Security BearerAuth
func CheckoutPosOrderHandler(c echo.Context) error {
	var payload PosCheckoutPayload
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	order, err := CheckoutPosOrderService(&payload)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to process POS checkout", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Transaction completed successfully", order)
}


// @Summary Get all PosOrder
// @Description Retrieve a list of all PosOrder
// @Tags sales-point_of_sale
// @Produce json
// @Success 200 {object} PosOrder
// @Router /api/sales/point_of_sale/posorder [get]
// @Security BearerAuth
func GetAllPosOrderHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllPosOrderService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")

	data, total, err := GetPaginatedPosOrderService(offset, limit, search, state)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Retrieved successfully", data, meta)
}
func GetPosOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PosOrder
// @Description Update an existing PosOrder
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Param id path int true "PosOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/posorder/{id} [put]
// @Security BearerAuth
func UpdatePosOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePosOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PosOrder
// @Description Delete PosOrder by ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "PosOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/posorder/{id} [delete]
// @Security BearerAuth
func DeletePosOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePosOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create PosOrderLine
// @Description Create a new PosOrderLine
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Success 201 {object} PosOrderLine
// @Param request body PosOrderLine true "Payload"
// @Router /api/sales/point_of_sale/posorderline [post]
// @Security BearerAuth
func CreatePosOrderLineHandler(c echo.Context) error {
	var data PosOrderLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePosOrderLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PosOrderLine
// @Description Retrieve a list of all PosOrderLine
// @Tags sales-point_of_sale
// @Produce json
// @Success 200 {object} PosOrderLine
// @Router /api/sales/point_of_sale/posorderline [get]
// @Security BearerAuth
func GetAllPosOrderLineHandler(c echo.Context) error {
	data, err := GetAllPosOrderLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPosOrderLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PosOrderLine
// @Description Update an existing PosOrderLine
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Param id path int true "PosOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/posorderline/{id} [put]
// @Security BearerAuth
func UpdatePosOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePosOrderLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PosOrderLine
// @Description Delete PosOrderLine by ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "PosOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/posorderline/{id} [delete]
// @Security BearerAuth
func DeletePosOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePosOrderLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create PosPayment
// @Description Create a new PosPayment
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Success 201 {object} PosPayment
// @Param request body PosPayment true "Payload"
// @Router /api/sales/point_of_sale/pospayment [post]
// @Security BearerAuth
func CreatePosPaymentHandler(c echo.Context) error {
	var data PosPayment
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePosPaymentService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PosPayment
// @Description Retrieve a list of all PosPayment
// @Tags sales-point_of_sale
// @Produce json
// @Success 200 {object} PosPayment
// @Router /api/sales/point_of_sale/pospayment [get]
// @Security BearerAuth
func GetAllPosPaymentHandler(c echo.Context) error {
	data, err := GetAllPosPaymentService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPosPaymentByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosPaymentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PosPayment
// @Description Update an existing PosPayment
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Param id path int true "PosPayment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/pospayment/{id} [put]
// @Security BearerAuth
func UpdatePosPaymentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPosPaymentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePosPaymentService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PosPayment
// @Description Delete PosPayment by ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "PosPayment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/pospayment/{id} [delete]
// @Security BearerAuth
func DeletePosPaymentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePosPaymentService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create LoyaltyProgram
// @Description Create a new LoyaltyProgram
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Success 201 {object} LoyaltyProgram
// @Param request body LoyaltyProgram true "Payload"
// @Router /api/sales/point_of_sale/loyaltyprogram [post]
// @Security BearerAuth
func CreateLoyaltyProgramHandler(c echo.Context) error {
	var data LoyaltyProgram
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateLoyaltyProgramService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all LoyaltyProgram
// @Description Retrieve a list of all LoyaltyProgram
// @Tags sales-point_of_sale
// @Produce json
// @Success 200 {object} LoyaltyProgram
// @Router /api/sales/point_of_sale/loyaltyprogram [get]
// @Security BearerAuth
func GetAllLoyaltyProgramHandler(c echo.Context) error {
	data, err := GetAllLoyaltyProgramService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetLoyaltyProgramByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLoyaltyProgramByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update LoyaltyProgram
// @Description Update an existing LoyaltyProgram
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Param id path int true "LoyaltyProgram ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/loyaltyprogram/{id} [put]
// @Security BearerAuth
func UpdateLoyaltyProgramHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLoyaltyProgramByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateLoyaltyProgramService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete LoyaltyProgram
// @Description Delete LoyaltyProgram by ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "LoyaltyProgram ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale/loyaltyprogram/{id} [delete]
// @Security BearerAuth
func DeleteLoyaltyProgramHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLoyaltyProgramService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


