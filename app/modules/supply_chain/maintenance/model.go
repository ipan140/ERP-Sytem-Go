package maintenance

import (
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/config"
	"time"
)

type MaintenanceEquipment struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Category       string    `gorm:"type:varchar(100)" json:"category"`
	WorkcenterID   *uint     `json:"workcenter_id"` // Terhubung ke mesin pabrik (MRP)
	Workcenter *manufacturing.MrpWorkcenter `gorm:"foreignKey:WorkcenterID" json:"workcenter,omitempty"` // Odoo relation mapped
	Cost           float64   `gorm:"type:numeric(15,2);default:0" json:"cost"`
	NextActionDate time.Time `json:"next_action_date"` // Jadwal servis berikutnya
	CreatedAt      time.Time `json:"created_at"`
}

type MaintenanceRequest struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"` // Mesin Potong Rusak
	EquipmentID  uint      `json:"equipment_id"`
	Equipment *MaintenanceEquipment `gorm:"foreignKey:EquipmentID" json:"equipment,omitempty"` // Odoo relation mapped
	Type         string    `gorm:"type:varchar(50);default:'corrective'" json:"type"` // corrective (rusak), preventive (servis rutin)
	State        string    `gorm:"type:varchar(50);default:'todo'" json:"state"`      // todo, progress, done, cancel
	ScheduleDate time.Time `json:"schedule_date"`
	Duration     float64   `gorm:"type:numeric(5,2);default:0" json:"duration"` // Lama perbaikan dalam jam
	CreatedAt    time.Time `json:"created_at"`
}


func (MaintenanceEquipment) TableName() string {
	return "supply_chain.maintenance_equipments"
}

func (MaintenanceRequest) TableName() string {
	return "supply_chain.maintenance_requests"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MaintenanceEquipment{}, &MaintenanceRequest{})
}
