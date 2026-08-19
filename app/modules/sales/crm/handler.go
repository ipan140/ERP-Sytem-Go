package crm

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateLead godoc
// @Summary Create a new Lead
// @Description Create a new Lead in the system
// @Tags sales-crm
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/crm [post]
// @Security BearerAuth
func CreateLeadHandler(c echo.Context) error {
	var data Lead
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateLeadService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllLead godoc
// @Summary Get all Lead
// @Description Retrieve a list of all Lead
// @Tags sales-crm
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm [get]
// @Security BearerAuth
func GetAllLeadHandler(c echo.Context) error {
	data, err := GetAllLeadService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetLeadByID godoc
// @Summary Get a Lead by ID
// @Description Retrieve a specific Lead by its ID
// @Tags sales-crm
// @Produce json
// @Param id path int true "Lead ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/{id} [get]
// @Security BearerAuth
func GetLeadByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeadByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateLead godoc
// @Summary Update a Lead
// @Description Update an existing Lead
// @Tags sales-crm
// @Accept json
// @Produce json
// @Param id path int true "Lead ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/{id} [put]
// @Security BearerAuth
func UpdateLeadHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeadByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateLeadService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteLead godoc
// @Summary Delete a Lead
// @Description Delete a Lead by ID
// @Tags sales-crm
// @Produce json
// @Param id path int true "Lead ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/{id} [delete]
// @Security BearerAuth
func DeleteLeadHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLeadService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

func CreateSalesTeamHandler(c echo.Context) error {
	var data SalesTeam
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSalesTeamService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllSalesTeamHandler(c echo.Context) error {
	data, err := GetAllSalesTeamService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSalesTeamByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesTeamByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateSalesTeamHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesTeamByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSalesTeamService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteSalesTeamHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSalesTeamService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateStageHandler(c echo.Context) error {
	var data Stage
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateStageService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllStageHandler(c echo.Context) error {
	data, err := GetAllStageService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStageByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStageByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateStageHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStageByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateStageService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteStageHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStageService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateActivityHandler(c echo.Context) error {
	var data Activity
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateActivityService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllActivityHandler(c echo.Context) error {
	data, err := GetAllActivityService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetActivityByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetActivityByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateActivityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetActivityByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateActivityService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteActivityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteActivityService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateSalesCommissionHandler(c echo.Context) error {
	var data SalesCommission
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSalesCommissionService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllSalesCommissionHandler(c echo.Context) error {
	data, err := GetAllSalesCommissionService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSalesCommissionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesCommissionByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateSalesCommissionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesCommissionByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSalesCommissionService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteSalesCommissionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSalesCommissionService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
