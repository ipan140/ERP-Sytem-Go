package subscriptions

import (
	"fmt"
	"time"

	"ERP-System/app/modules/finance/invoicing"
	"ERP-System/config"
)

func CreateSubscriptionService(data *Subscription) error {
	if data.State == "" {
		data.State = "active"
	}
	if data.StartDate.IsZero() {
		data.StartDate = time.Now()
	}
	if data.NextInvoiceDate.IsZero() {
		data.NextInvoiceDate = time.Now().AddDate(0, 1, 0) // Default 1 bulan
	}
	return CreateSubscription(data)
}

func GetAllSubscriptionService() ([]Subscription, error) {
	return GetAllSubscription()
}

func GetPaginatedSubscriptionService(offset int, limit int, search string, state string) ([]Subscription, int64, error) {
	return GetPaginatedSubscriptions(offset, limit, search, state)
}

func GetSubscriptionByIDService(id uint) (*Subscription, error) {
	return GetSubscriptionByID(id)
}

func UpdateSubscriptionService(data *Subscription) error {
	return UpdateSubscription(data)
}

func DeleteSubscriptionService(id uint) error {
	return DeleteSubscription(id)
}

func CreateSubscriptionPlanService(data *SubscriptionPlan) error { return CreateSubscriptionPlan(data) }
func GetAllSubscriptionPlanService() ([]SubscriptionPlan, error) { return GetAllSubscriptionPlan() }
func GetSubscriptionPlanByIDService(id uint) (*SubscriptionPlan, error) {
	return GetSubscriptionPlanByID(id)
}
func UpdateSubscriptionPlanService(data *SubscriptionPlan) error { return UpdateSubscriptionPlan(data) }
func DeleteSubscriptionPlanService(id uint) error                { return DeleteSubscriptionPlan(id) }

// GenerateSubscriptionInvoiceService: One-Click Recurring Invoicing ke modul Finance
func GenerateSubscriptionInvoiceService(subID uint) (*invoicing.Invoice, error) {
	sub, err := GetSubscriptionByID(subID)
	if err != nil {
		return nil, err
	}

	untaxed := sub.RecurringTotal
	tax := untaxed * 0.11
	total := untaxed + tax

	inv := invoicing.Invoice{
		Name:           fmt.Sprintf("INV/SUB/%s/%04d", time.Now().Format("200601"), sub.ID),
		PartnerID:      sub.PartnerID,
		InvoiceDate:    time.Now(),
		DueDate:        time.Now().AddDate(0, 0, 14), // Net 14
		State:          "draft",
		AmountUntaxed:  untaxed,
		AmountTax:      tax,
		AmountTotal:    total,
		ResidualAmount: total,
		CreatedAt:      time.Now(),
	}

	if err := config.DB.Create(&inv).Error; err != nil {
		return nil, err
	}

	// Buat baris faktur
	invLine := invoicing.InvoiceLine{
		InvoiceID:   inv.ID,
		Description: fmt.Sprintf("Tagihan Langganan: %s (Periode %s)", sub.Name, time.Now().Format("Jan 2006")),
		Quantity:    1,
		UnitPrice:   untaxed,
		SubTotal:    untaxed,
	}
	_ = config.DB.Create(&invLine)

	// Majukan tanggal tagihan berikutnya 1 bulan
	sub.NextInvoiceDate = sub.NextInvoiceDate.AddDate(0, 1, 0)
	_ = UpdateSubscription(sub)

	return &inv, nil
}

