package artificial_intelligence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: 60 * time.Second,
}

// CallGemini memanggil REST API Google Gemini v1beta
func CallGemini(apiKey, model, systemPrompt, userInput string) (string, error) {
	if model == "" {
		model = "gemini-1.5-flash"
	}
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)

	payload := map[string]interface{}{
		"systemInstruction": map[string]interface{}{
			"parts": []map[string]interface{}{
				{"text": systemPrompt},
			},
		},
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": userInput},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"temperature": 0.7,
		},
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("koneksi ke Google Gemini gagal: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Google Gemini error (%d): %s", resp.StatusCode, string(respBody))
	}

	var geminiResp struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.Unmarshal(respBody, &geminiResp); err != nil {
		return "", err
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		return geminiResp.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("respon kosong dari Google Gemini")
}

// CallOpenAI memanggil OpenAI Chat Completions API (GPT-4o, GPT-4o-mini, dll)
func CallOpenAI(apiKey, model, baseURL, systemPrompt, userInput string) (string, error) {
	if model == "" {
		model = "gpt-4o-mini"
	}
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	baseURL = strings.TrimRight(baseURL, "/")
	url := fmt.Sprintf("%s/chat/completions", baseURL)

	messages := []map[string]string{}
	if systemPrompt != "" {
		messages = append(messages, map[string]string{"role": "system", "content": systemPrompt})
	}
	messages = append(messages, map[string]string{"role": "user", "content": userInput})

	payload := map[string]interface{}{
		"model":       model,
		"messages":    messages,
		"temperature": 0.7,
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("koneksi ke OpenAI gagal: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OpenAI error (%d): %s", resp.StatusCode, string(respBody))
	}

	var openAiResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(respBody, &openAiResp); err != nil {
		return "", err
	}

	if len(openAiResp.Choices) > 0 {
		return openAiResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("respon kosong dari OpenAI")
}

// CallClaude memanggil Anthropic Claude API (Claude 3.5 Sonnet, Claude 3 Haiku)
func CallClaude(apiKey, model, systemPrompt, userInput string) (string, error) {
	if model == "" {
		model = "claude-3-5-sonnet-20241022"
	}
	url := "https://api.anthropic.com/v1/messages"

	payload := map[string]interface{}{
		"model":      model,
		"max_tokens": 2048,
		"system":     systemPrompt,
		"messages": []map[string]string{
			{"role": "user", "content": userInput},
		},
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("koneksi ke Anthropic Claude gagal: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Anthropic Claude error (%d): %s", resp.StatusCode, string(respBody))
	}

	var claudeResp struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal(respBody, &claudeResp); err != nil {
		return "", err
	}

	for _, item := range claudeResp.Content {
		if item.Type == "text" {
			return item.Text, nil
		}
	}

	return "", fmt.Errorf("respon kosong dari Anthropic Claude")
}

// CallOllama memanggil local Ollama server (Llama 3, Mistral, Qwen, DeepSeek)
func CallOllama(baseURL, model, systemPrompt, userInput string) (string, error) {
	if model == "" {
		model = "llama3"
	}
	if baseURL == "" {
		baseURL = "http://localhost:11434"
	}
	baseURL = strings.TrimRight(baseURL, "/")
	url := fmt.Sprintf("%s/api/generate", baseURL)

	prompt := userInput
	if systemPrompt != "" {
		prompt = fmt.Sprintf("System: %s\n\nUser: %s", systemPrompt, userInput)
	}

	payload := map[string]interface{}{
		"model":  model,
		"prompt": prompt,
		"stream": false,
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("koneksi ke Ollama gagal: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama error (%d): %s", resp.StatusCode, string(respBody))
	}

	var ollamaResp struct {
		Response string `json:"response"`
	}

	if err := json.Unmarshal(respBody, &ollamaResp); err != nil {
		return "", err
	}

	return ollamaResp.Response, nil
}

// ExecuteLLM adalah router cerdas yang memanggil provider yang dipilih (dengan fallback ke .env)
func ExecuteLLM(cfg *AIConfig, provider, model, systemPrompt, userInput string) (string, string, error) {
	// 1. Tentukan Provider yang aktif
	if provider == "" {
		if cfg != nil && cfg.ActiveProvider != "" {
			provider = cfg.ActiveProvider
		} else if envProv := os.Getenv("AI_DEFAULT_PROVIDER"); envProv != "" {
			provider = envProv
		} else {
			provider = "GEMINI"
		}
	}
	provider = strings.ToUpper(provider)

	var result string
	var err error
	actualModel := model

	switch provider {
	case "GEMINI":
		apiKey := ""
		if cfg != nil && cfg.GeminiApiKey != "" {
			apiKey = cfg.GeminiApiKey
		}
		if apiKey == "" {
			apiKey = os.Getenv("GEMINI_API_KEY")
		}
		if actualModel == "" {
			if cfg != nil && cfg.GeminiModel != "" {
				actualModel = cfg.GeminiModel
			} else if envModel := os.Getenv("GEMINI_MODEL"); envModel != "" {
				actualModel = envModel
			} else {
				actualModel = "gemini-1.5-flash"
			}
		}
		if apiKey == "" {
			return "", provider + " (" + actualModel + ")", fmt.Errorf("API Key Google Gemini belum dikonfigurasi (isi di Pengaturan AI atau .env)")
		}
		result, err = CallGemini(apiKey, actualModel, systemPrompt, userInput)

	case "OPENAI", "GPT":
		provider = "OPENAI"
		apiKey := ""
		baseURL := "https://api.openai.com/v1"
		if cfg != nil && cfg.OpenAiApiKey != "" {
			apiKey = cfg.OpenAiApiKey
			if cfg.OpenAiBaseUrl != "" {
				baseURL = cfg.OpenAiBaseUrl
			}
		}
		if apiKey == "" {
			apiKey = os.Getenv("OPENAI_API_KEY")
			if envBase := os.Getenv("OPENAI_BASE_URL"); envBase != "" {
				baseURL = envBase
			}
		}
		if actualModel == "" {
			if cfg != nil && cfg.OpenAiModel != "" {
				actualModel = cfg.OpenAiModel
			} else if envModel := os.Getenv("OPENAI_MODEL"); envModel != "" {
				actualModel = envModel
			} else {
				actualModel = "gpt-4o-mini"
			}
		}
		if apiKey == "" {
			return "", provider + " (" + actualModel + ")", fmt.Errorf("API Key OpenAI GPT belum dikonfigurasi (isi di Pengaturan AI atau .env)")
		}
		result, err = CallOpenAI(apiKey, actualModel, baseURL, systemPrompt, userInput)

	case "CLAUDE", "ANTHROPIC":
		provider = "CLAUDE"
		apiKey := ""
		if cfg != nil && cfg.ClaudeApiKey != "" {
			apiKey = cfg.ClaudeApiKey
		}
		if apiKey == "" {
			apiKey = os.Getenv("ANTHROPIC_API_KEY")
		}
		if actualModel == "" {
			if cfg != nil && cfg.ClaudeModel != "" {
				actualModel = cfg.ClaudeModel
			} else if envModel := os.Getenv("ANTHROPIC_MODEL"); envModel != "" {
				actualModel = envModel
			} else {
				actualModel = "claude-3-5-sonnet-20241022"
			}
		}
		if apiKey == "" {
			return "", provider + " (" + actualModel + ")", fmt.Errorf("API Key Anthropic Claude belum dikonfigurasi (isi di Pengaturan AI atau .env)")
		}
		result, err = CallClaude(apiKey, actualModel, systemPrompt, userInput)

	case "DEEPSEEK":
		apiKey := ""
		baseURL := "https://api.deepseek.com/v1"
		if cfg != nil && cfg.DeepseekApiKey != "" {
			apiKey = cfg.DeepseekApiKey
			if cfg.DeepseekBaseUrl != "" {
				baseURL = cfg.DeepseekBaseUrl
			}
		}
		if apiKey == "" {
			apiKey = os.Getenv("DEEPSEEK_API_KEY")
			if envBase := os.Getenv("DEEPSEEK_BASE_URL"); envBase != "" {
				baseURL = envBase
			}
		}
		if actualModel == "" {
			if cfg != nil && cfg.DeepseekModel != "" {
				actualModel = cfg.DeepseekModel
			} else if envModel := os.Getenv("DEEPSEEK_MODEL"); envModel != "" {
				actualModel = envModel
			} else {
				actualModel = "deepseek-chat"
			}
		}
		if apiKey == "" {
			return "", provider + " (" + actualModel + ")", fmt.Errorf("API Key DeepSeek belum dikonfigurasi (isi di Pengaturan AI atau .env)")
		}
		result, err = CallOpenAI(apiKey, actualModel, baseURL, systemPrompt, userInput)

	case "OLLAMA":
		baseURL := "http://localhost:11434"
		if cfg != nil && cfg.OllamaBaseUrl != "" {
			baseURL = cfg.OllamaBaseUrl
		} else if envBase := os.Getenv("OLLAMA_BASE_URL"); envBase != "" {
			baseURL = envBase
		}
		if actualModel == "" {
			if cfg != nil && cfg.OllamaModel != "" {
				actualModel = cfg.OllamaModel
			} else if envModel := os.Getenv("OLLAMA_MODEL"); envModel != "" {
				actualModel = envModel
			} else {
				actualModel = "llama3"
			}
		}
		result, err = CallOllama(baseURL, actualModel, systemPrompt, userInput)

	default:
		return "", provider, fmt.Errorf("provider AI '%s' belum didukung", provider)
	}

	return result, provider + " (" + actualModel + ")", err
}

