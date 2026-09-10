package barcode

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ScanBarcodeHandler(c echo.Context) error {
	var req BarcodeScanRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	result, err := ScanBarcodeService(req.Barcode)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to scan barcode", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Barcode scanned successfully", result)
}

func GetBarcodeSummaryHandler(c echo.Context) error {
	summary, err := GetBarcodeSummaryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve barcode summary", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Barcode summary retrieved successfully", summary)
}

func GetAllBarcodeConfigHandler(c echo.Context) error {
	data, err := GetAllBarcodeConfigService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

func GetBarcodeConfigByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBarcodeConfigByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

func CreateBarcodeConfigHandler(c echo.Context) error {
	var data BarcodeNomenclature
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateBarcodeConfigService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

func UpdateBarcodeConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBarcodeConfigByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateBarcodeConfigService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

func DeleteBarcodeConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteBarcodeConfigService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
