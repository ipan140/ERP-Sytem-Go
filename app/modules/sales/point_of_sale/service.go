package point_of_sale

import (
	"fmt"
	"time"

	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"gorm.io/gorm"
)

func CreatePosSessionService(data *PosSession) error {
	return CreatePosSession(data)
}

func GetAllPosSessionService() ([]PosSession, error) {
	return GetAllPosSession()
}

func GetPaginatedPosSessionService(offset int, limit int, search string, state string) ([]PosSession, int64, error) {
	return GetPaginatedPosSessions(offset, limit, search, state)
}

func GetPosSessionByIDService(id uint) (*PosSession, error) {
	return GetPosSessionByID(id)
}

func UpdatePosSessionService(data *PosSession) error {
	return UpdatePosSession(data)
}

func DeletePosSessionService(id uint) error {
	return DeletePosSession(id)
}

func CreatePosConfigService(data *PosConfig) error        { return CreatePosConfig(data) }
func GetAllPosConfigService() ([]PosConfig, error)        { return GetAllPosConfig() }
func GetPosConfigByIDService(id uint) (*PosConfig, error) { return GetPosConfigByID(id) }
func UpdatePosConfigService(data *PosConfig) error        { return UpdatePosConfig(data) }
func DeletePosConfigService(id uint) error                { return DeletePosConfig(id) }

func CreatePosOrderService(data *PosOrder) error        { return CreatePosOrder(data) }
func GetAllPosOrderService() ([]PosOrder, error)        { return GetAllPosOrder() }
func GetPaginatedPosOrderService(offset int, limit int, search string, state string) ([]PosOrder, int64, error) {
	return GetPaginatedPosOrders(offset, limit, search, state)
}
func GetPosOrderByIDService(id uint) (*PosOrder, error) { return GetPosOrderByID(id) }
func UpdatePosOrderService(data *PosOrder) error        { return UpdatePosOrder(data) }
func DeletePosOrderService(id uint) error               { return DeletePosOrder(id) }

func CreatePosOrderLineService(data *PosOrderLine) error        { return CreatePosOrderLine(data) }
func GetAllPosOrderLineService() ([]PosOrderLine, error)        { return GetAllPosOrderLine() }
func GetPosOrderLineByIDService(id uint) (*PosOrderLine, error) { return GetPosOrderLineByID(id) }
func UpdatePosOrderLineService(data *PosOrderLine) error        { return UpdatePosOrderLine(data) }
func DeletePosOrderLineService(id uint) error                   { return DeletePosOrderLine(id) }

func CreatePosPaymentService(data *PosPayment) error        { return CreatePosPayment(data) }
func GetAllPosPaymentService() ([]PosPayment, error)        { return GetAllPosPayment() }
func GetPosPaymentByIDService(id uint) (*PosPayment, error) { return GetPosPaymentByID(id) }
func UpdatePosPaymentService(data *PosPayment) error        { return UpdatePosPayment(data) }
func DeletePosPaymentService(id uint) error                 { return DeletePosPayment(id) }

func CreateLoyaltyProgramService(data *LoyaltyProgram) error        { return CreateLoyaltyProgram(data) }
func GetAllLoyaltyProgramService() ([]LoyaltyProgram, error)        { return GetAllLoyaltyProgram() }
func GetLoyaltyProgramByIDService(id uint) (*LoyaltyProgram, error) { return GetLoyaltyProgramByID(id) }
func UpdateLoyaltyProgramService(data *LoyaltyProgram) error        { return UpdateLoyaltyProgram(data) }
func DeleteLoyaltyProgramService(id uint) error                     { return DeleteLoyaltyProgram(id) }

type PosCheckoutItem struct {
	ProductID uint    `json:"product_id"`
	Qty       float64 `json:"qty"`
	PriceUnit float64 `json:"price_unit"`
	SubTotal  float64 `json:"sub_total"`
}

type PosCheckoutPayload struct {
	ReceiptNumber string            `json:"receipt_number"`
	PaymentMethod string            `json:"payment_method"` // Cash, QRIS, Card
	CashTendered  float64           `json:"cash_tendered"`
	ChangeAmount  float64           `json:"change_amount"`
	TotalAmount   float64           `json:"total_amount"`
	Items         []PosCheckoutItem `json:"items"`
}

// CheckoutPosOrderService memproses transaksi kasir POS, mencatat order & baris, serta memotong stok gudang
func CheckoutPosOrderService(payload *PosCheckoutPayload) (*PosOrder, error) {
	if payload.ReceiptNumber == "" {
		payload.ReceiptNumber = fmt.Sprintf("POS/%s/%05d", time.Now().Format("20060102"), time.Now().Unix()%100000)
	}

	order := PosOrder{
		Name:      payload.ReceiptNumber,
		Total:     payload.TotalAmount,
		State:     "paid",
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	for _, it := range payload.Items {
		line := PosOrderLine{
			OrderID:   order.ID,
			ProductID: it.ProductID,
			Qty:       it.Qty,
			PriceUnit: it.PriceUnit,
			SubTotal:  it.SubTotal,
		}
		_ = config.DB.Create(&line)

		// Kurangi stok produk jika product_id valid
		if it.ProductID > 0 {
			_ = config.DB.Model(&inventory.Product{}).
				Where("id = ?", it.ProductID).
				UpdateColumn("stock_qty", gorm.Expr("stock_qty - ?", it.Qty))
		}
	}

	// Catat PosPayment
	payment := PosPayment{
		OrderID: order.ID,
		Method:  payload.PaymentMethod,
		Amount:  payload.TotalAmount,
	}
	_ = config.DB.Create(&payment)

	return &order, nil
}

