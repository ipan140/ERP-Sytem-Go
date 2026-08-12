package elearning

func CreateCourseService(data *Course) error {
	return CreateCourse(data)
}

func GetAllCourseService() ([]Course, error) {
	return GetAllCourse()
}

func GetCourseByIDService(id uint) (*Course, error) {
	return GetCourseByID(id)
}

func UpdateCourseService(data *Course) error {
	return UpdateCourse(data)
}

func DeleteCourseService(id uint) error {
	return DeleteCourse(id)
}
