package rental

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateRentalOrder godoc
// @Summary Create a new RentalOrder
// @Description Create a new RentalOrder in the system
// @Tags sales-rental
// @Accept json
// @Produce json
// @Success 201 {object} RentalOrder
// @Param request body RentalOrder true "Payload"
// @Router /api/sales/rental [post]
// @Security BearerAuth
func CreateRentalOrderHandler(c echo.Context) error {
	var data RentalOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateRentalOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllRentalOrder godoc
// @Summary Get all RentalOrder
// @Description Retrieve a list of all RentalOrder
// @Tags sales-rental
// @Produce json
// @Success 200 {object} []RentalOrder
// @Router /api/sales/rental [get]
// @Security BearerAuth
func GetAllRentalOrderHandler(c echo.Context) error {
	data, err := GetAllRentalOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetRentalOrderByID godoc
// @Summary Get a RentalOrder by ID
// @Description Retrieve a specific RentalOrder by its ID
// @Tags sales-rental
// @Produce json
// @Param id path int true "RentalOrder ID"
// @Success 200 {object} RentalOrder
// @Router /api/sales/rental/{id} [get]
// @Security BearerAuth
func GetRentalOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRentalOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateRentalOrder godoc
// @Summary Update a RentalOrder
// @Description Update an existing RentalOrder
// @Tags sales-rental
// @Accept json
// @Produce json
// @Param id path int true "RentalOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental/{id} [put]
// @Security BearerAuth
func UpdateRentalOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRentalOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateRentalOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteRentalOrder godoc
// @Summary Delete a RentalOrder
// @Description Delete a RentalOrder by ID
// @Tags sales-rental
// @Produce json
// @Param id path int true "RentalOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental/{id} [delete]
// @Security BearerAuth
func DeleteRentalOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteRentalOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create RentalOrderLine
// @Description Create a new RentalOrderLine
// @Tags sales-rental
// @Accept json
// @Produce json
// @Success 201 {object} RentalOrderLine
// @Param request body RentalOrderLine true "Payload"
// @Router /api/sales/rental/rentalorderline [post]
// @Security BearerAuth
func CreateRentalOrderLineHandler(c echo.Context) error {
	var data RentalOrderLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateRentalOrderLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all RentalOrderLine
// @Description Retrieve a list of all RentalOrderLine
// @Tags sales-rental
// @Produce json
// @Success 200 {object} RentalOrderLine
// @Router /api/sales/rental/rentalorderline [get]
// @Security BearerAuth
func GetAllRentalOrderLineHandler(c echo.Context) error {
	data, err := GetAllRentalOrderLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetRentalOrderLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRentalOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update RentalOrderLine
// @Description Update an existing RentalOrderLine
// @Tags sales-rental
// @Accept json
// @Produce json
// @Param id path int true "RentalOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental/rentalorderline/{id} [put]
// @Security BearerAuth
func UpdateRentalOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRentalOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateRentalOrderLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete RentalOrderLine
// @Description Delete RentalOrderLine by ID
// @Tags sales-rental
// @Produce json
// @Param id path int true "RentalOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental/rentalorderline/{id} [delete]
// @Security BearerAuth
func DeleteRentalOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteRentalOrderLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


