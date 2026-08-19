package elearning

import (
	"ERP-System/config"
	"time"
)

type Course struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Employee Onboarding
	Description string    `gorm:"type:text" json:"description"`
	IsPublished bool      `gorm:"default:false" json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
}

type Slide struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CourseID  uint      `json:"course_id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Type      string    `gorm:"type:varchar(50);default:'document'" json:"type"` // document, video, quiz
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Course{}, &Slide{})
}
