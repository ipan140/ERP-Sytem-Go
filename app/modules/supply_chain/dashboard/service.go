package dashboard

func GetScmDashboardSummaryService() (*ScmDashboardSummary, error) {
	return GetScmDashboardSummary()
}

func GetScmCalendarEventsService() ([]ScmCalendarEvent, error) {
	return GetScmCalendarEvents()
}
