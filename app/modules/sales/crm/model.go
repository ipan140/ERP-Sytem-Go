package crm

import (
	"ERP-System/app/modules/auth"
	"ERP-System/app/modules/core/base"
	"ERP-System/config"
	"time"
)

type SalesTeam struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	ManagerID       *uint     `json:"manager_id"` // HR Employee ID
	Manager *SalesTeam `gorm:"foreignKey:ManagerID"` // Auto-added relation
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
	Lead *Lead `gorm:"foreignKey:LeadID"` // Auto-added relation
	Type      string    `gorm:"type:varchar(50);not null" json:"type"` // Call, Email, Meeting
	Summary   string    `gorm:"type:varchar(255)" json:"summary"`
	Deadline  time.Time `json:"deadline"`
	IsDone    bool      `gorm:"default:false" json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

type Lead struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"` // Opportunity name
	PartnerID        *uint     `json:"partner_id"`                             // Customer ID
	Partner *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"` // Cross-module relation
	Email            string    `gorm:"type:varchar(100)" json:"email"`
	Phone            string    `gorm:"type:varchar(50)" json:"phone"`
	ExpectedRevenue  float64   `gorm:"type:numeric(15,2);default:0" json:"expected_revenue"`
	Probability      float64   `gorm:"type:numeric(5,2);default:10" json:"probability"` // Percentage
	StageID          uint      `json:"stage_id"`
	Stage *Stage `gorm:"foreignKey:StageID"` // Auto-added relation
	SalesTeamID      *uint     `json:"sales_team_id"`
	SalesTeam *SalesTeam `gorm:"foreignKey:SalesTeamID"` // Auto-added relation
	SalespersonID    *uint     `json:"salesperson_id"` // Link to Employee (HR)
	Salesperson *auth.User `gorm:"foreignKey:SalespersonID" json:"salesperson,omitempty"` // Odoo relation mapped
	
	// FASE 4: Enterprise Lead Scoring Engine & Affiliate Referral Tracking
	LeadScore        int       `gorm:"default:0" json:"lead_score"`                       // 0 - 100+
	ScoreGrade       string    `gorm:"type:varchar(20);default:'Cold'" json:"score_grade"` // Hot (80+), Warm (40-79), Cold (<40)
	BehaviorPoints   string    `gorm:"type:text" json:"behavior_points"`                  // JSON breakdown log skor aktivitas
	ReferralCode     string    `gorm:"type:varchar(100)" json:"referral_code"`            // Kode afiliasi/mitra yang mereferensikan
	AffiliateName    string    `gorm:"type:varchar(255)" json:"affiliate_name"`           // Nama partner afiliasi
	UtmSource        string    `gorm:"type:varchar(100)" json:"utm_source"`               // google, meta, wa_blast, influencer
	CreatedAt        time.Time `json:"created_at"`
}

type SalesCommission struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	SalespersonID uint      `json:"salesperson_id"`
	Salesperson *auth.User `gorm:"foreignKey:SalespersonID" json:"salesperson,omitempty"` // Odoo relation mapped
	Date          time.Time `json:"date"`
	Amount        float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
	State         string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, paid
	CreatedAt     time.Time `json:"created_at"`
}


func (SalesTeam) TableName() string {
	return "sales.sales_teams"
}

func (Stage) TableName() string {
	return "sales.stages"
}

func (Activity) TableName() string {
	return "sales.activities"
}

func (Lead) TableName() string {
	return "sales.leads"
}

func (SalesCommission) TableName() string {
	return "sales.sales_commissions"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SalesTeam{}, &Stage{}, &Activity{}, &Lead{}, &SalesCommission{})
}
