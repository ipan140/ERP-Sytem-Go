package employees

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/employees", middleware.Auth())
	api.POST("", CreateEmployeeHandler)
	api.GET("", GetAllEmployeeHandler)
	api.GET("/:id", GetEmployeeByIDHandler)
	api.PUT("/:id", UpdateEmployeeHandler)
	api.DELETE("/:id", DeleteEmployeeHandler)

	api.POST("/jobposition", CreateJobPositionHandler)
	api.GET("/jobposition", GetAllJobPositionHandler)
	api.GET("/jobposition/:id", GetJobPositionByIDHandler)
	api.PUT("/jobposition/:id", UpdateJobPositionHandler)
	api.DELETE("/jobposition/:id", DeleteJobPositionHandler)

	api.POST("/workingschedule", CreateWorkingScheduleHandler)
	api.GET("/workingschedule", GetAllWorkingScheduleHandler)
	api.GET("/workingschedule/:id", GetWorkingScheduleByIDHandler)
	api.PUT("/workingschedule/:id", UpdateWorkingScheduleHandler)
	api.DELETE("/workingschedule/:id", DeleteWorkingScheduleHandler)

	api.POST("/contract", CreateContractHandler)
	api.GET("/contract", GetAllContractHandler)
	api.GET("/contract/:id", GetContractByIDHandler)
	api.PUT("/contract/:id", UpdateContractHandler)
	api.DELETE("/contract/:id", DeleteContractHandler)

	api.POST("/skill", CreateSkillHandler)
	api.GET("/skill", GetAllSkillHandler)
	api.GET("/skill/:id", GetSkillByIDHandler)
	api.PUT("/skill/:id", UpdateSkillHandler)
	api.DELETE("/skill/:id", DeleteSkillHandler)

	api.POST("/skilllevel", CreateSkillLevelHandler)
	api.GET("/skilllevel", GetAllSkillLevelHandler)
	api.GET("/skilllevel/:id", GetSkillLevelByIDHandler)
	api.PUT("/skilllevel/:id", UpdateSkillLevelHandler)
	api.DELETE("/skilllevel/:id", DeleteSkillLevelHandler)

	api.POST("/employeeskill", CreateEmployeeSkillHandler)
	api.GET("/employeeskill", GetAllEmployeeSkillHandler)
	api.GET("/employeeskill/:id", GetEmployeeSkillByIDHandler)
	api.PUT("/employeeskill/:id", UpdateEmployeeSkillHandler)
	api.DELETE("/employeeskill/:id", DeleteEmployeeSkillHandler)

	api.POST("/resumeline", CreateResumeLineHandler)
	api.GET("/resumeline", GetAllResumeLineHandler)
	api.GET("/resumeline/:id", GetResumeLineByIDHandler)
	api.PUT("/resumeline/:id", UpdateResumeLineHandler)
	api.DELETE("/resumeline/:id", DeleteResumeLineHandler)
}
