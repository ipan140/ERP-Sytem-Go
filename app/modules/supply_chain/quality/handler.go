package quality

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetQualitySummaryHandler godoc
// @Summary Get Quality summary KPI
// @Description Retrieve executive KPI summary for quality control
// @Tags supply_chain-quality
// @Produce json
// @Success 200 {object} QualitySummary
// @Router /api/supply_chain/quality/summary [get]
// @Security BearerAuth
func GetQualitySummaryHandler(c echo.Context) error {
	summary, err := GetQualitySummaryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil ringkasan mutu", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Ringkasan kendali mutu berhasil diambil", summary)
}

// CreateQualityCheck godoc
// @Summary Create a new QualityCheck
// @Description Create a new QualityCheck with auto sequence
// @Tags supply_chain-quality
// @Accept json
// @Produce json
// @Success 201 {object} QualityCheck
// @Param request body CreateQualityCheckRequest true "Payload"
// @Router /api/supply_chain/quality [post]
// @Security BearerAuth
func CreateQualityCheckHandler(c echo.Context) error {
	var req CreateQualityCheckRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if req.ProductID == 0 {
		return utils.SendError(c, http.StatusBadRequest, "Product ID wajib diisi", "Validation failed")
	}
	data, err := CreateQualityCheckWithSequenceService(req)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat uji kualitas", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Uji kualitas berhasil dibuat", data)
}

// GetAllQualityCheck godoc
// @Summary Get all QualityCheck (paginated)
// @Description Retrieve a list of QualityCheck with pagination and filters
// @Tags supply_chain-quality
// @Produce json
// @Success 200 {object} []QualityCheck
// @Router /api/supply_chain/quality [get]
// @Security BearerAuth
func GetAllQualityCheckHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllQualityCheckService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data uji mutu", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data uji mutu berhasil diambil", data)
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	search := c.QueryParam("search")
	result := c.QueryParam("result")
	pID, _ := strconv.Atoi(c.QueryParam("product_id"))

	data, total, err := GetPaginatedQualityChecksService(offset, limit, search, result, uint(pID))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data uji mutu paginasi", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data uji mutu berhasil diambil", data, meta)
}

func ProcessQualityCheckHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req ProcessQCRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	var inspectorID *uint
	if uID, ok := c.Get("user_id").(uint); ok && uID > 0 {
		inspectorID = &uID
	}

	data, err := ProcessQualityCheckService(uint(id), req, inspectorID)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, err.Error(), err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Hasil inspeksi kendali mutu berhasil disimpan", data)
}

func GetQualityPointsHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllQualityPointsService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil titik uji", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data titik uji berhasil diambil", data)
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit
	search := c.QueryParam("search")

	data, total, err := GetPaginatedQualityPointsService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data titik uji", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data titik uji berhasil diambil", data, meta)
}

func CreateQualityPointHandler(c echo.Context) error {
	var data QualityPoint
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if data.ProductID == 0 || data.Name == "" {
		return utils.SendError(c, http.StatusBadRequest, "Nama uji dan Product ID wajib diisi", "Validation failed")
	}
	if err := CreateQualityPointService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat titik uji mutu", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Titik uji mutu berhasil dibuat", data)
}

func DeleteQualityPointHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteQualityPointService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus titik uji mutu", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Titik uji mutu berhasil dihapus", nil)
}

// GetQualityCheckByID godoc
// @Summary Get a QualityCheck by ID
// @Description Retrieve a specific QualityCheck by its ID
// @Tags supply_chain-quality
// @Produce json
// @Param id path int true "QualityCheck ID"
// @Success 200 {object} QualityCheck
// @Router /api/supply_chain/quality/{id} [get]
// @Security BearerAuth
func GetQualityCheckByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQualityCheckByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateQualityCheck godoc
// @Summary Update a QualityCheck
// @Description Update an existing QualityCheck
// @Tags supply_chain-quality
// @Accept json
// @Produce json
// @Param id path int true "QualityCheck ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/quality/{id} [put]
// @Security BearerAuth
func UpdateQualityCheckHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQualityCheckByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateQualityCheckService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteQualityCheck godoc
// @Summary Delete a QualityCheck
// @Description Delete a QualityCheck by ID
// @Tags supply_chain-quality
// @Produce json
// @Param id path int true "QualityCheck ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/quality/{id} [delete]
// @Security BearerAuth
func DeleteQualityCheckHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteQualityCheckService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


