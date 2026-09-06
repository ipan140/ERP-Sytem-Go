package mass_mailing

import (
	"time"

	"ERP-System/config"
)

// 4. Mass Mailing & Tracking
type MailingCampaign struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255)" json:"name"`
	Subject        string    `gorm:"type:varchar(255)" json:"subject"`
	TargetAudience string    `gorm:"type:varchar(255)" json:"target_audience"`
	Status         string     `gorm:"type:varchar(50);default:'Draft'" json:"status"` // Draft, Scheduled, In-Queue, Sent
	SentCount      int        `gorm:"default:0" json:"sent_count"`
	OpenedCount    int        `gorm:"default:0" json:"opened_count"`
	ClickedCount   int        `gorm:"default:0" json:"clicked_count"`
	BouncedCount   int        `gorm:"default:0" json:"bounced_count"`
	ScheduledAt    *time.Time `json:"scheduled_at"`

	// Enterprise Fields: Budgeting & Approval
	BudgetAllocated float64    `gorm:"type:numeric(15,2);default:0" json:"budget_allocated"`
	ActualSpend     float64    `gorm:"type:numeric(15,2);default:0" json:"actual_spend"`
	ApprovalStatus  string     `gorm:"type:varchar(50);default:'Draft'" json:"approval_status"` // Draft, Waiting Approval, Approved, Rejected
	ApprovedByID    *uint      `json:"approved_by_id"`
	ApprovedAt      *time.Time `json:"approved_at"`
	RejectReason    *string    `gorm:"type:text" json:"reject_reason"`

	// Enterprise A/B Split Testing
	IsABTesting     bool       `gorm:"default:false" json:"is_ab_testing"`
	SubjectB        string     `gorm:"type:varchar(255)" json:"subject_b"`
	SampleSizePct   int        `gorm:"default:20" json:"sample_size_pct"` // Misal 20% dari audiens
	WinnerMetric    string     `gorm:"type:varchar(50);default:'open_rate'" json:"winner_metric"` // open_rate, click_rate
	WinnerVariant   string     `gorm:"type:varchar(10);default:''" json:"winner_variant"` // A, B, atau empty
	VariantAOpened  int        `gorm:"default:0" json:"variant_a_opened"`
	VariantBOpened  int        `gorm:"default:0" json:"variant_b_opened"`
}

// 10. UTM Link Tracker
type UtmTracker struct {
	ID               uint    `gorm:"primaryKey" json:"id"`
	CampaignID       uint    `json:"campaign_id"`
	Campaign *MailingCampaign `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"` // Odoo relation mapped
	UtmSource        string  `gorm:"type:varchar(50)" json:"utm_source"`          // misal: facebook
	UtmMedium        string  `gorm:"type:varchar(50)" json:"utm_medium"`          // misal: cpc
	GeneratedRevenue float64 `gorm:"type:numeric(15,2)" json:"generated_revenue"` // Untung dari link ini
}


func (MailingCampaign) TableName() string {
	return "marketing.mailing_campaigns"
}

func (UtmTracker) TableName() string {
	return "marketing.utm_trackers"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MailingCampaign{}, &UtmTracker{})
}
