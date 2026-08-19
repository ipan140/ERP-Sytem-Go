package subscriptions

import (
	"ERP-System/config"
	"time"
)

type SubscriptionPlan struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Monthly Basic
	BillingPeriod string `gorm:"type:varchar(50);not null" json:"billing_period"` // monthly, yearly
}

type Subscription struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"`
	PartnerID       uint      `json:"partner_id"`
	PlanID          uint      `json:"plan_id"`
	StartDate       time.Time `json:"start_date"`
	NextInvoiceDate time.Time `json:"next_invoice_date"`
	RecurringTotal  float64   `gorm:"type:numeric(15,2);default:0" json:"recurring_total"`
	State           string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, open, closed
	CreatedAt       time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SubscriptionPlan{}, &Subscription{})
}
