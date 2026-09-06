package rental

import (
	"fmt"
	"math"
	"time"
)

func CreateRentalOrderService(data *RentalOrder) error {
	if data.Name == "" {
		data.Name = fmt.Sprintf("RO/%s/%04d", time.Now().Format("2006"), time.Now().Unix()%10000)
	}
	if data.State == "" {
		data.State = "reserved"
	}
	if data.PickupDate.IsZero() {
		data.PickupDate = time.Now()
	}
	if data.ReturnDate.IsZero() {
		data.ReturnDate = time.Now().AddDate(0, 0, 7) // Default sewa 7 hari
	}
	return CreateRentalOrder(data)
}

func GetAllRentalOrderService() ([]RentalOrder, error) {
	return GetAllRentalOrder()
}

func GetRentalOrderByIDService(id uint) (*RentalOrder, error) {
	return GetRentalOrderByID(id)
}

func UpdateRentalOrderService(data *RentalOrder) error {
	return UpdateRentalOrder(data)
}

func DeleteRentalOrderService(id uint) error {
	return DeleteRentalOrder(id)
}

// PickupRentalOrderService: Serah terima barang fisik kepada penyewa
func PickupRentalOrderService(id uint) (*RentalOrder, error) {
	order, err := GetRentalOrderByID(id)
	if err != nil {
		return nil, err
	}
	order.State = "pickedup"
	order.PickupDate = time.Now()
	if err := UpdateRentalOrder(order); err != nil {
		return nil, err
	}
	return order, nil
}

type RentalReturnResult struct {
	Order          *RentalOrder `json:"order"`
	IsLate         bool         `json:"is_late"`
	LateDays       int          `json:"late_days"`
	LateFeePerDay  float64      `json:"late_fee_per_day"`
	TotalLateFee   float64      `json:"total_late_fee"`
	FinalAmountDue float64      `json:"final_amount_due"`
}

// ReturnRentalOrderService: Pengembalian fisik barang dan kalkulasi denda keterlambatan
func ReturnRentalOrderService(id uint) (*RentalReturnResult, error) {
	order, err := GetRentalOrderByID(id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	order.State = "returned"

	lateDays := 0
	totalLateFee := 0.0
	isLate := false
	lateFeePerDay := 150000.0 // Standar denda keterlambatan Rp 150.000 / hari

	if now.After(order.ReturnDate) {
		diff := now.Sub(order.ReturnDate)
		lateDays = int(math.Ceil(diff.Hours() / 24.0))
		if lateDays > 0 {
			isLate = true
			totalLateFee = float64(lateDays) * lateFeePerDay
		}
	}

	order.Total += totalLateFee
	if err := UpdateRentalOrder(order); err != nil {
		return nil, err
	}

	return &RentalReturnResult{
		Order:          order,
		IsLate:         isLate,
		LateDays:       lateDays,
		LateFeePerDay:  lateFeePerDay,
		TotalLateFee:   totalLateFee,
		FinalAmountDue: order.Total,
	}, nil
}

func CreateRentalOrderLineService(data *RentalOrderLine) error { return CreateRentalOrderLine(data) }
func GetAllRentalOrderLineService() ([]RentalOrderLine, error) { return GetAllRentalOrderLine() }
func GetRentalOrderLineByIDService(id uint) (*RentalOrderLine, error) {
	return GetRentalOrderLineByID(id)
}
func UpdateRentalOrderLineService(data *RentalOrderLine) error { return UpdateRentalOrderLine(data) }
func DeleteRentalOrderLineService(id uint) error               { return DeleteRentalOrderLine(id) }

