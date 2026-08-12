package website_builder

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreatePage godoc
// @Summary Create a new Page
// @Description Create a new Page in the system
// @Tags website-website_builder
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/website_builder [post]
// @Security BearerAuth
func CreatePageHandler(c echo.Context) error {
	var data Page
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreatePageService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllPage godoc
// @Summary Get all Page
// @Description Retrieve a list of all Page
// @Tags website-website_builder
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/website_builder [get]
// @Security BearerAuth
func GetAllPageHandler(c echo.Context) error {
	data, err := GetAllPageService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetPageByID godoc
// @Summary Get a Page by ID
// @Description Retrieve a specific Page by its ID
// @Tags website-website_builder
// @Produce json
// @Param id path int true "Page ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/website_builder/{id} [get]
// @Security BearerAuth
func GetPageByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPageByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdatePage godoc
// @Summary Update a Page
// @Description Update an existing Page
// @Tags website-website_builder
// @Accept json
// @Produce json
// @Param id path int true "Page ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/website_builder/{id} [put]
// @Security BearerAuth
func UpdatePageHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPageByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdatePageService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeletePage godoc
// @Summary Delete a Page
// @Description Delete a Page by ID
// @Tags website-website_builder
// @Produce json
// @Param id path int true "Page ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/website_builder/{id} [delete]
// @Security BearerAuth
func DeletePageHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePageService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
