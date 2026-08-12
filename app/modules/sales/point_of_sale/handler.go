package point_of_sale

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreatePosSession godoc
// @Summary Create a new PosSession
// @Description Create a new PosSession in the system
// @Tags sales-point_of_sale
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
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
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/point_of_sale [get]
// @Security BearerAuth
func GetAllPosSessionHandler(c echo.Context) error {
	data, err := GetAllPosSessionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetPosSessionByID godoc
// @Summary Get a PosSession by ID
// @Description Retrieve a specific PosSession by its ID
// @Tags sales-point_of_sale
// @Produce json
// @Param id path int true "PosSession ID"
// @Success 200 {object} map[string]interface{}
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
