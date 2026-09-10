package dashboard

import (
	"ERP-System/common/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetScmDashboardSummaryHandler(c echo.Context) error {
	summary, err := GetScmDashboardSummaryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve SCM dashboard summary", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "SCM dashboard summary retrieved successfully", summary)
}

func GetScmCalendarEventsHandler(c echo.Context) error {
	events, err := GetScmCalendarEventsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve SCM calendar events", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "SCM calendar events retrieved successfully", events)
}
