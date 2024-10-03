package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayfiyo/llms/dialogue/internal/api"
	"github.com/rayfiyo/llms/dialogue/internal/files"
	"github.com/rayfiyo/llms/dialogue/internal/flags"
	"github.com/rayfiyo/llms/dialogue/internal/format"
	"github.com/rayfiyo/llms/dialogue/internal/generate"
	"github.com/rayfiyo/llms/dialogue/models"
)

func main() {
	flags.Parse()
	prompt := flag.Arg(0)

	fileName := generate.File()

	// ログの Markdown に ヘッダー情報として書き込む
	if err := files.Header(fileName, prompt); err != nil {
		log.Fatal("Error writing options in file header: %w", err)
	}

	client := api.NewClient("http://172.27.167.204:11434")

	var messages []models.Message
	var content string
	var err error

	// ロールの初期化
	initialRole := "user"
	if *flags.CyclesLimit%2 == 0 {
		initialRole = "assistant"
	}
	messages = append(messages, models.Message{
		Role:    initialRole,
		Content: initialPrompt,
	})

	for i := 1; i < *flags.CyclesLimit+1; i++ {

		// 整形
		if *flags.Init != "" {
			i = 0
			*flags.Init = ""
			format.Prompt(i, prompt)
		} else {
			format.Prompt(i, prompt)
		}

		fmt.Print("\n- - - - - - - - - - - -\n")
		log.Printf("%3d:\n\n", i)

		messages = append(messages, models.Message{
			Role:    "user",
			Content: prompt,
		})

		request := &models.ChatRequest{
			Model:    *flags.Model,
			Messages: messages,
		}

		// APIの通信（リクエスト送信とレスポンス取得）
		content, err = client.Chat(request)
		if err != nil {
			log.Fatalf("Error in switch@%d: %v", i, err)
		}

		// ファイルに保存
		if err := files.Append(fileName,
			fmt.Sprintf("## %d\n%s\n", i, content),
		); err != nil {
			log.Fatalf("Error appending to file@%d: %v", i, err)
		}

		// 次のサイクルに繋げる後処理
		prompt = content
	}
}
