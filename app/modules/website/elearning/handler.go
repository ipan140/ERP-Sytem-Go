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

// @Summary Create Slide
// @Description Create a new Slide
// @Tags website-elearning
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/elearning/slide [post]
// @Security BearerAuth
func CreateSlideHandler(c echo.Context) error {
	var data Slide
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSlideService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all Slide
// @Description Retrieve a list of all Slide
// @Tags website-elearning
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/slide [get]
// @Security BearerAuth
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
// @Summary Update Slide
// @Description Update an existing Slide
// @Tags website-elearning
// @Accept json
// @Produce json
// @Param id path int true "Slide ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/slide/{id} [put]
// @Security BearerAuth
func UpdateSlideHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSlideByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSlideService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete Slide
// @Description Delete Slide by ID
// @Tags website-elearning
// @Produce json
// @Param id path int true "Slide ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/slide/{id} [delete]
// @Security BearerAuth
func DeleteSlideHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSlideService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Certification
// @Description Create a new Certification
// @Tags website-elearning
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/elearning/certification [post]
// @Security BearerAuth
func CreateCertificationHandler(c echo.Context) error { var data Certification; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateCertificationService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
// @Summary Get all Certification
// @Description Retrieve a list of all Certification
// @Tags website-elearning
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/certification [get]
// @Security BearerAuth
func GetAllCertificationHandler(c echo.Context) error { data, err := GetAllCertificationService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetCertificationByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetCertificationByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
// @Summary Update Certification
// @Description Update an existing Certification
// @Tags website-elearning
// @Accept json
// @Produce json
// @Param id path int true "Certification ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/certification/{id} [put]
// @Security BearerAuth
func UpdateCertificationHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetCertificationByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateCertificationService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
// @Summary Delete Certification
// @Description Delete Certification by ID
// @Tags website-elearning
// @Produce json
// @Param id path int true "Certification ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/elearning/certification/{id} [delete]
// @Security BearerAuth
func DeleteCertificationHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteCertificationService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

