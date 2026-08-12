package plm

func CreateBomService(data *Bom) error {
	return CreateBom(data)
}

func GetAllBomService() ([]Bom, error) {
	return GetAllBom()
}

func GetBomByIDService(id uint) (*Bom, error) {
	return GetBomByID(id)
}

func UpdateBomService(data *Bom) error {
	return UpdateBom(data)
}

func DeleteBomService(id uint) error {
	return DeleteBom(id)
}
