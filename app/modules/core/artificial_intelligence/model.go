package artificial_intelligence

import (
	"ERP-System/config"
	"time"
)

type AIPrompt struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"type:varchar(255);default:''" json:"title"`
	Name         string    `gorm:"type:varchar(255);default:''" json:"name"`
	Module       string    `gorm:"type:varchar(100);default:'HR & Payroll'" json:"module"`
	Model        string    `gorm:"type:varchar(100);default:'Gemini 1.5 Pro'" json:"model"`
	Description  string    `gorm:"type:text" json:"description"`
	SystemPrompt string    `gorm:"type:text" json:"system"`
	SampleInput  string    `gorm:"type:text" json:"sampleInput"`
	SampleOutput string    `gorm:"type:text" json:"sampleOutput"`
	Icon         string    `gorm:"type:varchar(50);default:'✨'" json:"icon"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AIConfig struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	ActiveProvider  string    `gorm:"type:varchar(50);default:'GEMINI'" json:"active_provider"` // GEMINI, OPENAI, CLAUDE, DEEPSEEK, OLLAMA
	GeminiApiKey    string    `gorm:"type:text" json:"gemini_api_key"`
	GeminiModel     string    `gorm:"type:varchar(100);default:'gemini-1.5-flash'" json:"gemini_model"`
	OpenAiApiKey    string    `gorm:"type:text" json:"openai_api_key"`
	OpenAiModel     string    `gorm:"type:varchar(100);default:'gpt-4o-mini'" json:"openai_model"`
	OpenAiBaseUrl   string    `gorm:"type:varchar(255);default:'https://api.openai.com/v1'" json:"openai_base_url"`
	ClaudeApiKey    string    `gorm:"type:text" json:"claude_api_key"`
	ClaudeModel     string    `gorm:"type:varchar(100);default:'claude-3-5-sonnet-20241022'" json:"claude_model"`
	DeepseekApiKey  string    `gorm:"type:text" json:"deepseek_api_key"`
	DeepseekModel   string    `gorm:"type:varchar(100);default:'deepseek-chat'" json:"deepseek_model"`
	DeepseekBaseUrl string    `gorm:"type:varchar(255);default:'https://api.deepseek.com/v1'" json:"deepseek_base_url"`
	OllamaBaseUrl   string    `gorm:"type:varchar(255);default:'http://localhost:11434'" json:"ollama_base_url"`
	OllamaModel     string    `gorm:"type:varchar(100);default:'llama3'" json:"ollama_model"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (AIConfig) TableName() string {
	return "setting.ai_configs"
}

func (AIPrompt) TableName() string {
	return "setting.ai_prompts"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &AIPrompt{}, &AIConfig{})
}