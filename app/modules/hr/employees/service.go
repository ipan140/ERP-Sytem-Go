package employees

func CreateEmployeeService(data *Employee) error {
	return CreateEmployee(data)
}

func GetAllEmployeeService() ([]Employee, error) {
	return GetAllEmployee()
}

func GetEmployeeByIDService(id uint) (*Employee, error) {
	return GetEmployeeByID(id)
}

func UpdateEmployeeService(data *Employee) error {
	return UpdateEmployee(data)
}

func DeleteEmployeeService(id uint) error {
	return DeleteEmployee(id)
}
