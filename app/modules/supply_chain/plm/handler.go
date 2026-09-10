package plm

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPlmSummaryHandler(c echo.Context) error {
	summary, err := GetPlmSummaryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve PLM summary", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "PLM summary retrieved successfully", summary)
}

func GetAllEcoHandler(c echo.Context) error {
	page, limit, _, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")

	list, total, err := GetPaginatedEcosService(page, limit, search, state)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve ECO records", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "ECO records retrieved successfully", list, meta)
}

func CreateEcoHandler(c echo.Context) error {
	var req CreateEcoRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if req.Name == "" || req.ProductID == 0 {
		return utils.SendError(c, http.StatusBadRequest, "ECO title and product are required", "")
	}

	eco, err := CreateEcoWithSequenceService(&req)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create ECO", err.Error())
	}

	return utils.SendSuccess(c, http.StatusCreated, "ECO created successfully", eco)
}

func UpdateEcoStateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateEcoStateRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if req.State != "draft" && req.State != "progress" && req.State != "approved" && req.State != "done" && req.State != "cancel" {
		return utils.SendError(c, http.StatusBadRequest, "Invalid ECO state", "")
	}

	var approverID *uint
	if userIDVal := c.Get("user_id"); userIDVal != nil {
		if uid, ok := userIDVal.(uint); ok {
			approverID = &uid
		}
	}

	eco, err := UpdateEcoStateService(uint(id), req.State, approverID)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update ECO state", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "ECO state updated successfully", eco)
}

func GetAllEcoTypesHandler(c echo.Context) error {
	types, err := GetAllEcoTypesService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve ECO types", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "ECO types retrieved successfully", types)
}

func GetEcoByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	eco, err := GetEcoByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "ECO record not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "ECO record retrieved successfully", eco)
}

func DeleteEcoHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEcoService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete ECO", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "ECO deleted successfully", nil)
}
