package approvals

import (
	"ERP-System/config"
	"time"
)

type ApprovalRequest struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Type          string    `gorm:"type:varchar(50);default:'Pengeluaran Dana'" json:"type"` // Pengeluaran Dana, Pengadaan, Kontrak
	Amount        float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
	RequesterName string    `gorm:"type:varchar(100)" json:"requester_name"`
	Stage         string    `gorm:"type:varchar(50);default:'Manager Dept'" json:"stage"` // Manager Dept, Finance Head, CFO / Direktur
	Status        string    `gorm:"type:varchar(30);default:'pending'" json:"status"`     // pending, approved, rejected
	ApproverName  string    `gorm:"type:varchar(100)" json:"approver_name"`
	Notes         string    `gorm:"type:text" json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ApprovalRequest{})
}


// ---- Auto-Generated TableName methods ----
func (ApprovalRequest) TableName() string {
	return "finance.approval_requests"
}