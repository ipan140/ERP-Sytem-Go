package quality

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateQualityCheck godoc
// @Summary Create a new QualityCheck
// @Description Create a new QualityCheck in the system
// @Tags supply_chain-quality
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/quality [post]
// @Security BearerAuth
func CreateQualityCheckHandler(c echo.Context) error {
	var data QualityCheck
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateQualityCheckService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllQualityCheck godoc
// @Summary Get all QualityCheck
// @Description Retrieve a list of all QualityCheck
// @Tags supply_chain-quality
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/quality [get]
// @Security BearerAuth
func GetAllQualityCheckHandler(c echo.Context) error {
	data, err := GetAllQualityCheckService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetQualityCheckByID godoc
// @Summary Get a QualityCheck by ID
// @Description Retrieve a specific QualityCheck by its ID
// @Tags supply_chain-quality
// @Produce json
// @Param id path int true "QualityCheck ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/quality/{id} [get]
// @Security BearerAuth
func GetQualityCheckByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQualityCheckByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateQualityCheck godoc
// @Summary Update a QualityCheck
// @Description Update an existing QualityCheck
// @Tags supply_chain-quality
// @Accept json
// @Produce json
// @Param id path int true "QualityCheck ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/quality/{id} [put]
// @Security BearerAuth
func UpdateQualityCheckHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQualityCheckByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateQualityCheckService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteQualityCheck godoc
// @Summary Delete a QualityCheck
// @Description Delete a QualityCheck by ID
// @Tags supply_chain-quality
// @Produce json
// @Param id path int true "QualityCheck ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/quality/{id} [delete]
// @Security BearerAuth
func DeleteQualityCheckHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteQualityCheckService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
