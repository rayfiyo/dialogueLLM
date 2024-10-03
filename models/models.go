package models

import "time"

// メッセージ（リクエストとレスポンスの中身）
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

// フラグ
type Flags struct {
	Model       string
	Model1      string
	Model2      string
	CyclesLimit string
	Head        string
	Head1       string
	Head2       string
	Tail        string
	Tail1       string
	Tail2       string
}
