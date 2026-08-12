package planning

func CreateShiftService(data *Shift) error {
	return CreateShift(data)
}

func GetAllShiftService() ([]Shift, error) {
	return GetAllShift()
}

func GetShiftByIDService(id uint) (*Shift, error) {
	return GetShiftByID(id)
}

func UpdateShiftService(data *Shift) error {
	return UpdateShift(data)
}

func DeleteShiftService(id uint) error {
	return DeleteShift(id)
}
