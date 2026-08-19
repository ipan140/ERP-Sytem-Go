package crm

func CreateLeadService(data *Lead) error {
	return CreateLead(data)
}

func GetAllLeadService() ([]Lead, error) {
	return GetAllLead()
}

func GetLeadByIDService(id uint) (*Lead, error) {
	return GetLeadByID(id)
}

func UpdateLeadService(data *Lead) error {
	return UpdateLead(data)
}

func DeleteLeadService(id uint) error {
	return DeleteLead(id)
}

func CreateSalesTeamService(data *SalesTeam) error { return CreateSalesTeam(data) }
func GetAllSalesTeamService() ([]SalesTeam, error) { return GetAllSalesTeam() }
func GetSalesTeamByIDService(id uint) (*SalesTeam, error) { return GetSalesTeamByID(id) }
func UpdateSalesTeamService(data *SalesTeam) error { return UpdateSalesTeam(data) }
func DeleteSalesTeamService(id uint) error { return DeleteSalesTeam(id) }

func CreateStageService(data *Stage) error { return CreateStage(data) }
func GetAllStageService() ([]Stage, error) { return GetAllStage() }
func GetStageByIDService(id uint) (*Stage, error) { return GetStageByID(id) }
func UpdateStageService(data *Stage) error { return UpdateStage(data) }
func DeleteStageService(id uint) error { return DeleteStage(id) }

func CreateActivityService(data *Activity) error { return CreateActivity(data) }
func GetAllActivityService() ([]Activity, error) { return GetAllActivity() }
func GetActivityByIDService(id uint) (*Activity, error) { return GetActivityByID(id) }
func UpdateActivityService(data *Activity) error { return UpdateActivity(data) }
func DeleteActivityService(id uint) error { return DeleteActivity(id) }

func CreateSalesCommissionService(data *SalesCommission) error { return CreateSalesCommission(data) }
func GetAllSalesCommissionService() ([]SalesCommission, error) { return GetAllSalesCommission() }
func GetSalesCommissionByIDService(id uint) (*SalesCommission, error) { return GetSalesCommissionByID(id) }
func UpdateSalesCommissionService(data *SalesCommission) error { return UpdateSalesCommission(data) }
func DeleteSalesCommissionService(id uint) error { return DeleteSalesCommission(id) }
