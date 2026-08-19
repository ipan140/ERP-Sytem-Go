package fleet

import (
	"ERP-System/config"
	"time"
)

type Vehicle struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ModelName    string    `gorm:"type:varchar(255);not null" json:"model_name"`
	LicensePlate string    `gorm:"type:varchar(50);not null;unique" json:"license_plate"`
	EmployeeID   *uint     `json:"employee_id"` // Driver
	State        string    `gorm:"type:varchar(50);default:'active'" json:"state"`
	CreatedAt    time.Time `json:"created_at"`
}

type VehicleLogContract struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	VehicleID  uint      `json:"vehicle_id"`
	Cost       float64   `gorm:"type:numeric(15,2);default:0" json:"cost"`
	Expiration time.Time `json:"expiration"`
}

type VehicleLogFuel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	VehicleID uint      `json:"vehicle_id"`
	Date      time.Time `json:"date"`
	Liters    float64   `gorm:"type:numeric(10,2);default:0" json:"liters"`
	Amount    float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
}

type VehicleLogServices struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	VehicleID   uint      `json:"vehicle_id"`
	Date        time.Time `json:"date"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Amount      float64   `gorm:"type:numeric(15,2);default:0" json:"amount"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Vehicle{}, &VehicleLogContract{}, &VehicleLogFuel{}, &VehicleLogServices{})
}
