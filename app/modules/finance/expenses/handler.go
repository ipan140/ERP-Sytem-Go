package expenses

import (
	"ERP-System/common/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateExpense godoc
// @Summary Create a new Expense
// @Description Create a new Expense in the system
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Success 201 {object} Expense
// @Param request body Expense true "Payload"
// @Router /api/finance/expenses [post]
// @Security BearerAuth
func CreateExpenseHandler(c echo.Context) error {
	var data Expense
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateExpenseService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllExpense godoc
// @Summary Get all Expense
// @Description Retrieve a list of all Expense
// @Tags finance-expenses
// @Produce json
// @Success 200 {object} []Expense
// @Router /api/finance/expenses [get]
// @Security BearerAuth
func GetAllExpenseHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllExpenseService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	status := c.QueryParam("status")

	data, total, err := GetPaginatedExpenseService(offset, limit, search, status)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetExpenseByID godoc
// @Summary Get a Expense by ID
// @Description Retrieve a specific Expense by its ID
// @Tags finance-expenses
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} Expense
// @Router /api/finance/expenses/{id} [get]
// @Security BearerAuth
func GetExpenseByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateExpense godoc
// @Summary Update a Expense
// @Description Update an existing Expense
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/{id} [put]
// @Security BearerAuth
func UpdateExpenseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateExpenseService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteExpense godoc
// @Summary Delete a Expense
// @Description Delete a Expense by ID
// @Tags finance-expenses
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/{id} [delete]
// @Security BearerAuth
func DeleteExpenseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteExpenseService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create ExpenseSheet
// @Description Create a new ExpenseSheet
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Success 201 {object} ExpenseSheet
// @Param request body ExpenseSheet true "Payload"
// @Router /api/finance/expenses/expensesheet [post]
// @Security BearerAuth
func CreateExpenseSheetHandler(c echo.Context) error {
	var data ExpenseSheet
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateExpenseSheetService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all ExpenseSheet
// @Description Retrieve a list of all ExpenseSheet
// @Tags finance-expenses
// @Produce json
// @Success 200 {object} ExpenseSheet
// @Router /api/finance/expenses/expensesheet [get]
// @Security BearerAuth
func GetAllExpenseSheetHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllExpenseSheetService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	status := c.QueryParam("status")

	data, total, err := GetPaginatedExpenseSheetService(offset, limit, search, status)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Retrieved successfully", data, meta)
}
func GetExpenseSheetByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseSheetByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update ExpenseSheet
// @Description Update an existing ExpenseSheet
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Param id path int true "ExpenseSheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/expensesheet/{id} [put]
// @Security BearerAuth
func UpdateExpenseSheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseSheetByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateExpenseSheetService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete ExpenseSheet
// @Description Delete ExpenseSheet by ID
// @Tags finance-expenses
// @Produce json
// @Param id path int true "ExpenseSheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/expensesheet/{id} [delete]
// @Security BearerAuth
func DeleteExpenseSheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteExpenseSheetService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// --- PETTY CASH (KAS KECIL) HANDLERS ---

// GetPettyCashHandler godoc
// @Summary Ambil saldo kas kecil dan riwayat transaksi
// @Description Mengambil data plafon kas kecil, saldo riil kasir, dan daftar bukti pengeluaran
// @Tags finance-expenses
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/petty-cash [get]
// @Security BearerAuth
func GetPettyCashHandler(c echo.Context) error {
	fund, txs, err := GetOrCreatePettyCashFundService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data kas kecil", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Petty cash retrieved", map[string]interface{}{
		"fund":         fund,
		"transactions": txs,
	})
}

// RecordPettyCashExpenseHandler godoc
// @Summary Catat nota pengeluaran kas kecil
// @Description Mencatat pengeluaran kas kecil dan memverifikasi sisa saldo fisik kasir
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Param request body PettyCashTransaction true "Payload Transaksi"
// @Success 201 {object} PettyCashTransaction
// @Router /api/finance/expenses/petty-cash/expense [post]
// @Security BearerAuth
func RecordPettyCashExpenseHandler(c echo.Context) error {
	var tx PettyCashTransaction
	if err := c.Bind(&tx); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if tx.Amount <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Nominal harus lebih besar dari 0", "")
	}
	if err := RecordPettyCashExpenseService(&tx); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Gagal mencatat kas kecil", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Pengeluaran kas kecil berhasil dicatat", tx)
}

// ReplenishPettyCashHandler godoc
// @Summary Pengisian kembali kas kecil (Imprest Fund Replenishment)
// @Description Mengisi kembali kas kecil ke batas plafon awal dan mencatat jurnal penggantian
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/petty-cash/replenish [post]
// @Security BearerAuth
func ReplenishPettyCashHandler(c echo.Context) error {
	type ReplenishReq struct {
		FundID     uint   `json:"fund_id"`
		RecordedBy string `json:"recorded_by"`
	}
	var req ReplenishReq
	_ = c.Bind(&req)
	if req.FundID == 0 {
		req.FundID = 1
	}
	if req.RecordedBy == "" {
		req.RecordedBy = "Bendahara Kasir"
	}

	amount, err := ReplenishPettyCashService(req.FundID, req.RecordedBy)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Gagal pengisian kas kecil", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, fmt.Sprintf("Penggantian kas kecil berhasil sebesar Rp %.0f!", amount), map[string]interface{}{
		"replenished_amount": amount,
	})
}


