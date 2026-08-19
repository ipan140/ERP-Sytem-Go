package plm

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateBom godoc
// @Summary Create a new PlmEco
// @Description Create a new PlmEco in the system
// @Tags supply_chain-plm
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/plm [post]
// @Security BearerAuth
func CreateBomHandler(c echo.Context) error {
	var data PlmEco
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateBomService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllBom godoc
// @Summary Get all PlmEco
// @Description Retrieve a list of all PlmEco
// @Tags supply_chain-plm
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/plm [get]
// @Security BearerAuth
func GetAllBomHandler(c echo.Context) error {
	data, err := GetAllBomService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetBomByID godoc
// @Summary Get a PlmEco by ID
// @Description Retrieve a specific PlmEco by its ID
// @Tags supply_chain-plm
// @Produce json
// @Param id path int true "PlmEco ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/plm/{id} [get]
// @Security BearerAuth
func GetBomByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBomByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateBom godoc
// @Summary Update a PlmEco
// @Description Update an existing PlmEco
// @Tags supply_chain-plm
// @Accept json
// @Produce json
// @Param id path int true "PlmEco ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/plm/{id} [put]
// @Security BearerAuth
func UpdateBomHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBomByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateBomService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteBom godoc
// @Summary Delete a PlmEco
// @Description Delete a PlmEco by ID
// @Tags supply_chain-plm
// @Produce json
// @Param id path int true "PlmEco ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/plm/{id} [delete]
// @Security BearerAuth
func DeleteBomHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteBomService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
