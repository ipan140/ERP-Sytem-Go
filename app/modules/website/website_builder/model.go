package website_builder

import (
	"ERP-System/config"
	"time"
)

type Page struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Announcement struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Category  string    `gorm:"type:varchar(100);default:'Corporate'" json:"category"`
	Content   string    `gorm:"type:text" json:"content"`
	Author    string    `gorm:"type:varchar(100);default:'Direksi & HR'" json:"author"`
	IsPinned  bool      `gorm:"default:false" json:"is_pinned"`
	CreatedAt time.Time `json:"created_at"`
}

type WBSReport struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TicketCode  string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"ticket_code"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	Subject     string    `gorm:"type:varchar(255);not null" json:"subject"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Location    string    `gorm:"type:varchar(150)" json:"location"`
	IncidentDate string   `gorm:"type:varchar(50)" json:"incident_date"`
	EvidenceURL string    `gorm:"type:varchar(500)" json:"evidence_url"`
	Status      string    `gorm:"type:varchar(50);default:'Sedang Ditelaah'" json:"status"` // Sedang Ditelaah, Investigasi Berjalan, Selesai Ditindaklanjuti
	Resolution  string    `gorm:"type:text" json:"resolution"`
	CreatedAt   time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Page{}, &Announcement{}, &WBSReport{})
}

// ---- Auto-Generated TableName methods ----
func (Page) TableName() string {
	return "website_portal.pages"
}

func (Announcement) TableName() string {
	return "website_portal.announcements"
}

func (WBSReport) TableName() string {
	return "website_portal.wbs_reports"
}