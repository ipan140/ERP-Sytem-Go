package maintenance

import (
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/config"
	"time"
)

type MaintenanceEquipment struct {
	ID             uint                         `gorm:"primaryKey" json:"id"`
	Name           string                       `gorm:"type:varchar(255);not null" json:"name"`
	Category       string                       `gorm:"type:varchar(100)" json:"category"`
	WorkcenterID   *uint                        `json:"workcenter_id"` // Terhubung ke mesin pabrik (MRP)
	Workcenter     *manufacturing.MrpWorkcenter `gorm:"foreignKey:WorkcenterID" json:"workcenter,omitempty"` // Odoo relation mapped
	Cost           float64                      `gorm:"type:numeric(15,2);default:0" json:"cost"`
	NextActionDate *time.Time                   `json:"next_action_date"` // Jadwal servis berikutnya
	CreatedAt      time.Time                    `json:"created_at"`
}

type MaintenanceRequest struct {
	ID           uint                  `gorm:"primaryKey" json:"id"`
	Name         string                `gorm:"type:varchar(255);not null" json:"name"` // Mesin Potong Rusak / Servis Rutin
	Code         string                `gorm:"type:varchar(50);uniqueIndex" json:"code"` // MR-20260909-001
	EquipmentID  uint                  `json:"equipment_id"`
	Equipment    *MaintenanceEquipment `gorm:"foreignKey:EquipmentID" json:"equipment,omitempty"` // Odoo relation mapped
	Type         string                `gorm:"type:varchar(50);default:'corrective'" json:"type"` // corrective (rusak), preventive (servis rutin)
	Priority     string                `gorm:"type:varchar(20);default:'normal'" json:"priority"` // low, normal, high, urgent
	State        string                `gorm:"type:varchar(50);default:'todo'" json:"state"`      // todo, progress, done, cancel
	ScheduleDate *time.Time            `json:"schedule_date"`
	Duration     float64               `gorm:"type:numeric(5,2);default:0" json:"duration"` // Lama perbaikan dalam jam
	Notes        string                `gorm:"type:text" json:"notes"`
	DateDone     *time.Time            `json:"date_done"`
	CreatedAt    time.Time             `json:"created_at"`
}

func (MaintenanceEquipment) TableName() string {
	return "supply_chain.maintenance_equipments"
}

func (MaintenanceRequest) TableName() string {
	return "supply_chain.maintenance_requests"
}

type MaintenanceSummary struct {
	TotalRequests   int64 `json:"total_requests"`
	TodoCount       int64 `json:"todo_count"`
	InProgressCount int64 `json:"in_progress_count"`
	DoneCount       int64 `json:"done_count"`
	TotalEquipments int64 `json:"total_equipments"`
}

type CreateMaintenanceRequestDto struct {
	Name         string   `json:"name"`
	EquipmentID  uint     `json:"equipment_id"`
	Type         string   `json:"type"`     // corrective, preventive
	Priority     string   `json:"priority"` // low, normal, high, urgent
	ScheduleDate *string  `json:"schedule_date"`
	Duration     *float64 `json:"duration"`
	Notes        string   `json:"notes"`
}

type UpdateMaintenanceStateRequest struct {
	State    string   `json:"state"` // todo, progress, done, cancel
	Duration *float64 `json:"duration"`
	Notes    *string  `json:"notes"`
}

type CreateMaintenanceEquipmentDto struct {
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	WorkcenterID   *uint   `json:"workcenter_id"`
	Cost           float64 `json:"cost"`
	NextActionDate *string `json:"next_action_date"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MaintenanceEquipment{}, &MaintenanceRequest{})
}
