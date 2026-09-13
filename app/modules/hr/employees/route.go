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
	api.GET("/department", GetAllDepartmentHandler)
	api.GET("/department/:id", GetDepartmentByIDHandler)
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
	api.GET("/warningletter", GetAllWarningLetterHandler)
	api.GET("/warningletter/:id", GetWarningLetterByIDHandler)
	api.GET("/employeetask", GetAllEmployeeTaskHandler)
	api.GET("/employeetask/:id", GetEmployeeTaskByIDHandler)
	api.GET("/overtime", GetAllOvertimeHandler)
	api.GET("/overtime/:id", GetOvertimeByIDHandler)
	api.GET("/employeeloan", GetAllEmployeeLoanHandler)
	api.GET("/employeeloan/:id", GetEmployeeLoanByIDHandler)
	api.GET("/expense", GetAllExpenseHandler)
	api.GET("/expense/:id", GetExpenseByIDHandler)
	api.GET("/payroll", GetAllPayslipsHandler)
	api.GET("/payroll/:id", GetPayslipByIDHandler)

	// ESS (Employee Self-Service) endpoints untuk karyawan biasa
	api.GET("/me/profile", GetMyProfileHandler)
	api.GET("/me/payslips", GetMyPayslipsHandler)
	api.GET("/me/warning-letters", GetMyWarningLettersHandler)

	// THR routes (Viewable by users, managed by HR Admin)
	api.GET("/thr", GetAllTHRHandler)

	// 2. Group HR Spesifik: HANYA BOLEH DIAKSES OLEH HR MANAGER (Create/Update/Delete)
	hrAdmin := api.Group("")

	// Karyawan biasa (EMPLOYEE) akan diblokir (403 Forbidden) jika mengakses endpoint di bawah ini:
	hrAdmin.POST("", CreateEmployeeHandler)
	hrAdmin.PUT("/:id", UpdateEmployeeHandler)
	hrAdmin.DELETE("/:id", DeleteEmployeeHandler)

	hrAdmin.POST("/department", CreateDepartmentHandler)
	hrAdmin.PUT("/department/:id", UpdateDepartmentHandler)
	hrAdmin.DELETE("/department/:id", DeleteDepartmentHandler)

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
	hrAdmin.POST("/warningletter", CreateWarningLetterHandler)
	hrAdmin.PUT("/warningletter/:id", UpdateWarningLetterHandler)
	hrAdmin.DELETE("/warningletter/:id", DeleteWarningLetterHandler)
	hrAdmin.POST("/employeetask", CreateEmployeeTaskHandler)
	hrAdmin.PUT("/employeetask/:id", UpdateEmployeeTaskHandler)
	hrAdmin.DELETE("/employeetask/:id", DeleteEmployeeTaskHandler)
	hrAdmin.POST("/overtime", CreateOvertimeHandler)
	hrAdmin.PUT("/overtime/:id", UpdateOvertimeHandler)
	hrAdmin.DELETE("/overtime/:id", DeleteOvertimeHandler)
	hrAdmin.POST("/employeeloan", CreateEmployeeLoanHandler)
	hrAdmin.PUT("/employeeloan/:id", UpdateEmployeeLoanHandler)
	hrAdmin.DELETE("/employeeloan/:id", DeleteEmployeeLoanHandler)
	hrAdmin.POST("/expense", CreateExpenseHandler)
	hrAdmin.PUT("/expense/:id", UpdateExpenseHandler)
	hrAdmin.DELETE("/expense/:id", DeleteExpenseHandler)
	hrAdmin.POST("/payroll/generate", GeneratePayrollHandler)
	hrAdmin.DELETE("/payroll/:id", DeletePayslipHandler)
	hrAdmin.PUT("/payroll/:id/pay", PayPayslipHandler)
	hrAdmin.PUT("/payroll/bulk-pay", BulkPayPayslipsHandler)
	hrAdmin.GET("/payroll/bank-transfer-export", ExportBankDisbursementHandler)

	// THR Admin routes
	hrAdmin.POST("/thr/generate", GenerateTHRHandler)
	hrAdmin.PUT("/thr/:id/status", UpdateTHRStatusHandler)

	// 3. Group Super Rahasia: KONTRAK KERJA (Hanya HR Manager)
	contract := e.Group("/api/hr/employees/contract", middleware.Auth(), middleware.GlobalAutoRBAC())
	contract.POST("", CreateContractHandler)
	contract.GET("", GetAllContractHandler)
	contract.GET("/:id", GetContractByIDHandler)
	contract.PUT("/:id", UpdateContractHandler)
	contract.DELETE("/:id", DeleteContractHandler)
}



