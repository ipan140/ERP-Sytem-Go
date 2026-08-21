package point_of_sale

func CreatePosSessionService(data *PosSession) error {
	return CreatePosSession(data)
}

func GetAllPosSessionService() ([]PosSession, error) {
	return GetAllPosSession()
}

func GetPosSessionByIDService(id uint) (*PosSession, error) {
	return GetPosSessionByID(id)
}

func UpdatePosSessionService(data *PosSession) error {
	return UpdatePosSession(data)
}

func DeletePosSessionService(id uint) error {
	return DeletePosSession(id)
}

func CreatePosConfigService(data *PosConfig) error        { return CreatePosConfig(data) }
func GetAllPosConfigService() ([]PosConfig, error)        { return GetAllPosConfig() }
func GetPosConfigByIDService(id uint) (*PosConfig, error) { return GetPosConfigByID(id) }
func UpdatePosConfigService(data *PosConfig) error        { return UpdatePosConfig(data) }
func DeletePosConfigService(id uint) error                { return DeletePosConfig(id) }

func CreatePosOrderService(data *PosOrder) error        { return CreatePosOrder(data) }
func GetAllPosOrderService() ([]PosOrder, error)        { return GetAllPosOrder() }
func GetPosOrderByIDService(id uint) (*PosOrder, error) { return GetPosOrderByID(id) }
func UpdatePosOrderService(data *PosOrder) error        { return UpdatePosOrder(data) }
func DeletePosOrderService(id uint) error               { return DeletePosOrder(id) }

func CreatePosOrderLineService(data *PosOrderLine) error        { return CreatePosOrderLine(data) }
func GetAllPosOrderLineService() ([]PosOrderLine, error)        { return GetAllPosOrderLine() }
func GetPosOrderLineByIDService(id uint) (*PosOrderLine, error) { return GetPosOrderLineByID(id) }
func UpdatePosOrderLineService(data *PosOrderLine) error        { return UpdatePosOrderLine(data) }
func DeletePosOrderLineService(id uint) error                   { return DeletePosOrderLine(id) }

func CreatePosPaymentService(data *PosPayment) error        { return CreatePosPayment(data) }
func GetAllPosPaymentService() ([]PosPayment, error)        { return GetAllPosPayment() }
func GetPosPaymentByIDService(id uint) (*PosPayment, error) { return GetPosPaymentByID(id) }
func UpdatePosPaymentService(data *PosPayment) error        { return UpdatePosPayment(data) }
func DeletePosPaymentService(id uint) error                 { return DeletePosPayment(id) }

func CreateLoyaltyProgramService(data *LoyaltyProgram) error        { return CreateLoyaltyProgram(data) }
func GetAllLoyaltyProgramService() ([]LoyaltyProgram, error)        { return GetAllLoyaltyProgram() }
func GetLoyaltyProgramByIDService(id uint) (*LoyaltyProgram, error) { return GetLoyaltyProgramByID(id) }
func UpdateLoyaltyProgramService(data *LoyaltyProgram) error        { return UpdateLoyaltyProgram(data) }
func DeleteLoyaltyProgramService(id uint) error                     { return DeleteLoyaltyProgram(id) }
