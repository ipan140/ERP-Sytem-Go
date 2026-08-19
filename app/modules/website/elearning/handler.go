package elearning

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateCourse godoc
// @Summary Create a new Course
// @Description Create a new Course in the system
// @Tags website-elearning
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/elearning [post]
// @Security BearerAuth
func CreateCourseHandler(c echo.Context) error {
	var data Course
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateCourseService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllCourse godoc
// @Summary Get all Course
// @Description Retrieve a list of all Course
// @Tags website-elearning
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning [get]
// @Security BearerAuth
func GetAllCourseHandler(c echo.Context) error {
	data, err := GetAllCourseService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetCourseByID godoc
// @Summary Get a Course by ID
// @Description Retrieve a specific Course by its ID
// @Tags website-elearning
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/{id} [get]
// @Security BearerAuth
func GetCourseByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCourseByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateCourse godoc
// @Summary Update a Course
// @Description Update an existing Course
// @Tags website-elearning
// @Accept json
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/{id} [put]
// @Security BearerAuth
func UpdateCourseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCourseByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateCourseService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteCourse godoc
// @Summary Delete a Course
// @Description Delete a Course by ID
// @Tags website-elearning
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/{id} [delete]
// @Security BearerAuth
func DeleteCourseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCourseService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

func CreateSlideHandler(c echo.Context) error {
	var data Slide
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSlideService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllSlideHandler(c echo.Context) error {
	data, err := GetAllSlideService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSlideByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSlideByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateSlideHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSlideByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSlideService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteSlideHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSlideService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
