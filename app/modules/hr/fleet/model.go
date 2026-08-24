package fleet

import (
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type Vehicle struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ModelName    string    `gorm:"type:varchar(255);not null" json:"model_name"`
	LicensePlate string    `gorm:"type:varchar(50);not null;unique" json:"license_plate"`
	EmployeeID   *uint     `json:"employee_id"` // Driver
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	State        string    `gorm:"type:varchar(50);default:'active'" json:"state"`
	CreatedAt    time.Time `json:"created_at"`
}

type VehicleLogContract struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	VehicleID  uint      `json:"vehicle_id"`
	Vehicle *Vehicle `gorm:"foreignKey:VehicleID"` // Auto-added relation
	Cost       float64   `gorm:"type:numeric(15,2);default:0" json:"cost"`
	Expiration time.Time `json:"expiration"`
}

type VehicleLogFuel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	VehicleID uint      `json:"vehicle_id"`
	Vehicle *Vehicle `gorm:"foreignKey:VehicleID"` // Auto-added relation
	Date      time.Time `json:"date"`
	Liters    float64   `gorm:"type:numeric(10,2);default:0" json:"liters"`
	Amount    float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
}

type VehicleLogServices struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	VehicleID   uint      `json:"vehicle_id"`
	Vehicle *Vehicle `gorm:"foreignKey:VehicleID"` // Auto-added relation
	Date        time.Time `json:"date"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Amount      float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
}


func (Vehicle) TableName() string {
	return "hrd.vehicles"
}

func (VehicleLogContract) TableName() string {
	return "hrd.vehicle_log_contracts"
}

func (VehicleLogFuel) TableName() string {
	return "hrd.vehicle_log_fuels"
}

func (VehicleLogServices) TableName() string {
	return "hrd.vehicle_log_serviceses"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Vehicle{}, &VehicleLogContract{}, &VehicleLogFuel{}, &VehicleLogServices{})
}
