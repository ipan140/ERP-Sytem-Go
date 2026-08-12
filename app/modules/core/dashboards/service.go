package dashboards

func CreateDashboardService(data *Dashboard) error {
	return CreateDashboard(data)
}

func GetAllDashboardService() ([]Dashboard, error) {
	return GetAllDashboard()
}

func GetDashboardByIDService(id uint) (*Dashboard, error) {
	return GetDashboardByID(id)
}

func UpdateDashboardService(data *Dashboard) error {
	return UpdateDashboard(data)
}

func DeleteDashboardService(id uint) error {
	return DeleteDashboard(id)
}
