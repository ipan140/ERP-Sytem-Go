package accounting

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/accounting", middleware.Auth(), middleware.GlobalAutoRBAC())
	// Chart of Accounts (COA)
	api.GET("/accounts", GetAllAccountsHandler)
	api.GET("/accounts/:id", GetAccountByIDHandler)
	api.POST("/accounts", CreateAccountHandler)
	api.PUT("/accounts/:id", UpdateAccountHandler)
	api.DELETE("/accounts/:id", DeleteAccountHandler)

	// Kas Masuk & Kas Keluar Sederhana
	api.POST("/cash-transaction", CreateCashTransactionHandler)

	// Laporan Keuangan Standar SAK
	api.GET("/reports/profit-loss", GetProfitLossReportHandler)
	api.GET("/reports/profit-loss/export-excel", ExportProfitLossExcelHandler)
	api.GET("/reports/balance-sheet", GetBalanceSheetReportHandler)
	api.GET("/reports/balance-sheet/export-excel", ExportBalanceSheetExcelHandler)

	// Jurnal Umum
	api.POST("", CreateJournalEntryHandler)
	api.GET("", GetAllJournalEntryHandler)
	api.GET("/:id", GetJournalEntryByIDHandler)
	api.PUT("/:id", UpdateJournalEntryHandler)
	api.DELETE("/:id", DeleteJournalEntryHandler)
	api.POST("/accountreconcilemodel", CreateAccountReconcileModelHandler)
	api.GET("/accountreconcilemodel", GetAllAccountReconcileModelHandler)
	api.GET("/accountreconcilemodel/:id", GetAccountReconcileModelByIDHandler)
	api.PUT("/accountreconcilemodel/:id", UpdateAccountReconcileModelHandler)
	api.DELETE("/accountreconcilemodel/:id", DeleteAccountReconcileModelHandler)

	api.POST("/followuprule", CreateFollowupRuleHandler)
	api.GET("/followuprule", GetAllFollowupRuleHandler)
	api.GET("/followuprule/:id", GetFollowupRuleByIDHandler)
	api.PUT("/followuprule/:id", UpdateFollowupRuleHandler)
	api.DELETE("/followuprule/:id", DeleteFollowupRuleHandler)

	api.POST("/accountlockdate", CreateAccountLockDateHandler)
	api.GET("/accountlockdate", GetAllAccountLockDateHandler)
	api.GET("/accountlockdate/:id", GetAccountLockDateByIDHandler)
	api.PUT("/accountlockdate/:id", UpdateAccountLockDateHandler)
	api.DELETE("/accountlockdate/:id", DeleteAccountLockDateHandler)
	api.POST("/paymentacquirer", CreatePaymentAcquirerHandler)
	api.GET("/paymentacquirer", GetAllPaymentAcquirerHandler)
	api.GET("/paymentacquirer/:id", GetPaymentAcquirerByIDHandler)
	api.PUT("/paymentacquirer/:id", UpdatePaymentAcquirerHandler)
	api.DELETE("/paymentacquirer/:id", DeletePaymentAcquirerHandler)

	api.POST("/paymenttransaction", CreatePaymentTransactionHandler)
	api.GET("/paymenttransaction", GetAllPaymentTransactionHandler)
	api.GET("/paymenttransaction/:id", GetPaymentTransactionByIDHandler)
	api.PUT("/paymenttransaction/:id", UpdatePaymentTransactionHandler)
	api.DELETE("/paymenttransaction/:id", DeletePaymentTransactionHandler)
	api.POST("/accountincoterms", CreateAccountIncotermsHandler)
	api.GET("/accountincoterms", GetAllAccountIncotermsHandler)
	api.GET("/accountincoterms/:id", GetAccountIncotermsByIDHandler)
	api.PUT("/accountincoterms/:id", UpdateAccountIncotermsHandler)
	api.DELETE("/accountincoterms/:id", DeleteAccountIncotermsHandler)

	api.POST("/webhook/midtrans", MidtransWebhookHandler)
}



