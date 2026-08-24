package surveys

import (
	"ERP-System/config"
	"time"
)

type Survey struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	State       string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, open, closed
	CreatedAt   time.Time `json:"created_at"`
}

type SurveyQuestion struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SurveyID     uint      `json:"survey_id"`
	Question     string    `gorm:"type:text;not null" json:"question"`
	QuestionType string    `gorm:"type:varchar(50);default:'text'" json:"question_type"` // text, multiple_choice, rating
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Survey{}, &SurveyQuestion{})
}


// ---- Auto-Generated TableName methods ----
func (Survey) TableName() string {
	return "marketing.surveies"
}

func (SurveyQuestion) TableName() string {
	return "marketing.survey_questions"
}