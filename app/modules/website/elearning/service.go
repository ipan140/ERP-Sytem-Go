package elearning

func CreateCourseService(data *Course) error {
	return CreateCourse(data)
}

func GetAllCourseService() ([]Course, error) {
	return GetAllCourse()
}

func GetPaginatedCourseService(offset, limit int, search string) ([]Course, int64, error) {
	return GetPaginatedCourses(offset, limit, search)
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

func CreateSlideService(data *Slide) error        { return CreateSlide(data) }
func GetAllSlideService() ([]Slide, error)        { return GetAllSlide() }
func GetSlideByIDService(id uint) (*Slide, error) { return GetSlideByID(id) }
func UpdateSlideService(data *Slide) error        { return UpdateSlide(data) }
func DeleteSlideService(id uint) error            { return DeleteSlide(id) }

func CreateCertificationService(data *Certification) error        { return CreateCertification(data) }
func GetAllCertificationService() ([]Certification, error)        { return GetAllCertification() }
func GetCertificationByIDService(id uint) (*Certification, error) { return GetCertificationByID(id) }
func UpdateCertificationService(data *Certification) error        { return UpdateCertification(data) }
func DeleteCertificationService(id uint) error                    { return DeleteCertification(id) }
