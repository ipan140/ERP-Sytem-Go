package sales_core

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSaleOrder godoc
// @Summary Create a new SaleOrder
// @Description Create a new SaleOrder in the system
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} SaleOrder
// @Param request body SaleOrder true "Payload"
// @Router /api/sales/sales_core [post]
// @Security BearerAuth
func CreateSaleOrderHandler(c echo.Context) error {
	var data SaleOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSaleOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSaleOrder godoc
// @Summary Get all SaleOrder
// @Description Retrieve a list of all SaleOrder
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} []SaleOrder
// @Router /api/sales/sales_core [get]
// @Security BearerAuth
func GetAllSaleOrderHandler(c echo.Context) error {
	data, err := GetAllSaleOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSaleOrderByID godoc
// @Summary Get a SaleOrder by ID
// @Description Retrieve a specific SaleOrder by its ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} SaleOrder
// @Router /api/sales/sales_core/{id} [get]
// @Security BearerAuth
func GetSaleOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSaleOrder godoc
// @Summary Update a SaleOrder
// @Description Update an existing SaleOrder
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [put]
// @Security BearerAuth
func UpdateSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSaleOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSaleOrder godoc
// @Summary Delete a SaleOrder
// @Description Delete a SaleOrder by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [delete]
// @Security BearerAuth
func DeleteSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSaleOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// ConfirmSaleOrderHandler godoc
// @Summary Confirm SaleOrder to Sales Order
// @Description Transition quotation status to 'sale' and check warehouse stock
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} SaleOrder
// @Router /api/sales/sales_core/{id}/confirm [post]
// @Security BearerAuth
func ConfirmSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := ConfirmSaleOrderService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to confirm sale order", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Sale order confirmed successfully", order)
}

// UpdateSaleOrderStatusHandler godoc
// @Summary Update SaleOrder state
// @Description Update quotation state (draft, sent, sale, done, cancel)
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Param request body map[string]string true "Status Payload (state)"
// @Success 200 {object} SaleOrder
// @Router /api/sales/sales_core/{id}/status [put]
// @Security BearerAuth
func UpdateSaleOrderStatusHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var body struct {
		State string `json:"state"`
	}
	if err := c.Bind(&body); err != nil || body.State == "" {
		return utils.SendError(c, http.StatusBadRequest, "Invalid status payload", "state is required")
	}
	order, err := UpdateSaleOrderStatusService(uint(id), body.State)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update state", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Status updated successfully", order)
}

// GenerateQuotationDocHandler godoc
// @Summary Trigger generation of quotation PDF / document
// @Description Queues quotation PDF generation and returns order details
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} SaleOrder
// @Router /api/sales/sales_core/{id}/print [get]
// @Security BearerAuth
func GenerateQuotationDocHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := GetSaleOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Sale order not found", err.Error())
	}
	_ = GenerateQuotationPDFService(uint(id), 1)
	return utils.SendSuccess(c, http.StatusOK, "Quotation document ready", order)
}

// CreateInvoiceFromSaleOrderHandler godoc
// @Summary One-Click Invoicing from Sales Order
// @Description Create a customer invoice in Finance Invoicing automatically from confirmed Sales Order
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 201 {object} invoicing.Invoice
// @Router /api/sales/sales_core/{id}/create-invoice [post]
// @Security BearerAuth
func CreateInvoiceFromSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	inv, err := CreateInvoiceFromSaleOrderService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create invoice from sale order", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Invoice created successfully in Finance module", inv)
}

// GeneratePaymentLinkHandler godoc
// @Summary Generate Direct Payment Link
// @Description Generate instant payment link (Midtrans / QRIS / VA) for SaleOrder
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} PaymentLinkResult
// @Router /api/sales/sales_core/{id}/payment-link [post]
// @Security BearerAuth
func GeneratePaymentLinkHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	link, err := GeneratePaymentLinkService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to generate payment link", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Payment link generated successfully", link)
}

// ApproveDiscountHandler godoc
// @Summary Approve or reject high discount on SaleOrder
// @Description Authorize discounts greater than 15% before quotation can be confirmed to SO
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Param request body map[string]string true "Approval Payload (status, approver)"
// @Success 200 {object} SaleOrder
// @Router /api/sales/sales_core/{id}/approve-discount [post]
// @Security BearerAuth
func ApproveDiscountHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var body struct {
		Status   string `json:"status"`   // Approved, Rejected
		Approver string `json:"approver"` // Manager name
	}
	if err := c.Bind(&body); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	order, err := ApproveDiscountSaleOrderService(uint(id), body.Status, body.Approver)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to process discount approval", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Discount approval updated successfully", order)
}

// GetSalesLeaderboardHandler godoc
// @Summary Get Salesperson Performance Leaderboard & Accrued Commissions
// @Description Aggregates total revenue, confirmed deals, KPI target quota achievement, and commission amounts per salesperson
// @Tags sales-sales_core
// @Produce json
// @Success 200 {array} SalesLeaderboardItem
// @Router /api/sales/sales_core/leaderboard [get]
// @Security BearerAuth
func GetSalesLeaderboardHandler(c echo.Context) error {
	board, err := GetSalesLeaderboardService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve sales leaderboard", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Sales leaderboard retrieved successfully", board)
}

// SignSaleOrderHandler godoc
// @Summary Digital E-Signature for Quotation / Sales Order
// @Description Signs a sales quotation or order digitally with the customer/signer's name and signature
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Param request body map[string]string false "Signer payload (signer_name)"
// @Success 200 {object} SaleOrder
// @Router /api/sales/sales_core/{id}/sign [post]
// @Security BearerAuth
func SignSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var body struct {
		SignerName string `json:"signer_name"`
	}
	_ = c.Bind(&body)
	order, err := SignSaleOrderService(uint(id), body.SignerName)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to sign sales order", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Sales order successfully signed", order)
}

// @Summary Create Pricelist
// @Description Create a new Pricelist
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} Pricelist
// @Param request body Pricelist true "Payload"
// @Router /api/sales/sales_core/pricelist [post]
// @Security BearerAuth
func CreatePricelistHandler(c echo.Context) error {
	var data Pricelist
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePricelistService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Pricelist
// @Description Retrieve a list of all Pricelist
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} Pricelist
// @Router /api/sales/sales_core/pricelist [get]
// @Security BearerAuth
func GetAllPricelistHandler(c echo.Context) error {
	data, err := GetAllPricelistService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPricelistByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Pricelist
// @Description Update an existing Pricelist
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "Pricelist ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelist/{id} [put]
// @Security BearerAuth
func UpdatePricelistHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePricelistService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Pricelist
// @Description Delete Pricelist by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "Pricelist ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelist/{id} [delete]
// @Security BearerAuth
func DeletePricelistHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePricelistService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create PricelistItem
// @Description Create a new PricelistItem
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} PricelistItem
// @Param request body PricelistItem true "Payload"
// @Router /api/sales/sales_core/pricelistitem [post]
// @Security BearerAuth
func CreatePricelistItemHandler(c echo.Context) error {
	var data PricelistItem
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePricelistItemService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PricelistItem
// @Description Retrieve a list of all PricelistItem
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} PricelistItem
// @Router /api/sales/sales_core/pricelistitem [get]
// @Security BearerAuth
func GetAllPricelistItemHandler(c echo.Context) error {
	data, err := GetAllPricelistItemService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPricelistItemByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistItemByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PricelistItem
// @Description Update an existing PricelistItem
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "PricelistItem ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelistitem/{id} [put]
// @Security BearerAuth
func UpdatePricelistItemHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPricelistItemByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePricelistItemService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PricelistItem
// @Description Delete PricelistItem by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "PricelistItem ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/pricelistitem/{id} [delete]
// @Security BearerAuth
func DeletePricelistItemHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePricelistItemService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create QuotationTemplate
// @Description Create a new QuotationTemplate
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} QuotationTemplate
// @Param request body QuotationTemplate true "Payload"
// @Router /api/sales/sales_core/quotationtemplate [post]
// @Security BearerAuth
func CreateQuotationTemplateHandler(c echo.Context) error {
	var data QuotationTemplate
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateQuotationTemplateService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all QuotationTemplate
// @Description Retrieve a list of all QuotationTemplate
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} QuotationTemplate
// @Router /api/sales/sales_core/quotationtemplate [get]
// @Security BearerAuth
func GetAllQuotationTemplateHandler(c echo.Context) error {
	data, err := GetAllQuotationTemplateService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetQuotationTemplateByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQuotationTemplateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update QuotationTemplate
// @Description Update an existing QuotationTemplate
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "QuotationTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/quotationtemplate/{id} [put]
// @Security BearerAuth
func UpdateQuotationTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetQuotationTemplateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateQuotationTemplateService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete QuotationTemplate
// @Description Delete QuotationTemplate by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "QuotationTemplate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/quotationtemplate/{id} [delete]
// @Security BearerAuth
func DeleteQuotationTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteQuotationTemplateService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create DeliveryMethod
// @Description Create a new DeliveryMethod
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} DeliveryMethod
// @Param request body DeliveryMethod true "Payload"
// @Router /api/sales/sales_core/deliverymethod [post]
// @Security BearerAuth
func CreateDeliveryMethodHandler(c echo.Context) error {
	var data DeliveryMethod
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateDeliveryMethodService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all DeliveryMethod
// @Description Retrieve a list of all DeliveryMethod
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} DeliveryMethod
// @Router /api/sales/sales_core/deliverymethod [get]
// @Security BearerAuth
func GetAllDeliveryMethodHandler(c echo.Context) error {
	data, err := GetAllDeliveryMethodService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetDeliveryMethodByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDeliveryMethodByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update DeliveryMethod
// @Description Update an existing DeliveryMethod
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "DeliveryMethod ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/deliverymethod/{id} [put]
// @Security BearerAuth
func UpdateDeliveryMethodHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDeliveryMethodByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateDeliveryMethodService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete DeliveryMethod
// @Description Delete DeliveryMethod by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "DeliveryMethod ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/deliverymethod/{id} [delete]
// @Security BearerAuth
func DeleteDeliveryMethodHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteDeliveryMethodService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create SaleOrderLine
// @Description Create a new SaleOrderLine
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} SaleOrderLine
// @Param request body SaleOrderLine true "Payload"
// @Router /api/sales/sales_core/saleorderline [post]
// @Security BearerAuth
func CreateSaleOrderLineHandler(c echo.Context) error {
	var data SaleOrderLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSaleOrderLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all SaleOrderLine
// @Description Retrieve a list of all SaleOrderLine
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} SaleOrderLine
// @Router /api/sales/sales_core/saleorderline [get]
// @Security BearerAuth
func GetAllSaleOrderLineHandler(c echo.Context) error {
	data, err := GetAllSaleOrderLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSaleOrderLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update SaleOrderLine
// @Description Update an existing SaleOrderLine
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/saleorderline/{id} [put]
// @Security BearerAuth
func UpdateSaleOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSaleOrderLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete SaleOrderLine
// @Description Delete SaleOrderLine by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/saleorderline/{id} [delete]
// @Security BearerAuth
func DeleteSaleOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSaleOrderLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


