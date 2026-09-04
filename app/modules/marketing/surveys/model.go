package surveys

import (
	"ERP-System/app/modules/auth"
	"ERP-System/config"
	"time"
)

type Survey struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          *uint      `json:"user_id"`
	User            *auth.User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Title           string    `gorm:"type:varchar(255);not null" json:"title"`
	Description     string    `gorm:"type:text" json:"description"`
	GformURL        string    `gorm:"type:varchar(500)" json:"gform_url"` // URL Google Form jika menggunakan form eksternal
	State           string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, open, closed
	NpsScore        int       `gorm:"default:0" json:"nps_score"` // Skala -100 sampai +100
	ResponsesCount  int       `gorm:"default:0" json:"responses_count"`
	PromotersCount  int       `gorm:"default:0" json:"promoters_count"`
	PassivesCount   int       `gorm:"default:0" json:"passives_count"`
	DetractorsCount int       `gorm:"default:0" json:"detractors_count"`
	CreatedAt       time.Time `json:"created_at"`
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