package purchase

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreatePurchaseOrder godoc
// @Summary Create a new PurchaseOrder
// @Description Create a new PurchaseOrder in the system
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Success 201 {object} PurchaseOrder
// @Param request body PurchaseOrder true "Payload"
// @Router /api/supply_chain/purchase [post]
// @Security BearerAuth
func CreatePurchaseOrderHandler(c echo.Context) error {
	var req CreatePORequest
	if err := c.Bind(&req); err == nil && len(req.Lines) > 0 {
		po, err := CreatePurchaseOrderWithLinesService(req)
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat Purchase Order", err.Error())
		}
		return utils.SendSuccess(c, http.StatusCreated, "Purchase Order berhasil dibuat", po)
	}

	var data PurchaseOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreatePurchaseOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllPurchaseOrder godoc
// @Summary Get all PurchaseOrder
// @Description Retrieve a list of all PurchaseOrder
// @Tags supply_chain-purchase
// @Produce json
// @Success 200 {object} []PurchaseOrder
// @Router /api/supply_chain/purchase [get]
// @Security BearerAuth
func GetAllPurchaseOrderHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllPurchaseOrderService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")
	partnerID, _ := strconv.Atoi(c.QueryParam("partner_id"))

	data, total, err := GetPaginatedPurchaseOrdersService(offset, limit, search, state, uint(partnerID))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve paginated purchase orders", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Purchase orders retrieved successfully", data, meta)
}

// GetPurchaseSummaryHandler godoc
// @Summary Get Purchase KPI summary
// @Description Total spent monthly, to approve count, to receive count, active vendors
// @Tags supply_chain-purchase
// @Produce json
// @Router /api/supply_chain/purchase/summary [get]
// @Security BearerAuth
func GetPurchaseSummaryHandler(c echo.Context) error {
	summary, err := GetPurchaseSummaryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve purchase summary", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Purchase summary retrieved successfully", summary)
}

// ConfirmPurchaseOrderHandler godoc
// @Summary Confirm Purchase Order / Send RFQ
// @Description Transition state from draft/sent to purchase or to_approve (> 50jt)
// @Tags supply_chain-purchase
// @Produce json
// @Router /api/supply_chain/purchase/:id/confirm [post]
// @Security BearerAuth
func ConfirmPurchaseOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	po, err := ConfirmPurchaseOrderService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, err.Error(), err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Status PO berhasil diperbarui", po)
}

// ApprovePurchaseOrderHandler godoc
// @Summary Manager Approval for Purchase Order
// @Description Approve high-value PO (> 50jt) transitioning to purchase state
// @Tags supply_chain-purchase
// @Produce json
// @Router /api/supply_chain/purchase/:id/approve [post]
// @Security BearerAuth
func ApprovePurchaseOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("user_id").(uint)
	po, err := ApprovePurchaseOrderService(uint(id), userID)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, err.Error(), err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "PO berhasil disetujui oleh manajer", po)
}

// ReceivePurchaseOrderProductsHandler godoc
// @Summary Receive products for PO (Three-Way Matching)
// @Description Increment stock qty, create incoming picking & valuation layer
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Router /api/supply_chain/purchase/:id/receive [post]
// @Security BearerAuth
func ReceivePurchaseOrderProductsHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req ReceiveGoodsRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format payload tidak valid", err.Error())
	}
	if err := ReceivePurchaseOrderProductsService(uint(id), req.WarehouseID, req.Items, req.Notes); err != nil {
		return utils.SendError(c, http.StatusBadRequest, err.Error(), err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Penerimaan barang berhasil dicatat ke stok fisik gudang", nil)
}

// GetPurchaseOrderByID godoc
// @Summary Get a PurchaseOrder by ID
// @Description Retrieve a specific PurchaseOrder by its ID
// @Tags supply_chain-purchase
// @Produce json
// @Param id path int true "PurchaseOrder ID"
// @Success 200 {object} PurchaseOrder
// @Router /api/supply_chain/purchase/{id} [get]
// @Security BearerAuth
func GetPurchaseOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdatePurchaseOrder godoc
// @Summary Update a PurchaseOrder
// @Description Update an existing PurchaseOrder
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Param id path int true "PurchaseOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/{id} [put]
// @Security BearerAuth
func UpdatePurchaseOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdatePurchaseOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeletePurchaseOrder godoc
// @Summary Delete a PurchaseOrder
// @Description Delete a PurchaseOrder by ID
// @Tags supply_chain-purchase
// @Produce json
// @Param id path int true "PurchaseOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/{id} [delete]
// @Security BearerAuth
func DeletePurchaseOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePurchaseOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create PurchaseRequisition
// @Description Create a new PurchaseRequisition
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Success 201 {object} PurchaseRequisition
// @Param request body PurchaseRequisition true "Payload"
// @Router /api/supply_chain/purchase/purchaserequisition [post]
// @Security BearerAuth
func CreatePurchaseRequisitionHandler(c echo.Context) error {
	var data PurchaseRequisition
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePurchaseRequisitionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PurchaseRequisition
// @Description Retrieve a list of all PurchaseRequisition
// @Tags supply_chain-purchase
// @Produce json
// @Success 200 {object} PurchaseRequisition
// @Router /api/supply_chain/purchase/purchaserequisition [get]
// @Security BearerAuth
func GetAllPurchaseRequisitionHandler(c echo.Context) error {
	data, err := GetAllPurchaseRequisitionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPurchaseRequisitionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseRequisitionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PurchaseRequisition
// @Description Update an existing PurchaseRequisition
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Param id path int true "PurchaseRequisition ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/purchaserequisition/{id} [put]
// @Security BearerAuth
func UpdatePurchaseRequisitionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseRequisitionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePurchaseRequisitionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PurchaseRequisition
// @Description Delete PurchaseRequisition by ID
// @Tags supply_chain-purchase
// @Produce json
// @Param id path int true "PurchaseRequisition ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/purchaserequisition/{id} [delete]
// @Security BearerAuth
func DeletePurchaseRequisitionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePurchaseRequisitionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create ProductSupplierInfo
// @Description Create a new ProductSupplierInfo
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Success 201 {object} ProductSupplierInfo
// @Param request body ProductSupplierInfo true "Payload"
// @Router /api/supply_chain/purchase/productsupplierinfo [post]
// @Security BearerAuth
func CreateProductSupplierInfoHandler(c echo.Context) error {
	var data ProductSupplierInfo
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateProductSupplierInfoService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all ProductSupplierInfo
// @Description Retrieve a list of all ProductSupplierInfo
// @Tags supply_chain-purchase
// @Produce json
// @Success 200 {object} ProductSupplierInfo
// @Router /api/supply_chain/purchase/productsupplierinfo [get]
// @Security BearerAuth
func GetAllProductSupplierInfoHandler(c echo.Context) error {
	data, err := GetAllProductSupplierInfoService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetProductSupplierInfoByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductSupplierInfoByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update ProductSupplierInfo
// @Description Update an existing ProductSupplierInfo
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Param id path int true "ProductSupplierInfo ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/productsupplierinfo/{id} [put]
// @Security BearerAuth
func UpdateProductSupplierInfoHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductSupplierInfoByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateProductSupplierInfoService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete ProductSupplierInfo
// @Description Delete ProductSupplierInfo by ID
// @Tags supply_chain-purchase
// @Produce json
// @Param id path int true "ProductSupplierInfo ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/productsupplierinfo/{id} [delete]
// @Security BearerAuth
func DeleteProductSupplierInfoHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProductSupplierInfoService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create PurchaseOrderLine
// @Description Create a new PurchaseOrderLine
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Success 201 {object} PurchaseOrderLine
// @Param request body PurchaseOrderLine true "Payload"
// @Router /api/supply_chain/purchase/purchaseorderline [post]
// @Security BearerAuth
func CreatePurchaseOrderLineHandler(c echo.Context) error {
	var data PurchaseOrderLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePurchaseOrderLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PurchaseOrderLine
// @Description Retrieve a list of all PurchaseOrderLine
// @Tags supply_chain-purchase
// @Produce json
// @Success 200 {object} PurchaseOrderLine
// @Router /api/supply_chain/purchase/purchaseorderline [get]
// @Security BearerAuth
func GetAllPurchaseOrderLineHandler(c echo.Context) error {
	data, err := GetAllPurchaseOrderLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPurchaseOrderLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PurchaseOrderLine
// @Description Update an existing PurchaseOrderLine
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Param id path int true "PurchaseOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/purchaseorderline/{id} [put]
// @Security BearerAuth
func UpdatePurchaseOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseOrderLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePurchaseOrderLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PurchaseOrderLine
// @Description Delete PurchaseOrderLine by ID
// @Tags supply_chain-purchase
// @Produce json
// @Param id path int true "PurchaseOrderLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/purchaseorderline/{id} [delete]
// @Security BearerAuth
func DeletePurchaseOrderLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePurchaseOrderLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// CreateTenderHandler godoc
// @Summary Create a new Purchase Tender (Multi-Vendor RFQ)
// @Description Creates a multi-vendor tender agreement and automatically generates draft RFQs for participating vendors
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Param request body CreateTenderRequest true "Payload"
// @Success 201 {object} utils.SuccessResponse{data=PurchaseRequisition}
// @Router /api/supply_chain/purchase/tender [post]
// @Security BearerAuth
func CreateTenderHandler(c echo.Context) error {
	var req CreateTenderRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, 400, "Invalid payload", err.Error())
	}
	data, err := CreateTenderMultiVendor(req)
	if err != nil {
		return utils.SendError(c, 500, "Failed to create tender", err.Error())
	}
	return utils.SendSuccess(c, 201, "Tender created successfully", data)
}

// SelectTenderWinnerHandler godoc
// @Summary Select winning vendor quotation for purchase tender
// @Description Confirms selected vendor's Purchase Order and cancels competing vendor RFQs
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Param id path int true "Purchase Requisition / Tender ID"
// @Param request body SelectTenderWinnerRequest true "Payload"
// @Success 200 {object} utils.SuccessResponse
// @Router /api/supply_chain/purchase/tender/{id}/select-winner [post]
// @Security BearerAuth
func SelectTenderWinnerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req SelectTenderWinnerRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, 400, "Invalid payload", err.Error())
	}
	if err := SelectTenderWinner(uint(id), req.PurchaseOrderID); err != nil {
		return utils.SendError(c, 500, "Failed to select winner", err.Error())
	}
	return utils.SendSuccess(c, 200, "Tender winner selected successfully", nil)
}


