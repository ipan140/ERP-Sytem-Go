package crm

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

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

// @Summary Create SalesTeam
// @Description Create a new SalesTeam
// @Tags sales-crm
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/crm/salesteam [post]
// @Security BearerAuth
func CreateSalesTeamHandler(c echo.Context) error {
	var data SalesTeam
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSalesTeamService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all SalesTeam
// @Description Retrieve a list of all SalesTeam
// @Tags sales-crm
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/salesteam [get]
// @Security BearerAuth
func GetAllSalesTeamHandler(c echo.Context) error {
	data, err := GetAllSalesTeamService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSalesTeamByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesTeamByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update SalesTeam
// @Description Update an existing SalesTeam
// @Tags sales-crm
// @Accept json
// @Produce json
// @Param id path int true "SalesTeam ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/salesteam/{id} [put]
// @Security BearerAuth
func UpdateSalesTeamHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesTeamByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSalesTeamService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete SalesTeam
// @Description Delete SalesTeam by ID
// @Tags sales-crm
// @Produce json
// @Param id path int true "SalesTeam ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/salesteam/{id} [delete]
// @Security BearerAuth
func DeleteSalesTeamHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSalesTeamService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Stage
// @Description Create a new Stage
// @Tags sales-crm
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/crm/stage [post]
// @Security BearerAuth
func CreateStageHandler(c echo.Context) error {
	var data Stage
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateStageService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Stage
// @Description Retrieve a list of all Stage
// @Tags sales-crm
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/stage [get]
// @Security BearerAuth
func GetAllStageHandler(c echo.Context) error {
	data, err := GetAllStageService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetStageByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStageByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Stage
// @Description Update an existing Stage
// @Tags sales-crm
// @Accept json
// @Produce json
// @Param id path int true "Stage ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/stage/{id} [put]
// @Security BearerAuth
func UpdateStageHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetStageByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateStageService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Stage
// @Description Delete Stage by ID
// @Tags sales-crm
// @Produce json
// @Param id path int true "Stage ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/stage/{id} [delete]
// @Security BearerAuth
func DeleteStageHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteStageService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Activity
// @Description Create a new Activity
// @Tags sales-crm
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/crm/activity [post]
// @Security BearerAuth
func CreateActivityHandler(c echo.Context) error {
	var data Activity
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateActivityService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Activity
// @Description Retrieve a list of all Activity
// @Tags sales-crm
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/activity [get]
// @Security BearerAuth
func GetAllActivityHandler(c echo.Context) error {
	data, err := GetAllActivityService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetActivityByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetActivityByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Activity
// @Description Update an existing Activity
// @Tags sales-crm
// @Accept json
// @Produce json
// @Param id path int true "Activity ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/activity/{id} [put]
// @Security BearerAuth
func UpdateActivityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetActivityByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateActivityService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Activity
// @Description Delete Activity by ID
// @Tags sales-crm
// @Produce json
// @Param id path int true "Activity ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/activity/{id} [delete]
// @Security BearerAuth
func DeleteActivityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteActivityService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create SalesCommission
// @Description Create a new SalesCommission
// @Tags sales-crm
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/crm/salescommission [post]
// @Security BearerAuth
func CreateSalesCommissionHandler(c echo.Context) error {
	var data SalesCommission
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSalesCommissionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all SalesCommission
// @Description Retrieve a list of all SalesCommission
// @Tags sales-crm
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/salescommission [get]
// @Security BearerAuth
func GetAllSalesCommissionHandler(c echo.Context) error {
	data, err := GetAllSalesCommissionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSalesCommissionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesCommissionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update SalesCommission
// @Description Update an existing SalesCommission
// @Tags sales-crm
// @Accept json
// @Produce json
// @Param id path int true "SalesCommission ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/salescommission/{id} [put]
// @Security BearerAuth
func UpdateSalesCommissionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalesCommissionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSalesCommissionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete SalesCommission
// @Description Delete SalesCommission by ID
// @Tags sales-crm
// @Produce json
// @Param id path int true "SalesCommission ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/crm/salescommission/{id} [delete]
// @Security BearerAuth
func DeleteSalesCommissionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSalesCommissionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
