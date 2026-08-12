package elearning

import (
	"ERP-System/config"
	"time"
)

type Course struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"type:varchar(255)" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	InstructorID uint      `json:"instructor_id"` // Employee ID
	State        string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, published
	CreatedAt    time.Time `json:"created_at"`
}

type Lesson struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CourseID  uint      `json:"course_id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	VideoURL  string    `gorm:"type:varchar(255)" json:"video_url"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Course{}, &Lesson{})
}
