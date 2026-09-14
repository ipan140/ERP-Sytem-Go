package documents

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateFinanceDocument godoc
// @Summary Create a new FinanceDocument
// @Description Create a new FinanceDocument in the system
// @Tags finance-documents
// @Accept json
// @Produce json
// @Success 201 {object} FinanceDocument
// @Param request body FinanceDocument true "Payload"
// @Router /api/finance/documents [post]
// @Security BearerAuth
func CreateFinanceDocumentHandler(c echo.Context) error {
	var data FinanceDocument
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateFinanceDocumentService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllFinanceDocument godoc
// @Summary Get all FinanceDocument
// @Description Retrieve a list of all FinanceDocument
// @Tags finance-documents
// @Produce json
// @Success 200 {object} []FinanceDocument
// @Router /api/finance/documents [get]
// @Security BearerAuth
func GetAllFinanceDocumentHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllFinanceDocumentService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedFinanceDocumentService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetFinanceDocumentByID godoc
// @Summary Get a FinanceDocument by ID
// @Description Retrieve a specific FinanceDocument by its ID
// @Tags finance-documents
// @Produce json
// @Param id path int true "FinanceDocument ID"
// @Success 200 {object} FinanceDocument
// @Router /api/finance/documents/{id} [get]
// @Security BearerAuth
func GetFinanceDocumentByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFinanceDocumentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateFinanceDocument godoc
// @Summary Update a FinanceDocument
// @Description Update an existing FinanceDocument
// @Tags finance-documents
// @Accept json
// @Produce json
// @Param id path int true "FinanceDocument ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/documents/{id} [put]
// @Security BearerAuth
func UpdateFinanceDocumentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFinanceDocumentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateFinanceDocumentService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteFinanceDocument godoc
// @Summary Delete a FinanceDocument
// @Description Delete a FinanceDocument by ID
// @Tags finance-documents
// @Produce json
// @Param id path int true "FinanceDocument ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/documents/{id} [delete]
// @Security BearerAuth
func DeleteFinanceDocumentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteFinanceDocumentService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


