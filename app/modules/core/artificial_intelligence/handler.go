package artificial_intelligence

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateAIPrompt godoc
// @Summary Create a new AIPrompt
// @Description Create a new AIPrompt in the system
// @Tags artificial_intelligence
// @Accept json
// @Produce json
// @Success 201 {object} AIPrompt
// @Param request body AIPrompt true "Payload"
// @Router /api/artificial_intelligence [post]
// @Security BearerAuth
func CreateAIPromptHandler(c echo.Context) error {
	var data AIPrompt
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAIPromptService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllAIPrompt godoc
// @Summary Get all AIPrompt
// @Description Retrieve a list of all AIPrompt
// @Tags artificial_intelligence
// @Produce json
// @Success 200 {object} []AIPrompt
// @Router /api/artificial_intelligence [get]
// @Security BearerAuth
func GetAllAIPromptHandler(c echo.Context) error {
	data, err := GetAllAIPromptService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetAIPromptByID godoc
// @Summary Get a AIPrompt by ID
// @Description Retrieve a specific AIPrompt by its ID
// @Tags artificial_intelligence
// @Produce json
// @Param id path int true "AIPrompt ID"
// @Success 200 {object} AIPrompt
// @Router /api/artificial_intelligence/{id} [get]
// @Security BearerAuth
func GetAIPromptByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAIPromptByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateAIPrompt godoc
// @Summary Update a AIPrompt
// @Description Update an existing AIPrompt
// @Tags artificial_intelligence
// @Accept json
// @Produce json
// @Param id path int true "AIPrompt ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/artificial_intelligence/{id} [put]
// @Security BearerAuth
func UpdateAIPromptHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAIPromptByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAIPromptService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteAIPrompt godoc
// @Summary Delete a AIPrompt
// @Description Delete a AIPrompt by ID
// @Tags artificial_intelligence
// @Produce json
// @Param id path int true "AIPrompt ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/artificial_intelligence/{id} [delete]
// @Security BearerAuth
func DeleteAIPromptHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAIPromptService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

func maskSecret(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	if len(s) <= 8 {
		return "••••••••"
	}
	return s[:4] + "••••••••" + s[len(s)-4:]
}

func GetAIConfigHandler(c echo.Context) error {
	cfg, err := GetActiveAIConfig()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil konfigurasi AI", err.Error())
	}

	envGemini := os.Getenv("GEMINI_API_KEY") != ""
	envOpenAI := os.Getenv("OPENAI_API_KEY") != ""
	envClaude := os.Getenv("ANTHROPIC_API_KEY") != ""
	envDeepseek := os.Getenv("DEEPSEEK_API_KEY") != ""

	resp := map[string]interface{}{
		"id":                  cfg.ID,
		"active_provider":     cfg.ActiveProvider,
		"gemini_api_key":      maskSecret(cfg.GeminiApiKey),
		"gemini_model":        cfg.GeminiModel,
		"gemini_configured":   cfg.GeminiApiKey != "" || envGemini,
		"gemini_from_env":     cfg.GeminiApiKey == "" && envGemini,
		"openai_api_key":      maskSecret(cfg.OpenAiApiKey),
		"openai_model":        cfg.OpenAiModel,
		"openai_base_url":     cfg.OpenAiBaseUrl,
		"openai_configured":   cfg.OpenAiApiKey != "" || envOpenAI,
		"openai_from_env":     cfg.OpenAiApiKey == "" && envOpenAI,
		"claude_api_key":      maskSecret(cfg.ClaudeApiKey),
		"claude_model":        cfg.ClaudeModel,
		"claude_configured":   cfg.ClaudeApiKey != "" || envClaude,
		"claude_from_env":     cfg.ClaudeApiKey == "" && envClaude,
		"deepseek_api_key":    maskSecret(cfg.DeepseekApiKey),
		"deepseek_model":      cfg.DeepseekModel,
		"deepseek_base_url":   cfg.DeepseekBaseUrl,
		"deepseek_configured": cfg.DeepseekApiKey != "" || envDeepseek,
		"deepseek_from_env":   cfg.DeepseekApiKey == "" && envDeepseek,
		"ollama_base_url":     cfg.OllamaBaseUrl,
		"ollama_model":        cfg.OllamaModel,
		"ollama_configured":   true,
	}

	return utils.SendSuccess(c, http.StatusOK, "Konfigurasi AI berhasil diambil", resp)
}

func SaveAIConfigHandler(c echo.Context) error {
	var payload AIConfig
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if err := SaveAIConfig(&payload); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan konfigurasi AI", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Konfigurasi AI berhasil disimpan", nil)
}

func GenerateAISimulationHandler(c echo.Context) error {
	var payload struct {
		PromptID uint   `json:"prompt_id"`
		Provider string `json:"provider"`
		Model    string `json:"model"`
		System   string `json:"system"`
		Input    string `json:"input"`
	}
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	systemPrompt := payload.System
	userInput := payload.Input
	var promptObj *AIPrompt

	if payload.PromptID > 0 {
		p, err := GetAIPromptByID(payload.PromptID)
		if err == nil && p != nil {
			promptObj = p
			if systemPrompt == "" {
				systemPrompt = p.SystemPrompt
			}
			if userInput == "" {
				userInput = p.SampleInput
			}
		}
	}

	if userInput == "" {
		userInput = "Berikan analisis komprehensif terkait efisiensi operasional ERP perusahaan."
	}

	cfg, _ := GetActiveAIConfig()
	start := time.Now()
	result, modelUsed, err := ExecuteLLM(cfg, payload.Provider, payload.Model, systemPrompt, userInput)
	latency := time.Since(start).Milliseconds()

	if err != nil {
		if promptObj != nil && promptObj.SampleOutput != "" {
			notice := fmt.Sprintf("<div class=\"p-2.5 mb-3 bg-amber-50 dark:bg-amber-900/30 border border-amber-200 dark:border-amber-800 rounded text-amber-700 dark:text-amber-300 text-xs font-medium\">ℹ️ <strong>Simulasi Offline:</strong> %s. Menampilkan data respons standar.</div>", err.Error())
			return utils.SendSuccess(c, http.StatusOK, "AI fallback output generated", map[string]interface{}{
				"result":        notice + promptObj.SampleOutput,
				"tokens_used":   135,
				"latency_ms":    latency,
				"model_version": modelUsed + " (Offline Mode)",
				"is_live":       false,
			})
		}
		return utils.SendError(c, http.StatusBadRequest, "Eksekusi AI Gagal", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "AI output generated successfully", map[string]interface{}{
		"result":        result,
		"tokens_used":   (len(userInput) + len(result)) / 4,
		"latency_ms":    latency,
		"model_version": modelUsed,
		"is_live":       true,
	})
}


