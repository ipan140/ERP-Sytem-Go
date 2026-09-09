package repairs

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateRepairOrder godoc
// @Summary Create a new RepairOrder
// @Description Create a new RepairOrder in the system
// @Tags services-repairs
// @Accept json
// @Produce json
// @Success 201 {object} RepairOrder
// @Param request body RepairOrder true "Payload"
// @Router /api/services/repairs [post]
// @Security BearerAuth
func CreateRepairOrderHandler(c echo.Context) error {
	var data RepairOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateRepairOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllRepairOrder godoc
// @Summary Get all RepairOrder
// @Description Retrieve a list of all RepairOrder
// @Tags services-repairs
// @Produce json
// @Success 200 {object} []RepairOrder
// @Router /api/services/repairs [get]
// @Security BearerAuth
func GetAllRepairOrderHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllRepairOrderService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")
	warranty := c.QueryParam("warranty_status")
	companyID, _ := strconv.Atoi(c.QueryParam("company_id"))
	technicianID, _ := strconv.Atoi(c.QueryParam("technician_id"))

	userRole, _ := c.Get("role").(string)
	if c.QueryParam("my_only") == "true" && (userRole == "staff" || userRole == "technician") {
		if currentTechID, _ := strconv.Atoi(c.QueryParam("my_technician_id")); currentTechID > 0 {
			technicianID = currentTechID
		}
	}

	data, total, err := GetPaginatedRepairOrderService(offset, limit, search, state, warranty, uint(companyID), uint(technicianID))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve paginated data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetRepairOrderByID godoc
// @Summary Get a RepairOrder by ID
// @Description Retrieve a specific RepairOrder by its ID
// @Tags services-repairs
// @Produce json
// @Param id path int true "RepairOrder ID"
// @Success 200 {object} RepairOrder
// @Router /api/services/repairs/{id} [get]
// @Security BearerAuth
func GetRepairOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRepairOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateRepairOrder godoc
// @Summary Update a RepairOrder
// @Description Update an existing RepairOrder
// @Tags services-repairs
// @Accept json
// @Produce json
// @Param id path int true "RepairOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/repairs/{id} [put]
// @Security BearerAuth
func UpdateRepairOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRepairOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateRepairOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteRepairOrder godoc
// @Summary Delete a RepairOrder
// @Description Delete a RepairOrder by ID
// @Tags services-repairs
// @Produce json
// @Param id path int true "RepairOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/repairs/{id} [delete]
// @Security BearerAuth
func DeleteRepairOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteRepairOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// Fase 2: Quality Control Gate Handler
func PassQCHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	inspectorID := uint(1)
	if uid, ok := c.Get("user_id").(uint); ok && uid > 0 {
		inspectorID = uid
	} else if uidFloat, ok := c.Get("user_id").(float64); ok && uidFloat > 0 {
		inspectorID = uint(uidFloat)
	}

	var req struct {
		Notes string `json:"notes"`
	}
	_ = c.Bind(&req)

	if err := PassQCRepairOrderService(uint(id), inspectorID, req.Notes); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memproses validasi QC", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Unit reparasi telah lulus inspeksi QC dan siap diserahterimakan", nil)
}

// PublicTrackingHandler godoc
// @Summary Public repair tracking by RMA and Serial Number
// @Tags services-repairs
// @Router /api/public/services/tracking [get]
func PublicTrackingHandler(c echo.Context) error {
	rmaStr := c.QueryParam("rma")
	if rmaStr == "" {
		rmaStr = c.QueryParam("id")
	}
	// Support #RO-0001 format or raw number
	cleanRMA := ""
	for _, ch := range rmaStr {
		if ch >= '0' && ch <= '9' {
			cleanRMA += string(ch)
		}
	}

	id, err := strconv.Atoi(cleanRMA)
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Nomor RMA tidak valid. Gunakan format seperti #RO-0001 atau angka ID", "")
	}

	sn := c.QueryParam("sn")
	if sn == "" {
		sn = c.QueryParam("serial_number")
	}

	order, err := GetRepairByRMAAndSerialService(uint(id), sn)
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data perbaikan tidak ditemukan. Periksa kembali No. RMA dan Nomor Seri unit Anda.", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Status reparasi berhasil dimuat", order)
}

// PublicApproveEstimateHandler godoc
// @Summary Customer approves repair cost estimate from public portal
// @Tags services-repairs
// @Router /api/public/services/tracking/approve-estimate [post]
func PublicApproveEstimateHandler(c echo.Context) error {
	var req struct {
		RMA          string `json:"rma"`
		ID           uint   `json:"id"`
		SerialNumber string `json:"serial_number"`
		Note         string `json:"note"`
	}
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Payload request tidak valid", err.Error())
	}

	targetID := req.ID
	if targetID == 0 && req.RMA != "" {
		cleanRMA := ""
		for _, ch := range req.RMA {
			if ch >= '0' && ch <= '9' {
				cleanRMA += string(ch)
			}
		}
		id, _ := strconv.Atoi(cleanRMA)
		targetID = uint(id)
	}

	if targetID == 0 {
		return utils.SendError(c, http.StatusBadRequest, "Nomor RMA / ID perbaikan wajib disertakan", "")
	}

	order, err := ApproveRepairEstimateService(targetID, req.Note)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyetujui estimasi biaya", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Estimasi biaya berhasil disetujui. Teknisi kami akan segera memproses perbaikan.", order)
}



