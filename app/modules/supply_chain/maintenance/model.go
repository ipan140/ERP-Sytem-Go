package maintenance

import (
	"ERP-System/config"
	"time"
)

type MaintenanceEquipment struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Category       string    `gorm:"type:varchar(100)" json:"category"`
	WorkcenterID   *uint     `json:"workcenter_id"` // Terhubung ke mesin pabrik (MRP)
	Cost           float64   `gorm:"type:numeric(15,2);default:0" json:"cost"`
	NextActionDate time.Time `json:"next_action_date"` // Jadwal servis berikutnya
	CreatedAt      time.Time `json:"created_at"`
}

type MaintenanceRequest struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"` // Mesin Potong Rusak
	EquipmentID  uint      `json:"equipment_id"`
	Type         string    `gorm:"type:varchar(50);default:'corrective'" json:"type"` // corrective (rusak), preventive (servis rutin)
	State        string    `gorm:"type:varchar(50);default:'todo'" json:"state"` // todo, progress, done, cancel
	ScheduleDate time.Time `json:"schedule_date"`
	Duration     float64   `gorm:"type:numeric(5,2);default:0" json:"duration"` // Lama perbaikan dalam jam
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MaintenanceEquipment{}, &MaintenanceRequest{})
}
