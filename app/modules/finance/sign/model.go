package sign

import (
	"ERP-System/config"
	"time"
)

type SignatureRequest struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SignatureRequest{})
}


// ---- Auto-Generated TableName methods ----
func (SignatureRequest) TableName() string {
	return "finance.signature_requests"
}