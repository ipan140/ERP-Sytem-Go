package crm

import (
	"ERP-System/config"
	"time"
)

type Lead struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"` // Opportunity name
	Email           string    `gorm:"type:varchar(100)" json:"email"`
	Phone           string    `gorm:"type:varchar(50)" json:"phone"`
	ExpectedRevenue float64   `gorm:"type:numeric(15,2);default:0" json:"expected_revenue"`
	Probability     float64   `gorm:"type:numeric(5,2);default:10" json:"probability"` // Percentage
	Stage           string    `gorm:"type:varchar(50);default:'new'" json:"stage"`     // new, qualified, won, lost
	SalespersonID   *uint     `json:"salesperson_id"`                                  // Link to Employee (HR)
	CreatedAt       time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Lead{})
}
