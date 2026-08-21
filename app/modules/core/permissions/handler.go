package permissions

import (
	"net/http"

	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// GetAllPermissionsHandler godoc
// @Summary Get all role permissions
// @Description Retrieve a list of all role permissions and their toggle states
// @Tags Core - Permissions (Dynamic RBAC)
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /core/permissions [get]
func GetAllPermissionsHandler(c echo.Context) error {
	perms, err := GetAllPermissionsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Berhasil mengambil konfigurasi toggle", perms)
}

// TogglePermissionHandler godoc
// @Summary Toggle a specific permission
// @Description Turn a permission (read, write, delete) ON or OFF for a specific role and module
// @Tags Core - Permissions (Dynamic RBAC)
// @Accept json
// @Produce json
// @Param body body TogglePermissionRequest true "Toggle Request Payload"
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /core/permissions/toggle [post]
func TogglePermissionHandler(c echo.Context) error {
	var req TogglePermissionRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format Data Salah", err.Error())
	}

	err := TogglePermissionService(req)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan toggle", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Toggle Akses berhasil diperbarui!", nil)
}

// GetAvailableModulesHandler godoc
// @Summary Get list of available modules
// @Description Retrieve a list of all ERP modules available for permission toggling
// @Tags Core - Permissions (Dynamic RBAC)
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /core/permissions/modules [get]
func GetAvailableModulesHandler(c echo.Context) error {
	modules := []string{
		"core/artificial_intelligence", "core/base", "core/dashboards", "core/discuss",
		"core/documents", "core/iot", "core/knowledge", "core/mailer", "core/permissions",
		"core/report", "core/storage", "core/user_roles", "core/voip", "core/whatsapp",

		"finance/accounting", "finance/approvals", "finance/consolidation", "finance/documents",
		"finance/expenses", "finance/invoicing", "finance/sign", "finance/spreadsheet_bi",

		"hr/appraisals", "hr/attendances", "hr/employees", "hr/fleet", "hr/lunch",
		"hr/payroll", "hr/recruitment", "hr/referrals", "hr/time_off",

		"marketing/events", "marketing/marketing_automation", "marketing/mass_mailing",
		"marketing/sms_marketing", "marketing/social_marketing", "marketing/surveys",

		"sales/crm", "sales/point_of_sale", "sales/rental", "sales/sales_core", "sales/subscriptions",

		"services/appointments", "services/field_service", "services/helpdesk", "services/planning",
		"services/project", "services/repairs", "services/timesheets",

		"supply_chain/barcode", "supply_chain/inventory", "supply_chain/maintenance",
		"supply_chain/manufacturing", "supply_chain/plm", "supply_chain/purchase", "supply_chain/quality",

		"website/blog", "website/ecommerce", "website/elearning", "website/forum",
		"website/live_chat", "website/website_builder",
	}
	return utils.SendSuccess(c, http.StatusOK, "Daftar modul ERP berhasil diambil", modules)
}
