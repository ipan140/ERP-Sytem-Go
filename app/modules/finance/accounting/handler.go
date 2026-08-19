package accounting

import (
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
	"ERP-System/config"
	"strconv"
)

// CreateJournalEntry godoc
// @Summary Create a new JournalEntry
// @Description Create a new JournalEntry in the system
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting [post]
// @Security BearerAuth
func CreateJournalEntryHandler(c echo.Context) error {
	var data JournalEntry
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateJournalEntryService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllJournalEntry godoc
// @Summary Get all JournalEntry
// @Description Retrieve a list of all JournalEntry
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting [get]
// @Security BearerAuth
func GetAllJournalEntryHandler(c echo.Context) error {
	data, err := GetAllJournalEntryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetJournalEntryByID godoc
// @Summary Get a JournalEntry by ID
// @Description Retrieve a specific JournalEntry by its ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [get]
// @Security BearerAuth
func GetJournalEntryByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJournalEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateJournalEntry godoc
// @Summary Update a JournalEntry
// @Description Update an existing JournalEntry
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [put]
// @Security BearerAuth
func UpdateJournalEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJournalEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateJournalEntryService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteJournalEntry godoc
// @Summary Delete a JournalEntry
// @Description Delete a JournalEntry by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [delete]
// @Security BearerAuth
func DeleteJournalEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteJournalEntryService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// GetGeneralLedgerHandler godoc
// @Summary Get General Ledger
// @Description Get grouped debit/credit balance per account
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/ledger [get]
// @Security BearerAuth
func GetGeneralLedgerHandler(c echo.Context) error {
	data, err := GetGeneralLedgerService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menarik laporan buku besar", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Buku Besar berhasil diambil", data)
}

func CreateAccountReconcileModelHandler(c echo.Context) error { var data AccountReconcileModel; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateAccountReconcileModelService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllAccountReconcileModelHandler(c echo.Context) error { data, err := GetAllAccountReconcileModelService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetAccountReconcileModelByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetAccountReconcileModelByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateAccountReconcileModelHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetAccountReconcileModelByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateAccountReconcileModelService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteAccountReconcileModelHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteAccountReconcileModelService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreateFollowupRuleHandler(c echo.Context) error { var data FollowupRule; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateFollowupRuleService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllFollowupRuleHandler(c echo.Context) error { data, err := GetAllFollowupRuleService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetFollowupRuleByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetFollowupRuleByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateFollowupRuleHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetFollowupRuleByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateFollowupRuleService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteFollowupRuleHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteFollowupRuleService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreateAccountLockDateHandler(c echo.Context) error { var data AccountLockDate; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateAccountLockDateService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllAccountLockDateHandler(c echo.Context) error { data, err := GetAllAccountLockDateService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetAccountLockDateByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetAccountLockDateByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateAccountLockDateHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetAccountLockDateByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateAccountLockDateService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteAccountLockDateHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteAccountLockDateService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreatePaymentAcquirerHandler(c echo.Context) error { var data PaymentAcquirer; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreatePaymentAcquirerService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllPaymentAcquirerHandler(c echo.Context) error { data, err := GetAllPaymentAcquirerService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetPaymentAcquirerByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetPaymentAcquirerByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdatePaymentAcquirerHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetPaymentAcquirerByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdatePaymentAcquirerService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeletePaymentAcquirerHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeletePaymentAcquirerService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreatePaymentTransactionHandler(c echo.Context) error { var data PaymentTransaction; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreatePaymentTransactionService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllPaymentTransactionHandler(c echo.Context) error { data, err := GetAllPaymentTransactionService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetPaymentTransactionByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetPaymentTransactionByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdatePaymentTransactionHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetPaymentTransactionByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdatePaymentTransactionService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeletePaymentTransactionHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeletePaymentTransactionService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreateAccountIncotermsHandler(c echo.Context) error { var data AccountIncoterms; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateAccountIncotermsService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllAccountIncotermsHandler(c echo.Context) error { data, err := GetAllAccountIncotermsService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetAccountIncotermsByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetAccountIncotermsByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateAccountIncotermsHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetAccountIncotermsByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateAccountIncotermsService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteAccountIncotermsHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteAccountIncotermsService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

// MidtransWebhookHandler adalah CCTV 24 Jam yang menerima sinyal dari Server Midtrans
func MidtransWebhookHandler(c echo.Context) error {
	var payload map[string]interface{}
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid Payload", err.Error())
	}

	// 1. Ambil Data dari Midtrans
	orderID, ok := payload["order_id"].(string)
	transactionStatus, _ := payload["transaction_status"].(string)

	if ok && (transactionStatus == "settlement" || transactionStatus == "capture") {
		// 2. Cari Transaksi berdasarkan OrderID (Reference)
		var trx PaymentTransaction
		if err := config.DB.Where("reference = ?", orderID).First(&trx).Error; err == nil {
			
			// 3. Ubah status Transaksi jadi DONE
			trx.State = "done"
			config.DB.Save(&trx)

			// 4. OTOMATIS LUNASKAN INVOICE!
			if trx.InvoiceID != nil {
				// Kita langsung eksekusi Query Update ke tabel invoices
				config.DB.Exec("UPDATE invoices SET state = 'paid' WHERE id = ?", *trx.InvoiceID)
				fmt.Printf("[WEBHOOK] Hore! Invoice ID %d telah Otomatis LUNAS dari Midtrans!\n", *trx.InvoiceID)
			}
		}
	}

	// Midtrans mewajibkan kita membalas dengan status 200 OK
	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

