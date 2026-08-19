package project

import (
	"ERP-System/config"
)

func CreateProject(data *Project) error {
	return config.DB.Create(data).Error
}

func GetAllProject() ([]Project, error) {
	var list []Project
	err := config.DB.Find(&list).Error
	return list, err
}

func GetProjectByID(id uint) (*Project, error) {
	var data Project
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateProject(data *Project) error {
	return config.DB.Save(data).Error
}

func DeleteProject(id uint) error {
	return config.DB.Delete(&Project{}, id).Error
}

func CreateProjectMilestone(data *ProjectMilestone) error { return config.DB.Create(data).Error }
func GetAllProjectMilestone() ([]ProjectMilestone, error) { var list []ProjectMilestone; err := config.DB.Find(&list).Error; return list, err }
func GetProjectMilestoneByID(id uint) (*ProjectMilestone, error) { var data ProjectMilestone; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateProjectMilestone(data *ProjectMilestone) error { return config.DB.Save(data).Error }
func DeleteProjectMilestone(id uint) error { return config.DB.Delete(&ProjectMilestone{}, id).Error }

func CreateTaskDependency(data *TaskDependency) error { return config.DB.Create(data).Error }
func GetAllTaskDependency() ([]TaskDependency, error) { var list []TaskDependency; err := config.DB.Find(&list).Error; return list, err }
func GetTaskDependencyByID(id uint) (*TaskDependency, error) { var data TaskDependency; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateTaskDependency(data *TaskDependency) error { return config.DB.Save(data).Error }
func DeleteTaskDependency(id uint) error { return config.DB.Delete(&TaskDependency{}, id).Error }

func CreateResourceForecast(data *ResourceForecast) error { return config.DB.Create(data).Error }
func GetAllResourceForecast() ([]ResourceForecast, error) { var list []ResourceForecast; err := config.DB.Find(&list).Error; return list, err }
func GetResourceForecastByID(id uint) (*ResourceForecast, error) { var data ResourceForecast; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateResourceForecast(data *ResourceForecast) error { return config.DB.Save(data).Error }
func DeleteResourceForecast(id uint) error { return config.DB.Delete(&ResourceForecast{}, id).Error }


func CreateTask(data *Task) error { return config.DB.Create(data).Error }
func GetAllTask() ([]Task, error) { var list []Task; err := config.DB.Find(&list).Error; return list, err }
func GetTaskByID(id uint) (*Task, error) { var data Task; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateTask(data *Task) error { return config.DB.Save(data).Error }
func DeleteTask(id uint) error { return config.DB.Delete(&Task{}, id).Error }
