package create

import (
	"github.com/rayfiyo/dialogueLLM/internal/constants"
	"github.com/rayfiyo/dialogueLLM/internal/flags"
	"github.com/rayfiyo/dialogueLLM/models"
)

// func (models.ChatRequest, string, int) models.ChatRequest
func Request(input models.ChatRequest, formattedPrompt string, cycle int) models.ChatRequest {
	// 返り値の定義
	newMessage := models.Message{
		Role:    constants.User,
		Content: formattedPrompt,
	}
	output := models.ChatRequest{
		Model:    *flags.Model,
		Messages: make([]models.Message, len(input.Messages)+1),
	}

	// モデルの修正
	if cycle%2 != 0 {
		// 1 odd
		if *flags.Model1 != "" {
			output.Model = *flags.Model1
		}
	} else {
		// 2 even
		if *flags.Model2 != "" {
			output.Model = *flags.Model2
		}
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

	// 追記
	output.Messages[len(input.Messages)] = models.Message{
		Role:    newMessage.Role,
		Content: newMessage.Content,
	}

	return output
}
