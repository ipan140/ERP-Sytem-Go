package consolidation

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateConsolidationEntry godoc
// @Summary Create a new ConsolidationEntry
// @Description Create a new ConsolidationEntry in the system
// @Tags finance-consolidation
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/consolidation [post]
// @Security BearerAuth
func CreateConsolidationEntryHandler(c echo.Context) error {
	var data ConsolidationEntry
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateConsolidationEntryService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllConsolidationEntry godoc
// @Summary Get all ConsolidationEntry
// @Description Retrieve a list of all ConsolidationEntry
// @Tags finance-consolidation
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/consolidation [get]
// @Security BearerAuth
func GetAllConsolidationEntryHandler(c echo.Context) error {
	data, err := GetAllConsolidationEntryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetConsolidationEntryByID godoc
// @Summary Get a ConsolidationEntry by ID
// @Description Retrieve a specific ConsolidationEntry by its ID
// @Tags finance-consolidation
// @Produce json
// @Param id path int true "ConsolidationEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/consolidation/{id} [get]
// @Security BearerAuth
func GetConsolidationEntryByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetConsolidationEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateConsolidationEntry godoc
// @Summary Update a ConsolidationEntry
// @Description Update an existing ConsolidationEntry
// @Tags finance-consolidation
// @Accept json
// @Produce json
// @Param id path int true "ConsolidationEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/consolidation/{id} [put]
// @Security BearerAuth
func UpdateConsolidationEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetConsolidationEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateConsolidationEntryService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteConsolidationEntry godoc
// @Summary Delete a ConsolidationEntry
// @Description Delete a ConsolidationEntry by ID
// @Tags finance-consolidation
// @Produce json
// @Param id path int true "ConsolidationEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/consolidation/{id} [delete]
// @Security BearerAuth
func DeleteConsolidationEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteConsolidationEntryService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
