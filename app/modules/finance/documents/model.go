package documents

import (
	"ERP-System/app/modules/core/base"
	"ERP-System/config"
	"time"
)

type FinanceDocument struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PartnerID uint `json:"partner_id"`
	Partner *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &FinanceDocument{})
}


// ---- Auto-Generated TableName methods ----
func (FinanceDocument) TableName() string {
	return "finance.finance_documents"
}