package project

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateProject godoc
// @Summary Create a new Project
// @Description Create a new Project in the system
// @Tags services-project
// @Accept json
// @Produce json
// @Success 201 {object} Project
// @Param request body Project true "Payload"
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
// @Success 200 {object} []Project
// @Router /api/services/project [get]
// @Security BearerAuth
func GetAllProjectHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllProjectService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedProjectService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve paginated data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetProjectByID godoc
// @Summary Get a Project by ID
// @Description Retrieve a specific Project by its ID
// @Tags services-project
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} Project
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

// @Summary Create ProjectMilestone
// @Description Create a new ProjectMilestone
// @Tags services-project
// @Accept json
// @Produce json
// @Success 201 {object} ProjectMilestone
// @Param request body ProjectMilestone true "Payload"
// @Router /api/services/project/projectmilestone [post]
// @Security BearerAuth
func CreateProjectMilestoneHandler(c echo.Context) error {
	var data ProjectMilestone
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateProjectMilestoneService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all ProjectMilestone
// @Description Retrieve a list of all ProjectMilestone
// @Tags services-project
// @Produce json
// @Success 200 {object} ProjectMilestone
// @Router /api/services/project/projectmilestone [get]
// @Security BearerAuth
func GetAllProjectMilestoneHandler(c echo.Context) error {
	data, err := GetAllProjectMilestoneService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetProjectMilestoneByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProjectMilestoneByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update ProjectMilestone
// @Description Update an existing ProjectMilestone
// @Tags services-project
// @Accept json
// @Produce json
// @Param id path int true "ProjectMilestone ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/projectmilestone/{id} [put]
// @Security BearerAuth
func UpdateProjectMilestoneHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProjectMilestoneByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateProjectMilestoneService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete ProjectMilestone
// @Description Delete ProjectMilestone by ID
// @Tags services-project
// @Produce json
// @Param id path int true "ProjectMilestone ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/projectmilestone/{id} [delete]
// @Security BearerAuth
func DeleteProjectMilestoneHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProjectMilestoneService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create TaskDependency
// @Description Create a new TaskDependency
// @Tags services-project
// @Accept json
// @Produce json
// @Success 201 {object} TaskDependency
// @Param request body TaskDependency true "Payload"
// @Router /api/services/project/taskdependency [post]
// @Security BearerAuth
func CreateTaskDependencyHandler(c echo.Context) error {
	var data TaskDependency
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateTaskDependencyService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all TaskDependency
// @Description Retrieve a list of all TaskDependency
// @Tags services-project
// @Produce json
// @Success 200 {object} TaskDependency
// @Router /api/services/project/taskdependency [get]
// @Security BearerAuth
func GetAllTaskDependencyHandler(c echo.Context) error {
	data, err := GetAllTaskDependencyService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetTaskDependencyByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTaskDependencyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update TaskDependency
// @Description Update an existing TaskDependency
// @Tags services-project
// @Accept json
// @Produce json
// @Param id path int true "TaskDependency ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/taskdependency/{id} [put]
// @Security BearerAuth
func UpdateTaskDependencyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTaskDependencyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateTaskDependencyService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete TaskDependency
// @Description Delete TaskDependency by ID
// @Tags services-project
// @Produce json
// @Param id path int true "TaskDependency ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/taskdependency/{id} [delete]
// @Security BearerAuth
func DeleteTaskDependencyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteTaskDependencyService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create ResourceForecast
// @Description Create a new ResourceForecast
// @Tags services-project
// @Accept json
// @Produce json
// @Success 201 {object} ResourceForecast
// @Param request body ResourceForecast true "Payload"
// @Router /api/services/project/resourceforecast [post]
// @Security BearerAuth
func CreateResourceForecastHandler(c echo.Context) error {
	var data ResourceForecast
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateResourceForecastService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all ResourceForecast
// @Description Retrieve a list of all ResourceForecast
// @Tags services-project
// @Produce json
// @Success 200 {object} ResourceForecast
// @Router /api/services/project/resourceforecast [get]
// @Security BearerAuth
func GetAllResourceForecastHandler(c echo.Context) error {
	data, err := GetAllResourceForecastService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetResourceForecastByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetResourceForecastByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update ResourceForecast
// @Description Update an existing ResourceForecast
// @Tags services-project
// @Accept json
// @Produce json
// @Param id path int true "ResourceForecast ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/resourceforecast/{id} [put]
// @Security BearerAuth
func UpdateResourceForecastHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetResourceForecastByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateResourceForecastService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete ResourceForecast
// @Description Delete ResourceForecast by ID
// @Tags services-project
// @Produce json
// @Param id path int true "ResourceForecast ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/resourceforecast/{id} [delete]
// @Security BearerAuth
func DeleteResourceForecastHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteResourceForecastService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create Task
// @Description Create a new Task
// @Tags services-project
// @Accept json
// @Produce json
// @Success 201 {object} Task
// @Param request body Task true "Payload"
// @Router /api/services/project/task [post]
// @Security BearerAuth
func CreateTaskHandler(c echo.Context) error {
	var data Task
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateTaskService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all Task
// @Description Retrieve a list of all Task
// @Tags services-project
// @Produce json
// @Success 200 {object} []Task
// @Router /api/services/project/task [get]
// @Security BearerAuth
func GetAllTaskHandler(c echo.Context) error {
	data, err := GetAllTaskService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

func GetTaskByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTaskByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update Task
// @Description Update an existing Task
// @Tags services-project
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/task/{id} [put]
// @Security BearerAuth
func UpdateTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTaskByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateTaskService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete Task
// @Description Delete Task by ID
// @Tags services-project
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/project/task/{id} [delete]
// @Security BearerAuth
func DeleteTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteTaskService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}



