package project

func CreateProjectService(data *Project) error {
	return CreateProject(data)
}

func GetAllProjectService() ([]Project, error) {
	return GetAllProject()
}

func GetProjectByIDService(id uint) (*Project, error) {
	return GetProjectByID(id)
}

func UpdateProjectService(data *Project) error {
	return UpdateProject(data)
}

func DeleteProjectService(id uint) error {
	return DeleteProject(id)
}
