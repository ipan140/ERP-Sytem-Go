package lunch

func CreateLunchOrderService(data *LunchOrder) error {
	return CreateLunchOrder(data)
}

func GetAllLunchOrderService() ([]LunchOrder, error) {
	return GetAllLunchOrder()
}

func GetLunchOrderByIDService(id uint) (*LunchOrder, error) {
	return GetLunchOrderByID(id)
}

func UpdateLunchOrderService(data *LunchOrder) error {
	return UpdateLunchOrder(data)
}

func DeleteLunchOrderService(id uint) error {
	return DeleteLunchOrder(id)
}

func CreateLunchCashmoveService(data *LunchCashmove) error        { return CreateLunchCashmove(data) }
func GetAllLunchCashmoveService() ([]LunchCashmove, error)        { return GetAllLunchCashmove() }
func GetLunchCashmoveByIDService(id uint) (*LunchCashmove, error) { return GetLunchCashmoveByID(id) }
func UpdateLunchCashmoveService(data *LunchCashmove) error        { return UpdateLunchCashmove(data) }
func DeleteLunchCashmoveService(id uint) error                    { return DeleteLunchCashmove(id) }
