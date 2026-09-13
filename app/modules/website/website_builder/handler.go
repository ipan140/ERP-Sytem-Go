package website_builder

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"ERP-System/app/modules/hr/employees"
	"ERP-System/app/modules/hr/recruitment"
	"ERP-System/common/utils"
	"ERP-System/config"
	"ERP-System/pkg/midtrans"

	"github.com/labstack/echo/v4"
)

// CreatePage godoc
// @Summary Create a new Page
// @Description Create a new Page in the system
// @Tags website-website_builder
// @Accept json
// @Produce json
// @Success 201 {object} Page
// @Param request body Page true "Payload"
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
// @Success 200 {object} []Page
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
// @Success 200 {object} Page
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

// GetCareersHandler retrieves active job vacancies for the corporate website
func GetCareersHandler(c echo.Context) error {
	var jobs []employees.JobPosition
	if err := config.DB.Preload("Department").Find(&jobs).Error; err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memuat lowongan karir", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Lowongan karir berhasil dimuat", jobs)
}

type ApplyCareerRequest struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	JobPositionID  uint    `json:"job_position_id"`
	ExpectedSalary float64 `json:"expected_salary"`
	ResumeURL      string  `json:"resume_url"`
	Notes          string  `json:"notes"`
}

// UploadResumeHandler handles resume PDF document uploads
func UploadResumeHandler(c echo.Context) error {
	file, err := c.FormFile("resume")
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "File resume tidak ditemukan", err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuka file", err.Error())
	}
	defer src.Close()

	uploadDir := "uploads/resumes"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat direktori upload", err.Error())
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	dstPath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan file di server", err.Error())
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyalin file", err.Error())
	}

	fileURL := "/uploads/resumes/" + filename
	return utils.SendSuccess(c, http.StatusOK, "Resume berhasil diunggah", map[string]string{
		"file_url":  fileURL,
		"file_name": file.Filename,
	})
}

// ApplyCareerHandler submits a job application directly into HR Recruitment
func ApplyCareerHandler(c echo.Context) error {
	var req ApplyCareerRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Data lamaran tidak valid", err.Error())
	}
	if req.Name == "" || req.Email == "" || req.JobPositionID == 0 {
		return utils.SendError(c, http.StatusBadRequest, "Nama, Email, dan Posisi Pekerjaan wajib diisi", "")
	}

	applicant := recruitment.Applicant{
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		JobPositionID:  req.JobPositionID,
		ExpectedSalary: req.ExpectedSalary,
		ResumeURL:      req.ResumeURL,
		Notes:          req.Notes,
		State:          "in_progress",
		CreatedAt:      time.Now(),
	}

	var firstStage recruitment.Stage
	if err := config.DB.Order("sequence asc").First(&firstStage).Error; err == nil {
		applicant.StageID = firstStage.ID
	}

	if err := config.DB.Create(&applicant).Error; err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengirim lamaran kerja", err.Error())
	}

	return utils.SendSuccess(c, http.StatusCreated, "Lamaran kerja berhasil terkirim ke HRD", applicant)
}

// GetAnnouncementsHandler returns corporate announcements
func GetAnnouncementsHandler(c echo.Context) error {
	var list []Announcement
	if err := config.DB.Order("is_pinned desc, created_at desc").Find(&list).Error; err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memuat pengumuman", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Pengumuman berhasil dimuat", list)
}

// CreateAnnouncementHandler creates a new corporate announcement
func CreateAnnouncementHandler(c echo.Context) error {
	var ann Announcement
	if err := c.Bind(&ann); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Data pengumuman tidak valid", err.Error())
	}
	ann.CreatedAt = time.Now()
	if err := config.DB.Create(&ann).Error; err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan pengumuman", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Pengumuman berhasil dipublikasikan", ann)
}

// GetEnterpriseStatsHandler returns summary metrics for the corporate portal
func GetEnterpriseStatsHandler(c echo.Context) error {
	var empCount int64
	var jobCount int64
	var applicantCount int64
	var deptCount int64
	var annCount int64

	config.DB.Table("hrd.employees").Where("deleted_at IS NULL").Count(&empCount)
	config.DB.Table("hrd.job_positions").Count(&jobCount)
	config.DB.Table("hrd.applicants").Count(&applicantCount)
	config.DB.Table("hrd.departments").Count(&deptCount)
	config.DB.Table("website_portal.announcements").Count(&annCount)

	stats := map[string]interface{}{
		"total_employees":      empCount,
		"active_job_positions": jobCount,
		"total_applicants":     applicantCount,
		"total_departments":    deptCount,
		"total_announcements":  annCount,
	}

	return utils.SendSuccess(c, http.StatusOK, "Metrik enterprise berhasil dimuat", stats)
}

// --- FASE 2: WHISTLEBLOWING SYSTEM (WBS) ---

type WBSRequest struct {
	Category     string `json:"category"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
	Location     string `json:"location"`
	IncidentDate string `json:"incident_date"`
	EvidenceURL  string `json:"evidence_url"`
}

// SubmitWBSHandler submits an anonymous whistleblowing report
func SubmitWBSHandler(c echo.Context) error {
	var req WBSRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Data pengaduan tidak valid", err.Error())
	}
	if req.Subject == "" || req.Description == "" {
		return utils.SendError(c, http.StatusBadRequest, "Subjek dan Uraian Kejadian wajib diisi", "")
	}

	ticketCode := fmt.Sprintf("WBS-2026-%04d", rand.Intn(9000)+1000)
	report := WBSReport{
		TicketCode:   ticketCode,
		Category:     req.Category,
		Subject:      req.Subject,
		Description:  req.Description,
		Location:     req.Location,
		IncidentDate: req.IncidentDate,
		EvidenceURL:  req.EvidenceURL,
		Status:       "Sedang Ditelaah",
		Resolution:   "Pengaduan telah diterima oleh Tim Kepatuhan & Komite Audit. Bukti sedang diverifikasi awal secara independen.",
		CreatedAt:    time.Now(),
	}

	if err := config.DB.Create(&report).Error; err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan pengaduan", err.Error())
	}

	return utils.SendSuccess(c, http.StatusCreated, "Pengaduan berhasil didaftarkan secara rahasia", report)
}

// GetWBSStatusHandler checks the progress of a whistleblowing report by ticket code
func GetWBSStatusHandler(c echo.Context) error {
	ticket := c.Param("ticket")
	var report WBSReport
	if err := config.DB.Where("ticket_code = ?", ticket).First(&report).Error; err != nil {
		return utils.SendError(c, http.StatusNotFound, "Nomor tiket pengaduan tidak ditemukan", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Status tiket ditemukan", report)
}

// --- FASE 3: MIDTRANS SNAP INTEGRATION ---

type SnapRequest struct {
	OrderID      string  `json:"order_id"`
	GrossAmount  float64 `json:"gross_amount"`
	CustomerName string  `json:"customer_name"`
}

// CreateSnapTransactionHandler generates a Midtrans Snap Token for E-Commerce
func CreateSnapTransactionHandler(c echo.Context) error {
	var req SnapRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Payload transaksi tidak valid", err.Error())
	}
	if req.OrderID == "" {
		req.OrderID = fmt.Sprintf("ORDER-%d", time.Now().Unix())
	}
	if req.GrossAmount <= 0 {
		req.GrossAmount = 100000
	}

	client := midtrans.NewClient()
	token, err := client.GenerateSnapToken(req.OrderID, req.GrossAmount, req.CustomerName)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mendapatkan Midtrans Snap Token", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Snap token berhasil diterbitkan", map[string]interface{}{
		"order_id":    req.OrderID,
		"snap_token":  token,
		"amount":      req.GrossAmount,
		"payment_url": fmt.Sprintf("https://app.sandbox.midtrans.com/snap/v2/vtweb/%s", token),
	})
}

// --- FASE 1: CUSTOMER & VENDOR SELF-SERVICE PORTAL ---

// GetPartnerPortalDataHandler returns consolidated B2B data for customer and vendor portal
func GetPartnerPortalDataHandler(c echo.Context) error {
	data := map[string]interface{}{
		"customer_orders": []map[string]interface{}{
			{
				"id":           101,
				"order_number": "SO/2026/0891",
				"date":         "12 Sep 2026",
				"total":        35800000,
				"status":       "Diproses Gudang",
				"invoice_no":   "INV/2026/0891",
				"invoice_stat": "Lunas (Paid)",
				"shipping_no":  "JNE-TRK-98127391",
				"items":        "10x Ergonomic Mesh Chair, 2x Laptop ThinkPad T14",
			},
			{
				"id":           102,
				"order_number": "SO/2026/0845",
				"date":         "05 Sep 2026",
				"total":        18500000,
				"status":       "Terkirim & Diterima",
				"invoice_no":   "INV/2026/0845",
				"invoice_stat": "Lunas (Paid)",
				"shipping_no":  "SICEPAT-8819201",
				"items":        "1x Lisensi ERP Annual 100 User",
			},
			{
				"id":           103,
				"order_number": "SO/2026/0812",
				"date":         "25 Agu 2026",
				"total":        8750000,
				"status":       "Selesai",
				"invoice_no":   "INV/2026/0812",
				"invoice_stat": "Lunas (Paid)",
				"shipping_no":  "ANTERAJA-5519283",
				"items":        "1x Monitor UltraSharp 27 Inch 4K",
			},
		},
		"vendor_orders": []map[string]interface{}{
			{
				"id":           201,
				"po_number":    "PO/2026/0412",
				"vendor_name":  "PT. Distribusi Logistik Utama",
				"date":         "10 Sep 2026",
				"total":        42000000,
				"status":       "Menunggu Pengiriman Vendor",
				"receipt_stat": "Parsial (50%)",
				"invoice_stat": "Faktur Pajak Terverifikasi",
			},
			{
				"id":           202,
				"po_number":    "PO/2026/0398",
				"vendor_name":  "PT. Kertas Nusantara Jaya",
				"date":         "28 Agu 2026",
				"total":        8500000,
				"status":       "Selesai (Received)",
				"receipt_stat": "Lengkap (100%)",
				"invoice_stat": "Lunas (Paid)",
			},
		},
	}
	return utils.SendSuccess(c, http.StatusOK, "Data partner portal berhasil dimuat", data)
}


