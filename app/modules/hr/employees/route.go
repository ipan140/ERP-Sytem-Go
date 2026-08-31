package employees

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// 1. Group HR Dasar: Minimal harus punya role EMPLOYEE (Karyawan) atau HR_MANAGER
	api := e.Group("/api/hr/employees", middleware.Auth(), middleware.GlobalAutoRBAC())

	// Karyawan biasa boleh melihat profil (GET)
	api.GET("", GetAllEmployeeHandler)
	api.GET("/:id", GetEmployeeByIDHandler)
	api.GET("/jobposition", GetAllJobPositionHandler)
	api.GET("/jobposition/:id", GetJobPositionByIDHandler)
	api.GET("/workingschedule", GetAllWorkingScheduleHandler)
	api.GET("/workingschedule/:id", GetWorkingScheduleByIDHandler)
	api.GET("/skill", GetAllSkillHandler)
	api.GET("/skill/:id", GetSkillByIDHandler)
	api.GET("/skilllevel", GetAllSkillLevelHandler)
	api.GET("/skilllevel/:id", GetSkillLevelByIDHandler)
	api.GET("/employeeskill", GetAllEmployeeSkillHandler)
	api.GET("/employeeskill/:id", GetEmployeeSkillByIDHandler)
	api.GET("/resumeline", GetAllResumeLineHandler)
	api.GET("/resumeline/:id", GetResumeLineByIDHandler)

	// 2. Group HR Spesifik: HANYA BOLEH DIAKSES OLEH HR MANAGER (Create/Update/Delete)
	hrAdmin := api.Group("")

	// Karyawan biasa (EMPLOYEE) akan diblokir (403 Forbidden) jika mengakses endpoint di bawah ini:
	hrAdmin.POST("", CreateEmployeeHandler)
	hrAdmin.PUT("/:id", UpdateEmployeeHandler)
	hrAdmin.DELETE("/:id", DeleteEmployeeHandler)

	hrAdmin.POST("/jobposition", CreateJobPositionHandler)
	hrAdmin.PUT("/jobposition/:id", UpdateJobPositionHandler)
	hrAdmin.DELETE("/jobposition/:id", DeleteJobPositionHandler)

	hrAdmin.POST("/workingschedule", CreateWorkingScheduleHandler)
	hrAdmin.PUT("/workingschedule/:id", UpdateWorkingScheduleHandler)
	hrAdmin.DELETE("/workingschedule/:id", DeleteWorkingScheduleHandler)

	hrAdmin.POST("/skill", CreateSkillHandler)
	hrAdmin.PUT("/skill/:id", UpdateSkillHandler)
	hrAdmin.DELETE("/skill/:id", DeleteSkillHandler)

	hrAdmin.POST("/skilllevel", CreateSkillLevelHandler)
	hrAdmin.PUT("/skilllevel/:id", UpdateSkillLevelHandler)
	hrAdmin.DELETE("/skilllevel/:id", DeleteSkillLevelHandler)

	hrAdmin.POST("/employeeskill", CreateEmployeeSkillHandler)
	hrAdmin.PUT("/employeeskill/:id", UpdateEmployeeSkillHandler)
	hrAdmin.DELETE("/employeeskill/:id", DeleteEmployeeSkillHandler)

	hrAdmin.POST("/resumeline", CreateResumeLineHandler)
	hrAdmin.PUT("/resumeline/:id", UpdateResumeLineHandler)
	hrAdmin.DELETE("/resumeline/:id", DeleteResumeLineHandler)

	// 3. Group Super Rahasia: KONTRAK KERJA (Hanya HR Manager)
	contract := e.Group("/api/hr/employees/contract", middleware.Auth(), middleware.GlobalAutoRBAC())
	contract.POST("", CreateContractHandler)
	contract.GET("", GetAllContractHandler)
	contract.GET("/:id", GetContractByIDHandler)
	contract.PUT("/:id", UpdateContractHandler)
	contract.DELETE("/:id", DeleteContractHandler)
}



