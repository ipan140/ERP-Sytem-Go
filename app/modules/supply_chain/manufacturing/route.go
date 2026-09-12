package manufacturing

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/manufacturing", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.GET("/summary", GetMrpSummaryHandler)
	api.POST("", CreateMrpProductionHandler)
	api.GET("", GetAllMrpProductionHandler)
	api.GET("/:id", GetMrpProductionByIDHandler)
	api.PUT("/:id", UpdateMrpProductionHandler)
	api.DELETE("/:id", DeleteMrpProductionHandler)

	api.POST("/:id/confirm", ConfirmProductionHandler)
	api.POST("/:id/start", StartProductionHandler)
	api.POST("/:id/finish", FinishProductionHandler)
	api.POST("/:id/cancel", CancelProductionHandler)

	api.POST("/mrpworkcenter", CreateMrpWorkcenterHandler)
	api.GET("/mrpworkcenter", GetAllMrpWorkcenterHandler)
	api.GET("/mrpworkcenter/:id", GetMrpWorkcenterByIDHandler)
	api.PUT("/mrpworkcenter/:id", UpdateMrpWorkcenterHandler)
	api.DELETE("/mrpworkcenter/:id", DeleteMrpWorkcenterHandler)

	api.POST("/mrpbom", CreateMrpBomHandler)
	api.GET("/mrpbom", GetAllMrpBomHandler)
	api.GET("/mrpbom/:id", GetMrpBomByIDHandler)
	api.PUT("/mrpbom/:id", UpdateMrpBomHandler)
	api.DELETE("/mrpbom/:id", DeleteMrpBomHandler)

	api.POST("/mrpbomline", CreateMrpBomLineHandler)
	api.GET("/mrpbomline", GetAllMrpBomLineHandler)
	api.GET("/mrpbomline/:id", GetMrpBomLineByIDHandler)
	api.PUT("/mrpbomline/:id", UpdateMrpBomLineHandler)
	api.DELETE("/mrpbomline/:id", DeleteMrpBomLineHandler)

	api.POST("/mrpbombyproduct", CreateMrpBomByproductHandler)
	api.GET("/mrpbombyproduct", GetAllMrpBomByproductHandler)
	api.GET("/mrpbombyproduct/:id", GetMrpBomByproductByIDHandler)
	api.PUT("/mrpbombyproduct/:id", UpdateMrpBomByproductHandler)
	api.DELETE("/mrpbombyproduct/:id", DeleteMrpBomByproductHandler)

	api.POST("/mrpworkorder", CreateMrpWorkorderHandler)
	api.GET("/mrpworkorder", GetAllMrpWorkorderHandler)
	api.GET("/mrpworkorder/:id", GetMrpWorkorderByIDHandler)
	api.PUT("/mrpworkorder/:id", UpdateMrpWorkorderHandler)
	api.DELETE("/mrpworkorder/:id", DeleteMrpWorkorderHandler)

	// FASE 9: MPS & OEE Routes
	api.POST("/mps", CreateMPSHandler)
	api.GET("/mps", GetMPSHandler)
	api.POST("/mps/:id/generate-mo", GenerateMOFromMPSHandler)
	api.POST("/workorder/log-time", LogOEEHandler)
	api.GET("/workcenter/:id/oee", GetOEEHandler)
}



