package knowledge

import (
	"ERP-System/config"
	"time"
)

type Article struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);default:''" json:"title"`
	Name        string    `gorm:"type:varchar(255);default:''" json:"name"`
	Category    string    `gorm:"type:varchar(100);default:'HR & Kepegawaian'" json:"category"`
	Summary     string    `gorm:"type:text" json:"summary"`
	Content     string    `gorm:"type:text" json:"content"`
	Author      string    `gorm:"type:varchar(100);default:'HR Directorate'" json:"author"`
	ReadingTime string    `gorm:"type:varchar(50);default:'5 Menit'" json:"reading_time"`
	Views       int       `gorm:"default:0" json:"views"`
	Status      string    `gorm:"type:varchar(50);default:'PUBLISHED'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Article) TableName() string {
	return "setting.articles"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Article{})
}
