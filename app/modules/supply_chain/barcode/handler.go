package barcode

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateBarcodeConfig godoc
// @Summary Create a new BarcodeNomenclature
// @Description Create a new BarcodeNomenclature in the system
// @Tags supply_chain-barcode
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/barcode [post]
// @Security BearerAuth
func CreateBarcodeConfigHandler(c echo.Context) error {
	var data BarcodeNomenclature
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateBarcodeConfigService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllBarcodeConfig godoc
// @Summary Get all BarcodeNomenclature
// @Description Retrieve a list of all BarcodeNomenclature
// @Tags supply_chain-barcode
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/barcode [get]
// @Security BearerAuth
func GetAllBarcodeConfigHandler(c echo.Context) error {
	data, err := GetAllBarcodeConfigService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetBarcodeConfigByID godoc
// @Summary Get a BarcodeNomenclature by ID
// @Description Retrieve a specific BarcodeNomenclature by its ID
// @Tags supply_chain-barcode
// @Produce json
// @Param id path int true "BarcodeNomenclature ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/barcode/{id} [get]
// @Security BearerAuth
func GetBarcodeConfigByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBarcodeConfigByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateBarcodeConfig godoc
// @Summary Update a BarcodeNomenclature
// @Description Update an existing BarcodeNomenclature
// @Tags supply_chain-barcode
// @Accept json
// @Produce json
// @Param id path int true "BarcodeNomenclature ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/barcode/{id} [put]
// @Security BearerAuth
func UpdateBarcodeConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBarcodeConfigByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateBarcodeConfigService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteBarcodeConfig godoc
// @Summary Delete a BarcodeNomenclature
// @Description Delete a BarcodeNomenclature by ID
// @Tags supply_chain-barcode
// @Produce json
// @Param id path int true "BarcodeNomenclature ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/barcode/{id} [delete]
// @Security BearerAuth
func DeleteBarcodeConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteBarcodeConfigService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
