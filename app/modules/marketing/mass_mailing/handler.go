package mass_mailing

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// CreateMailingCampaign godoc
// @Summary Create a new MailingCampaign
// @Description Create a new MailingCampaign in the system
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Success 201 {object} MailingCampaign
// @Param request body MailingCampaign true "Payload"
// @Router /api/marketing/mass_mailing [post]
// @Security BearerAuth
func CreateMailingCampaignHandler(c echo.Context) error {
	var data MailingCampaign
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateMailingCampaignService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllMailingCampaign godoc
// @Summary Get all MailingCampaign
// @Description Retrieve a list of all MailingCampaign
// @Tags marketing-mass_mailing
// @Produce json
// @Success 200 {object} []MailingCampaign
// @Router /api/marketing/mass_mailing [get]
// @Security BearerAuth
func GetAllMailingCampaignHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllMailingCampaignService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedMailingCampaignService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetMailingCampaignByID godoc
// @Summary Get a MailingCampaign by ID
// @Description Retrieve a specific MailingCampaign by its ID
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} MailingCampaign
// @Router /api/marketing/mass_mailing/{id} [get]
// @Security BearerAuth
func GetMailingCampaignByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMailingCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateMailingCampaign godoc
// @Summary Update a MailingCampaign
// @Description Update an existing MailingCampaign
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id} [put]
// @Security BearerAuth
func UpdateMailingCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMailingCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if data.Status == "Sent" && data.ApprovalStatus != "Approved" {
		return utils.SendError(c, http.StatusForbidden, "Persetujuan Diperlukan", "Kampanye harus disetujui (Approved) oleh Manajer terlebih dahulu sebelum dapat dieksekusi/dikirim.")
	}
	if err := UpdateMailingCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// RequestApprovalMailingCampaign godoc
// @Summary Request Approval for a MailingCampaign
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id}/request-approval [put]
// @Security BearerAuth
func RequestApprovalMailingCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMailingCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	data.ApprovalStatus = "Waiting Approval"
	if err := UpdateMailingCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update approval status", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Kampanye berhasil diajukan untuk persetujuan Manajer", data)
}

// ApproveMailingCampaign godoc
// @Summary Approve a MailingCampaign
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id}/approve [put]
// @Security BearerAuth
func ApproveMailingCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMailingCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	data.ApprovalStatus = "Approved"
	now := time.Now()
	data.ApprovedAt = &now
	data.RejectReason = nil
	if err := UpdateMailingCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to approve campaign", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Kampanye berhasil disetujui (Approved)", data)
}

// RejectMailingCampaign godoc
// @Summary Reject a MailingCampaign
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id}/reject [put]
// @Security BearerAuth
func RejectMailingCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMailingCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	var payload struct {
		Reason string `json:"reason"`
	}
	_ = c.Bind(&payload)
	data.ApprovalStatus = "Rejected"
	data.RejectReason = &payload.Reason
	if err := UpdateMailingCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to reject campaign", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Kampanye telah ditolak (Rejected)", data)
}

// DeleteMailingCampaign godoc
// @Summary Delete a MailingCampaign
// @Description Delete a MailingCampaign by ID
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id} [delete]
// @Security BearerAuth
func DeleteMailingCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMailingCampaignService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create UtmTracker
// @Description Create a new UtmTracker
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Success 201 {object} UtmTracker
// @Param request body UtmTracker true "Payload"
// @Router /api/marketing/mass_mailing/utmtracker [post]
// @Security BearerAuth
func CreateUtmTrackerHandler(c echo.Context) error {
	var data UtmTracker
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateUtmTrackerService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all UtmTracker
// @Description Retrieve a list of all UtmTracker
// @Tags marketing-mass_mailing
// @Produce json
// @Success 200 {object} UtmTracker
// @Router /api/marketing/mass_mailing/utmtracker [get]
// @Security BearerAuth
func GetAllUtmTrackerHandler(c echo.Context) error {
	data, err := GetAllUtmTrackerService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetUtmTrackerByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUtmTrackerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update UtmTracker
// @Description Update an existing UtmTracker
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Param id path int true "UtmTracker ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/utmtracker/{id} [put]
// @Security BearerAuth
func UpdateUtmTrackerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUtmTrackerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateUtmTrackerService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete UtmTracker
// @Description Delete UtmTracker by ID
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "UtmTracker ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/utmtracker/{id} [delete]
// @Security BearerAuth
func DeleteUtmTrackerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteUtmTrackerService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// RunABSplitTestHandler godoc
// @Summary Run A/B Split Test
// @Description Evaluate Version A vs Version B and automatically pick the winner for blast
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id}/ab-test [post]
// @Security BearerAuth
func RunABSplitTestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := RunABSplitTestService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to run A/B split test", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "A/B Split Test berhasil dievaluasi! Pemenang telah ditentukan.", data)
}



