package sign

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateSignatureRequest godoc
// @Summary Create a new SignatureRequest
// @Description Create a new SignatureRequest in the system
// @Tags finance-sign
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/sign [post]
// @Security BearerAuth
func CreateSignatureRequestHandler(c echo.Context) error {
	var data SignatureRequest
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSignatureRequestService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSignatureRequest godoc
// @Summary Get all SignatureRequest
// @Description Retrieve a list of all SignatureRequest
// @Tags finance-sign
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/sign [get]
// @Security BearerAuth
func GetAllSignatureRequestHandler(c echo.Context) error {
	data, err := GetAllSignatureRequestService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSignatureRequestByID godoc
// @Summary Get a SignatureRequest by ID
// @Description Retrieve a specific SignatureRequest by its ID
// @Tags finance-sign
// @Produce json
// @Param id path int true "SignatureRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/sign/{id} [get]
// @Security BearerAuth
func GetSignatureRequestByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSignatureRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSignatureRequest godoc
// @Summary Update a SignatureRequest
// @Description Update an existing SignatureRequest
// @Tags finance-sign
// @Accept json
// @Produce json
// @Param id path int true "SignatureRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/sign/{id} [put]
// @Security BearerAuth
func UpdateSignatureRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSignatureRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSignatureRequestService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSignatureRequest godoc
// @Summary Delete a SignatureRequest
// @Description Delete a SignatureRequest by ID
// @Tags finance-sign
// @Produce json
// @Param id path int true "SignatureRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/sign/{id} [delete]
// @Security BearerAuth
func DeleteSignatureRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSignatureRequestService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
