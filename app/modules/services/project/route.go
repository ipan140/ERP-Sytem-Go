package project

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/project", middleware.Auth())
	api.POST("", CreateProjectHandler)
	api.GET("", GetAllProjectHandler)
	api.GET("/:id", GetProjectByIDHandler)
	api.PUT("/:id", UpdateProjectHandler)
	api.DELETE("/:id", DeleteProjectHandler)

	api.POST("/projectmilestone", CreateProjectMilestoneHandler)
	api.GET("/projectmilestone", GetAllProjectMilestoneHandler)
	api.GET("/projectmilestone/:id", GetProjectMilestoneByIDHandler)
	api.PUT("/projectmilestone/:id", UpdateProjectMilestoneHandler)
	api.DELETE("/projectmilestone/:id", DeleteProjectMilestoneHandler)
	api.POST("/taskdependency", CreateTaskDependencyHandler)
	api.GET("/taskdependency", GetAllTaskDependencyHandler)
	api.GET("/taskdependency/:id", GetTaskDependencyByIDHandler)
	api.PUT("/taskdependency/:id", UpdateTaskDependencyHandler)
	api.DELETE("/taskdependency/:id", DeleteTaskDependencyHandler)
	api.POST("/resourceforecast", CreateResourceForecastHandler)
	api.GET("/resourceforecast", GetAllResourceForecastHandler)
	api.GET("/resourceforecast/:id", GetResourceForecastByIDHandler)
	api.PUT("/resourceforecast/:id", UpdateResourceForecastHandler)
	api.DELETE("/resourceforecast/:id", DeleteResourceForecastHandler)

}
