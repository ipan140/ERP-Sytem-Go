package voip

import (
	"ERP-System/config"
	"time"
)

type VoipExtension struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Ext       string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"ext"`
	User      string    `gorm:"type:varchar(255);not null" json:"user"`
	Dept      string    `gorm:"type:varchar(255);default:''" json:"dept"`
	Device    string    `gorm:"type:varchar(255);default:'WebRTC Softphone'" json:"device"`
	IP        string    `gorm:"type:varchar(50);default:'192.168.10.100'" json:"ip"`
	Status    string    `gorm:"type:varchar(50);default:'AVAILABLE'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CallRecord struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Caller    string    `gorm:"type:varchar(50);default:''" json:"caller"`
	Callee    string    `gorm:"type:varchar(50);default:''" json:"callee"`
	Duration  int       `gorm:"default:0" json:"duration"`
	Status    string    `gorm:"type:varchar(50);default:'ANSWERED'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (VoipExtension) TableName() string {
	return "setting.voip_extensions"
}

func (CallRecord) TableName() string {
	return "setting.call_records"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &VoipExtension{}, &CallRecord{})
}
