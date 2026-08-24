package elearning

import (
	"ERP-System/app/modules/core/base"
	"ERP-System/config"
)

// 9. eLearning Certification
type Certification struct {
	ID         uint          `gorm:"primaryKey" json:"id"`
	CourseID   uint          `json:"course_id"`
	Course     *Course       `gorm:"foreignKey:CourseID"` // Auto-added relation
	CustomerID uint          `json:"customer_id"`
	Customer   *base.Partner `gorm:"foreignKey:CustomerID" json:"customer,omitempty"` // Cross-module relation
	Score      float64       `gorm:"type:numeric(5,2)" json:"score"`
	PdfUrl     string        `gorm:"type:varchar(255)" json:"pdf_url"` // Link sertifikat
}

func (Certification) TableName() string {
	return "website_portal.certifications"
}

func (Course) TableName() string {
	return "website_portal.courses"
}

func (Slide) TableName() string {
	return "website_portal.slides"
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &Certification{}) }

type Course struct {
	ID uint `gorm:"primaryKey"`
}
type Slide struct {
	ID uint `gorm:"primaryKey"`
}
