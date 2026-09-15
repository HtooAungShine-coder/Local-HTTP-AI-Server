package API

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// The individual message structure
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Used only for the request payload
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"` // Corrected: plural slice
}

// A single choice wrapper returned in the choices array
type Choice struct {
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"` // Kept singular: OpenRouter responses return a single message per choice
}

// The token usage metadata block
type TokenUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// The main structural wrapper matching OpenRouter's response payload
type OpenRouterResponse struct {
	ID       string     `json:"id"`
	Provider string     `json:"provider"`
	Model    string     `json:"model"`
	Object   string     `json:"object"`
	Created  int        `json:"created"`
	Choices  []Choice   `json:"choices"`
	Usage    TokenUsage `json:"usage"`
}

func SendMessage(message string) (string, error) {
	request := ChatRequest{
		Model: "openrouter/free",
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: message,
			},
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("error: OPENROUTER_API_KEY environment variable is not set")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("HTTP-Referer", "http://localhost:8080") // Required by OpenRouter
	req.Header.Set("X-Title", "Luna AI Terminal")           // Required by OpenRouter

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Unmarshal the response into your OpenRouterResponse struct
	var openRouterResp OpenRouterResponse
	err = json.Unmarshal(body, &openRouterResp)
	if err != nil {
		return "", err
	}

	//Return the text content of the first choice response
	if len(openRouterResp.Choices) > 0 {
		return openRouterResp.Choices[0].Message.Content, nil
	}

	return "No response generated.", nil
}
