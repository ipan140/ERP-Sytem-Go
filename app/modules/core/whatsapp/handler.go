package whatsapp

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateWaTemplate godoc
// @Summary Create a new WaTemplate
// @Description Create a new WaTemplate in the system
// @Tags whatsapp
// @Accept json
// @Produce json
// @Success 201 {object} WaTemplate
// @Param request body WaTemplate true "Payload"
// @Router /api/whatsapp [post]
// @Security BearerAuth
func CreateWaTemplateHandler(c echo.Context) error {
	var data WaTemplate
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateWaTemplateService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllWaTemplate godoc
// @Summary Get all WaTemplate
// @Description Retrieve a list of all WaTemplate
// @Tags whatsapp
// @Produce json
// @Success 200 {object} []WaTemplate
// @Router /api/whatsapp [get]
// @Security BearerAuth
func GetAllWaTemplateHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllWaTemplateService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedWaTemplateService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetWaTemplateByID godoc
// @Summary Get a WaTemplate by ID
// @Description Retrieve a specific WaTemplate by its ID
// @Tags whatsapp
// @Produce json
// @Param id path int true "WaTemplate ID"
// @Success 200 {object} WaTemplate
// @Router /api/whatsapp/{id} [get]
// @Security BearerAuth
func GetWaTemplateByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWaTemplateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateWaTemplate godoc
// @Summary Update a WaTemplate
// @Description Update an existing WaTemplate
// @Tags whatsapp
// @Accept json
// @Produce json
// @Param id path int true "WaTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/whatsapp/{id} [put]
// @Security BearerAuth
func UpdateWaTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWaTemplateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateWaTemplateService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteWaTemplate godoc
// @Summary Delete a WaTemplate
// @Description Delete a WaTemplate by ID
// @Tags whatsapp
// @Produce json
// @Param id path int true "WaTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/whatsapp/{id} [delete]
// @Security BearerAuth
func DeleteWaTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWaTemplateService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


