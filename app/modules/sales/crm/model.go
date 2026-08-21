package crm

import (
	"ERP-System/config"
	"time"
)

type SalesTeam struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	ManagerID       *uint     `json:"manager_id"` // HR Employee ID
	InvoicingTarget float64   `gorm:"type:numeric(15,2);default:0" json:"invoicing_target"`
	CreatedAt       time.Time `json:"created_at"`
}

type Stage struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"` // New, Meeting, Proposal, Won, Lost
	Sequence int    `gorm:"default:10" json:"sequence"`
}

type Activity struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LeadID    uint      `json:"lead_id"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"` // Call, Email, Meeting
	Summary   string    `gorm:"type:varchar(255)" json:"summary"`
	Deadline  time.Time `json:"deadline"`
	IsDone    bool      `gorm:"default:false" json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

type Lead struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"` // Opportunity name
	PartnerID       *uint     `json:"partner_id"`                             // Customer ID
	Email           string    `gorm:"type:varchar(100)" json:"email"`
	Phone           string    `gorm:"type:varchar(50)" json:"phone"`
	ExpectedRevenue float64   `gorm:"type:numeric(15,2);default:0" json:"expected_revenue"`
	Probability     float64   `gorm:"type:numeric(5,2);default:10" json:"probability"` // Percentage
	StageID         uint      `json:"stage_id"`
	SalesTeamID     *uint     `json:"sales_team_id"`
	SalespersonID   *uint     `json:"salesperson_id"` // Link to Employee (HR)
	CreatedAt       time.Time `json:"created_at"`
}

type SalesCommission struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	SalespersonID uint      `json:"salesperson_id"`
	Date          time.Time `json:"date"`
	Amount        float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
	State         string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, paid
	CreatedAt     time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SalesTeam{}, &Stage{}, &Activity{}, &Lead{}, &SalesCommission{})
}
