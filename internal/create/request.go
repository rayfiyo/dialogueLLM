package create

import (
	"github.com/rayfiyo/dialogueLLM/internal/constants"
	"github.com/rayfiyo/dialogueLLM/models"
)

// func (models.ChatRequest, models.Message) models.ChatRequest
func Request(input models.ChatRequest, newMessage models.Message) models.ChatRequest {
	// 返り値の定義
	output := models.ChatRequest{
		Model:    input.Model,
		Messages: make([]models.Message, len(input.Messages)+1),
	}

	// ロールの入れ替え
	for i, msg := range input.Messages {
		newRole := constants.User
		if msg.Role == newRole {
			newRole = constants.Assistant
		}
		output.Messages[i] = models.Message{
			Role:    newRole,
			Content: msg.Content,
		}
	}

	// 追記分
	output.Messages[len(input.Messages)] = models.Message{
		Role:    newMessage.Role,
		Content: newMessage.Content,
	}

	return output
}
