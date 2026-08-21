package employees

import (
	"ERP-System/config"
)

func CreateEmployee(data *Employee) error {
	return config.DB.Create(data).Error
}

func GetAllEmployee() ([]Employee, error) {
	var list []Employee
	err := config.DB.Find(&list).Error
	return list, err
}

func GetEmployeeByID(id uint) (*Employee, error) {
	var data Employee
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateEmployee(data *Employee) error {
	return config.DB.Save(data).Error
}

func DeleteEmployee(id uint) error {
	return config.DB.Delete(&Employee{}, id).Error
}

func CreateJobPosition(data *JobPosition) error { return config.DB.Create(data).Error }
func GetAllJobPosition() ([]JobPosition, error) {
	var list []JobPosition
	err := config.DB.Find(&list).Error
	return list, err
}
func GetJobPositionByID(id uint) (*JobPosition, error) {
	var data JobPosition
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateJobPosition(data *JobPosition) error { return config.DB.Save(data).Error }
func DeleteJobPosition(id uint) error           { return config.DB.Delete(&JobPosition{}, id).Error }

func CreateWorkingSchedule(data *WorkingSchedule) error { return config.DB.Create(data).Error }
func GetAllWorkingSchedule() ([]WorkingSchedule, error) {
	var list []WorkingSchedule
	err := config.DB.Find(&list).Error
	return list, err
}
func GetWorkingScheduleByID(id uint) (*WorkingSchedule, error) {
	var data WorkingSchedule
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateWorkingSchedule(data *WorkingSchedule) error { return config.DB.Save(data).Error }
func DeleteWorkingSchedule(id uint) error               { return config.DB.Delete(&WorkingSchedule{}, id).Error }

func CreateContract(data *Contract) error { return config.DB.Create(data).Error }
func GetAllContract() ([]Contract, error) {
	var list []Contract
	err := config.DB.Find(&list).Error
	return list, err
}
func GetContractByID(id uint) (*Contract, error) {
	var data Contract
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateContract(data *Contract) error { return config.DB.Save(data).Error }
func DeleteContract(id uint) error        { return config.DB.Delete(&Contract{}, id).Error }

func CreateSkill(data *Skill) error { return config.DB.Create(data).Error }
func GetAllSkill() ([]Skill, error) {
	var list []Skill
	err := config.DB.Find(&list).Error
	return list, err
}
func GetSkillByID(id uint) (*Skill, error) {
	var data Skill
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateSkill(data *Skill) error { return config.DB.Save(data).Error }
func DeleteSkill(id uint) error     { return config.DB.Delete(&Skill{}, id).Error }

func CreateSkillLevel(data *SkillLevel) error { return config.DB.Create(data).Error }
func GetAllSkillLevel() ([]SkillLevel, error) {
	var list []SkillLevel
	err := config.DB.Find(&list).Error
	return list, err
}
func GetSkillLevelByID(id uint) (*SkillLevel, error) {
	var data SkillLevel
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateSkillLevel(data *SkillLevel) error { return config.DB.Save(data).Error }
func DeleteSkillLevel(id uint) error          { return config.DB.Delete(&SkillLevel{}, id).Error }

func CreateEmployeeSkill(data *EmployeeSkill) error { return config.DB.Create(data).Error }
func GetAllEmployeeSkill() ([]EmployeeSkill, error) {
	var list []EmployeeSkill
	err := config.DB.Find(&list).Error
	return list, err
}
func GetEmployeeSkillByID(id uint) (*EmployeeSkill, error) {
	var data EmployeeSkill
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateEmployeeSkill(data *EmployeeSkill) error { return config.DB.Save(data).Error }
func DeleteEmployeeSkill(id uint) error             { return config.DB.Delete(&EmployeeSkill{}, id).Error }

func CreateResumeLine(data *ResumeLine) error { return config.DB.Create(data).Error }
func GetAllResumeLine() ([]ResumeLine, error) {
	var list []ResumeLine
	err := config.DB.Find(&list).Error
	return list, err
}
func GetResumeLineByID(id uint) (*ResumeLine, error) {
	var data ResumeLine
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateResumeLine(data *ResumeLine) error { return config.DB.Save(data).Error }
func DeleteResumeLine(id uint) error          { return config.DB.Delete(&ResumeLine{}, id).Error }
