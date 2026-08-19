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

func CreateSlideService(data *Slide) error { return CreateSlide(data) }
func GetAllSlideService() ([]Slide, error) { return GetAllSlide() }
func GetSlideByIDService(id uint) (*Slide, error) { return GetSlideByID(id) }
func UpdateSlideService(data *Slide) error { return UpdateSlide(data) }
func DeleteSlideService(id uint) error { return DeleteSlide(id) }
