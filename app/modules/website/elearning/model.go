package elearning

import "ERP-System/config"

// 9. eLearning Certification
type Certification struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	CourseID   uint    `json:"course_id"`
	CustomerID uint    `json:"customer_id"`
	Score      float64 `gorm:"type:numeric(5,2)" json:"score"`
	PdfUrl     string  `gorm:"type:varchar(255)" json:"pdf_url"` // Link sertifikat
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &Certification{}) }

type Course struct {
	ID uint `gorm:"primaryKey"`
}
type Slide struct {
	ID uint `gorm:"primaryKey"`
}
