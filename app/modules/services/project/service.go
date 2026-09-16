package project

import "fmt"

func CreateProjectService(data *Project) error {
	return CreateProject(data)
}

func GetAllProjectService() ([]Project, error) {
	return GetAllProject()
}

func GetPaginatedProjectService(offset, limit int, search string) ([]Project, int64, error) {
	return GetPaginatedProject(offset, limit, search)
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

func CreateProjectMilestoneService(data *ProjectMilestone) error { return CreateProjectMilestone(data) }
func GetAllProjectMilestoneService() ([]ProjectMilestone, error) { return GetAllProjectMilestone() }
func GetProjectMilestoneByIDService(id uint) (*ProjectMilestone, error) {
	return GetProjectMilestoneByID(id)
}
func UpdateProjectMilestoneService(data *ProjectMilestone) error {
	if data.IsReached {
		TriggerMilestoneBilling(data.ID)
	}
	return UpdateProjectMilestone(data)
}
func DeleteProjectMilestoneService(id uint) error { return DeleteProjectMilestone(id) }

func CreateTaskDependencyService(data *TaskDependency) error        { return CreateTaskDependency(data) }
func GetAllTaskDependencyService() ([]TaskDependency, error)        { return GetAllTaskDependency() }
func GetTaskDependencyByIDService(id uint) (*TaskDependency, error) { return GetTaskDependencyByID(id) }
func UpdateTaskDependencyService(data *TaskDependency) error        { return UpdateTaskDependency(data) }
func DeleteTaskDependencyService(id uint) error                     { return DeleteTaskDependency(id) }

func CreateResourceForecastService(data *ResourceForecast) error { return CreateResourceForecast(data) }
func GetAllResourceForecastService() ([]ResourceForecast, error) { return GetAllResourceForecast() }
func GetResourceForecastByIDService(id uint) (*ResourceForecast, error) {
	return GetResourceForecastByID(id)
}
func UpdateResourceForecastService(data *ResourceForecast) error { return UpdateResourceForecast(data) }
func DeleteResourceForecastService(id uint) error                { return DeleteResourceForecast(id) }

func CreateTaskService(data *Task) error        { return CreateTask(data) }
func GetAllTaskService() ([]Task, error)        { return GetAllTask() }
func GetTaskByIDService(id uint) (*Task, error) { return GetTaskByID(id) }
func UpdateTaskService(data *Task) error {
	if data.AssigneeID != 0 {
		CheckTechnicianSkillEligibility(data.ID, data.AssigneeID)
	}
	if data.Stage != "todo" {
		if !CanStartTask(data.ID) {
			return fmt.Errorf("Tugas diblokir oleh dependensi yang belum selesai!")
		}
	}
	return UpdateTask(data)
}
func DeleteTaskService(id uint) error { return DeleteTask(id) }
