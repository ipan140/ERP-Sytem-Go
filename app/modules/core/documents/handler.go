package documents

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateWorkspace godoc
// @Summary Create a new Workspace
// @Description Create a new Workspace in the system
// @Tags documents
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/documents [post]
// @Security BearerAuth
func CreateWorkspaceHandler(c echo.Context) error {
	var data Workspace
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c,  http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateWorkspaceService(&data); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllWorkspace godoc
// @Summary Get all Workspace
// @Description Retrieve a list of all Workspace
// @Tags documents
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/documents [get]
// @Security BearerAuth
func GetAllWorkspaceHandler(c echo.Context) error {
	data, err := GetAllWorkspaceService()
	if err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetWorkspaceByID godoc
// @Summary Get a Workspace by ID
// @Description Retrieve a specific Workspace by its ID
// @Tags documents
// @Produce json
// @Param id path int true "Workspace ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/documents/{id} [get]
// @Security BearerAuth
func GetWorkspaceByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkspaceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c,  http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateWorkspace godoc
// @Summary Update a Workspace
// @Description Update an existing Workspace
// @Tags documents
// @Accept json
// @Produce json
// @Param id path int true "Workspace ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/documents/{id} [put]
// @Security BearerAuth
func UpdateWorkspaceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkspaceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c,  http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c,  http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateWorkspaceService(data); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteWorkspace godoc
// @Summary Delete a Workspace
// @Description Delete a Workspace by ID
// @Tags documents
// @Produce json
// @Param id path int true "Workspace ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/documents/{id} [delete]
// @Security BearerAuth
func DeleteWorkspaceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWorkspaceService(uint(id)); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
