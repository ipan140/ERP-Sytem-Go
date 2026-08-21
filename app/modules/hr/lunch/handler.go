package lunch

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateLunchOrder godoc
// @Summary Create a new LunchOrder
// @Description Create a new LunchOrder in the system
// @Tags hr-lunch
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/lunch [post]
// @Security BearerAuth
func CreateLunchOrderHandler(c echo.Context) error {
	var data LunchOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateLunchOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllLunchOrder godoc
// @Summary Get all LunchOrder
// @Description Retrieve a list of all LunchOrder
// @Tags hr-lunch
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/lunch [get]
// @Security BearerAuth
func GetAllLunchOrderHandler(c echo.Context) error {
	data, err := GetAllLunchOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetLunchOrderByID godoc
// @Summary Get a LunchOrder by ID
// @Description Retrieve a specific LunchOrder by its ID
// @Tags hr-lunch
// @Produce json
// @Param id path int true "LunchOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/lunch/{id} [get]
// @Security BearerAuth
func GetLunchOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLunchOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateLunchOrder godoc
// @Summary Update a LunchOrder
// @Description Update an existing LunchOrder
// @Tags hr-lunch
// @Accept json
// @Produce json
// @Param id path int true "LunchOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/lunch/{id} [put]
// @Security BearerAuth
func UpdateLunchOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLunchOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateLunchOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteLunchOrder godoc
// @Summary Delete a LunchOrder
// @Description Delete a LunchOrder by ID
// @Tags hr-lunch
// @Produce json
// @Param id path int true "LunchOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/lunch/{id} [delete]
// @Security BearerAuth
func DeleteLunchOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLunchOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create LunchCashmove
// @Description Create a new LunchCashmove
// @Tags hr-lunch
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/lunch/lunchcashmove [post]
// @Security BearerAuth
func CreateLunchCashmoveHandler(c echo.Context) error {
	var data LunchCashmove
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateLunchCashmoveService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all LunchCashmove
// @Description Retrieve a list of all LunchCashmove
// @Tags hr-lunch
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/lunch/lunchcashmove [get]
// @Security BearerAuth
func GetAllLunchCashmoveHandler(c echo.Context) error {
	data, err := GetAllLunchCashmoveService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetLunchCashmoveByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLunchCashmoveByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update LunchCashmove
// @Description Update an existing LunchCashmove
// @Tags hr-lunch
// @Accept json
// @Produce json
// @Param id path int true "LunchCashmove ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/lunch/lunchcashmove/{id} [put]
// @Security BearerAuth
func UpdateLunchCashmoveHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLunchCashmoveByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateLunchCashmoveService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete LunchCashmove
// @Description Delete LunchCashmove by ID
// @Tags hr-lunch
// @Produce json
// @Param id path int true "LunchCashmove ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/lunch/lunchcashmove/{id} [delete]
// @Security BearerAuth
func DeleteLunchCashmoveHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLunchCashmoveService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
