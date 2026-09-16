package documents

func CreateWorkspaceService(data *Workspace) error {
	return CreateWorkspace(data)
}

func GetAllWorkspaceService() ([]Workspace, error) {
	return GetAllWorkspace()
}

func GetPaginatedWorkspaceService(offset, limit int, search string) ([]Workspace, int64, error) {
	return GetPaginatedWorkspaces(offset, limit, search)
}

func GetWorkspaceByIDService(id uint) (*Workspace, error) {
	return GetWorkspaceByID(id)
}

func UpdateWorkspaceService(data *Workspace) error {
	return UpdateWorkspace(data)
}

func DeleteWorkspaceService(id uint) error {
	return DeleteWorkspace(id)
}
