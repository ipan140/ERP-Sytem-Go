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

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Page{})
}


// ---- Auto-Generated TableName methods ----
func (Page) TableName() string {
	return "website_portal.pages"
}