package plm

func CreateBomService(data *PlmEco) error {
	return CreateBom(data)
}

func GetAllBomService() ([]PlmEco, error) {
	return GetAllBom()
}

func GetBomByIDService(id uint) (*PlmEco, error) {
	return GetBomByID(id)
}

func UpdateBomService(data *PlmEco) error {
	return UpdateBom(data)
}

func DeleteBomService(id uint) error {
	return DeleteBom(id)
}
