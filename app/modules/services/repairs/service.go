package repairs

func CreateRepairOrderService(data *RepairOrder) error {
	return CreateRepairOrder(data)
}

func GetAllRepairOrderService() ([]RepairOrder, error) {
	return GetAllRepairOrder()
}

func GetPaginatedRepairOrderService(offset int, limit int, search string, state string, warranty string, companyID uint, technicianID uint) ([]RepairOrder, int64, error) {
	return GetPaginatedRepairOrders(offset, limit, search, state, warranty, companyID, technicianID)
}

func GetRepairOrderByIDService(id uint) (*RepairOrder, error) {
	return GetRepairOrderByID(id)
}

func UpdateRepairOrderService(data *RepairOrder) error {
	return UpdateRepairOrder(data)
}

func DeleteRepairOrderService(id uint) error {
	return DeleteRepairOrder(id)
}

// Fase 2: Quality Control Gate
func PassQCRepairOrderService(id uint, inspectorID uint, notes string) error {
	return PassQCRepairOrder(id, inspectorID, notes)
}

// Fase 4: Public Tracking & Customer Approval
func GetRepairByRMAAndSerialService(id uint, serialNumber string) (*RepairOrder, error) {
	return GetRepairByRMAAndSerial(id, serialNumber)
}

func ApproveRepairEstimateService(id uint, note string) (*RepairOrder, error) {
	return ApproveRepairEstimate(id, note)
}

