package accounting

import (
	"ERP-System/common/utils"
	"ERP-System/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

// @Summary Create AccountReconcileModel
// @Description Create a new AccountReconcileModel
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting/accountreconcilemodel [post]
// @Security BearerAuth
func CreateAccountReconcileModelHandler(c echo.Context) error {
	var data AccountReconcileModel
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateAccountReconcileModelService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all AccountReconcileModel
// @Description Retrieve a list of all AccountReconcileModel
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountreconcilemodel [get]
// @Security BearerAuth
func GetAllAccountReconcileModelHandler(c echo.Context) error {
	data, err := GetAllAccountReconcileModelService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetAccountReconcileModelByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountReconcileModelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update AccountReconcileModel
// @Description Update an existing AccountReconcileModel
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "AccountReconcileModel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountreconcilemodel/{id} [put]
// @Security BearerAuth
func UpdateAccountReconcileModelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountReconcileModelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateAccountReconcileModelService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete AccountReconcileModel
// @Description Delete AccountReconcileModel by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "AccountReconcileModel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountreconcilemodel/{id} [delete]
// @Security BearerAuth
func DeleteAccountReconcileModelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAccountReconcileModelService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create FollowupRule
// @Description Create a new FollowupRule
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting/followuprule [post]
// @Security BearerAuth
func CreateFollowupRuleHandler(c echo.Context) error {
	var data FollowupRule
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateFollowupRuleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all FollowupRule
// @Description Retrieve a list of all FollowupRule
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/followuprule [get]
// @Security BearerAuth
func GetAllFollowupRuleHandler(c echo.Context) error {
	data, err := GetAllFollowupRuleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetFollowupRuleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFollowupRuleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update FollowupRule
// @Description Update an existing FollowupRule
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "FollowupRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/followuprule/{id} [put]
// @Security BearerAuth
func UpdateFollowupRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFollowupRuleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateFollowupRuleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete FollowupRule
// @Description Delete FollowupRule by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "FollowupRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/followuprule/{id} [delete]
// @Security BearerAuth
func DeleteFollowupRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteFollowupRuleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create AccountLockDate
// @Description Create a new AccountLockDate
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting/accountlockdate [post]
// @Security BearerAuth
func CreateAccountLockDateHandler(c echo.Context) error {
	var data AccountLockDate
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateAccountLockDateService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all AccountLockDate
// @Description Retrieve a list of all AccountLockDate
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountlockdate [get]
// @Security BearerAuth
func GetAllAccountLockDateHandler(c echo.Context) error {
	data, err := GetAllAccountLockDateService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetAccountLockDateByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountLockDateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update AccountLockDate
// @Description Update an existing AccountLockDate
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "AccountLockDate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountlockdate/{id} [put]
// @Security BearerAuth
func UpdateAccountLockDateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountLockDateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateAccountLockDateService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete AccountLockDate
// @Description Delete AccountLockDate by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "AccountLockDate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountlockdate/{id} [delete]
// @Security BearerAuth
func DeleteAccountLockDateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAccountLockDateService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create PaymentAcquirer
// @Description Create a new PaymentAcquirer
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting/paymentacquirer [post]
// @Security BearerAuth
func CreatePaymentAcquirerHandler(c echo.Context) error {
	var data PaymentAcquirer
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePaymentAcquirerService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all PaymentAcquirer
// @Description Retrieve a list of all PaymentAcquirer
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymentacquirer [get]
// @Security BearerAuth
func GetAllPaymentAcquirerHandler(c echo.Context) error {
	data, err := GetAllPaymentAcquirerService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetPaymentAcquirerByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentAcquirerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update PaymentAcquirer
// @Description Update an existing PaymentAcquirer
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "PaymentAcquirer ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymentacquirer/{id} [put]
// @Security BearerAuth
func UpdatePaymentAcquirerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentAcquirerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdatePaymentAcquirerService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete PaymentAcquirer
// @Description Delete PaymentAcquirer by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "PaymentAcquirer ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymentacquirer/{id} [delete]
// @Security BearerAuth
func DeletePaymentAcquirerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePaymentAcquirerService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create PaymentTransaction
// @Description Create a new PaymentTransaction
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting/paymenttransaction [post]
// @Security BearerAuth
func CreatePaymentTransactionHandler(c echo.Context) error {
	var data PaymentTransaction
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePaymentTransactionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all PaymentTransaction
// @Description Retrieve a list of all PaymentTransaction
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymenttransaction [get]
// @Security BearerAuth
func GetAllPaymentTransactionHandler(c echo.Context) error {
	data, err := GetAllPaymentTransactionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetPaymentTransactionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentTransactionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update PaymentTransaction
// @Description Update an existing PaymentTransaction
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "PaymentTransaction ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymenttransaction/{id} [put]
// @Security BearerAuth
func UpdatePaymentTransactionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentTransactionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdatePaymentTransactionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete PaymentTransaction
// @Description Delete PaymentTransaction by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "PaymentTransaction ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymenttransaction/{id} [delete]
// @Security BearerAuth
func DeletePaymentTransactionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePaymentTransactionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create AccountIncoterms
// @Description Create a new AccountIncoterms
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting/accountincoterms [post]
// @Security BearerAuth
func CreateAccountIncotermsHandler(c echo.Context) error {
	var data AccountIncoterms
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateAccountIncotermsService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all AccountIncoterms
// @Description Retrieve a list of all AccountIncoterms
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountincoterms [get]
// @Security BearerAuth
func GetAllAccountIncotermsHandler(c echo.Context) error {
	data, err := GetAllAccountIncotermsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetAccountIncotermsByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountIncotermsByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update AccountIncoterms
// @Description Update an existing AccountIncoterms
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "AccountIncoterms ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountincoterms/{id} [put]
// @Security BearerAuth
func UpdateAccountIncotermsHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountIncotermsByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateAccountIncotermsService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete AccountIncoterms
// @Description Delete AccountIncoterms by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "AccountIncoterms ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountincoterms/{id} [delete]
// @Security BearerAuth
func DeleteAccountIncotermsHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAccountIncotermsService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

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
