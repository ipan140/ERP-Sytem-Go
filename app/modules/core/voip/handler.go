package voip

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateCallRecord godoc
// @Summary Create a new CallRecord
// @Description Create a new CallRecord in the system
// @Tags voip
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/voip [post]
// @Security BearerAuth
func CreateCallRecordHandler(c echo.Context) error {
	var data CallRecord
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c,  http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateCallRecordService(&data); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllCallRecord godoc
// @Summary Get all CallRecord
// @Description Retrieve a list of all CallRecord
// @Tags voip
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/voip [get]
// @Security BearerAuth
func GetAllCallRecordHandler(c echo.Context) error {
	data, err := GetAllCallRecordService()
	if err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetCallRecordByID godoc
// @Summary Get a CallRecord by ID
// @Description Retrieve a specific CallRecord by its ID
// @Tags voip
// @Produce json
// @Param id path int true "CallRecord ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/voip/{id} [get]
// @Security BearerAuth
func GetCallRecordByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCallRecordByIDService(uint(id))
	if err != nil {
		return utils.SendError(c,  http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateCallRecord godoc
// @Summary Update a CallRecord
// @Description Update an existing CallRecord
// @Tags voip
// @Accept json
// @Produce json
// @Param id path int true "CallRecord ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/voip/{id} [put]
// @Security BearerAuth
func UpdateCallRecordHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCallRecordByIDService(uint(id))
	if err != nil {
		return utils.SendError(c,  http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c,  http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateCallRecordService(data); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteCallRecord godoc
// @Summary Delete a CallRecord
// @Description Delete a CallRecord by ID
// @Tags voip
// @Produce json
// @Param id path int true "CallRecord ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/voip/{id} [delete]
// @Security BearerAuth
func DeleteCallRecordHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCallRecordService(uint(id)); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
