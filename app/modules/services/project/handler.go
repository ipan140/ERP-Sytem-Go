package project

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateProject godoc
// @Summary Create a new Project
// @Description Create a new Project in the system
// @Tags services-project
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/services/project [post]
// @Security BearerAuth
func CreateProjectHandler(c echo.Context) error {
	var data Project
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateProjectService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllProject godoc
// @Summary Get all Project
// @Description Retrieve a list of all Project
// @Tags services-project
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project [get]
// @Security BearerAuth
func GetAllProjectHandler(c echo.Context) error {
	data, err := GetAllProjectService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetProjectByID godoc
// @Summary Get a Project by ID
// @Description Retrieve a specific Project by its ID
// @Tags services-project
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/{id} [get]
// @Security BearerAuth
func GetProjectByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProjectByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateProject godoc
// @Summary Update a Project
// @Description Update an existing Project
// @Tags services-project
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/{id} [put]
// @Security BearerAuth
func UpdateProjectHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProjectByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateProjectService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteProject godoc
// @Summary Delete a Project
// @Description Delete a Project by ID
// @Tags services-project
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/{id} [delete]
// @Security BearerAuth
func DeleteProjectHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProjectService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

func CreateProjectMilestoneHandler(c echo.Context) error { var data ProjectMilestone; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateProjectMilestoneService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllProjectMilestoneHandler(c echo.Context) error { data, err := GetAllProjectMilestoneService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetProjectMilestoneByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetProjectMilestoneByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateProjectMilestoneHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetProjectMilestoneByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateProjectMilestoneService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteProjectMilestoneHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteProjectMilestoneService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreateTaskDependencyHandler(c echo.Context) error { var data TaskDependency; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateTaskDependencyService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllTaskDependencyHandler(c echo.Context) error { data, err := GetAllTaskDependencyService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetTaskDependencyByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetTaskDependencyByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateTaskDependencyHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetTaskDependencyByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateTaskDependencyService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteTaskDependencyHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteTaskDependencyService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreateResourceForecastHandler(c echo.Context) error { var data ResourceForecast; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateResourceForecastService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllResourceForecastHandler(c echo.Context) error { data, err := GetAllResourceForecastService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetResourceForecastByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetResourceForecastByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateResourceForecastHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetResourceForecastByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateResourceForecastService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteResourceForecastHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteResourceForecastService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

