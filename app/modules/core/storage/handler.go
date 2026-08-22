package storage

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAttachment godoc
// @Summary Create a new Attachment
// @Description Create a new Attachment in the system
// @Tags core-storage
// @Accept json
// @Produce json
// @Success 201 {object} Attachment
// @Param request body Attachment true "Payload"
// @Router /api/core/storage [post]
// @Security BearerAuth
func CreateAttachmentHandler(c echo.Context) error {
	var data Attachment
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAttachmentService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllAttachment godoc
// @Summary Get all Attachment
// @Description Retrieve a list of all Attachment
// @Tags core-storage
// @Produce json
// @Success 200 {object} []Attachment
// @Router /api/core/storage [get]
// @Security BearerAuth
func GetAllAttachmentHandler(c echo.Context) error {
	data, err := GetAllAttachmentService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetAttachmentByID godoc
// @Summary Get a Attachment by ID
// @Description Retrieve a specific Attachment by its ID
// @Tags core-storage
// @Produce json
// @Param id path int true "Attachment ID"
// @Success 200 {object} Attachment
// @Router /api/core/storage/{id} [get]
// @Security BearerAuth
func GetAttachmentByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAttachmentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateAttachment godoc
// @Summary Update a Attachment
// @Description Update an existing Attachment
// @Tags core-storage
// @Accept json
// @Produce json
// @Param id path int true "Attachment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/storage/{id} [put]
// @Security BearerAuth
func UpdateAttachmentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAttachmentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAttachmentService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteAttachment godoc
// @Summary Delete a Attachment
// @Description Delete a Attachment by ID
// @Tags core-storage
// @Produce json
// @Param id path int true "Attachment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/storage/{id} [delete]
// @Security BearerAuth
func DeleteAttachmentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAttachmentService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


