package sign

import (
	"ERP-System/config"
	"time"
)

type SignatureRequest struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Name          string     `gorm:"type:varchar(255);not null" json:"name"`
	DocumentTitle string     `gorm:"type:varchar(255)" json:"document_title"`
	SignerName    string     `gorm:"type:varchar(100)" json:"signer_name"`
	SignerRole    string     `gorm:"type:varchar(100)" json:"signer_role"`
	Status        string     `gorm:"type:varchar(30);default:'pending'" json:"status"` // pending, signed
	SignedAt      *time.Time `json:"signed_at"`
	SignatureHash string     `gorm:"type:varchar(255)" json:"signature_hash"`
	CreatedAt     time.Time  `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SignatureRequest{})
}


// ---- Auto-Generated TableName methods ----
func (SignatureRequest) TableName() string {
	return "finance.signature_requests"
}