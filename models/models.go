package models

import "time"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// リクエスト

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// レスポンス

type ChatResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   Message   `json:"message"`
	Done      bool      `json:"done"`
}
