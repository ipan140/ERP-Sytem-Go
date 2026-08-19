package events



import "ERP-System/config"

// 5. Event Ticketing & Barcode
type EventTicket struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	EventID    uint   `json:"event_id"`
	CustomerID uint   `json:"customer_id"`
	Barcode    string `gorm:"type:varchar(100);unique" json:"barcode"` // Di-scan saat acara
	IsScanned  bool   `gorm:"default:false" json:"is_scanned"`
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &EventTicket{}) }


type Event struct { ID uint `gorm:"primaryKey"` }

