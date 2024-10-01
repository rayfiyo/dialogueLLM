package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayfiyo/llms/dialogue/cmd/api"
	"github.com/rayfiyo/llms/dialogue/cmd/files"
	"github.com/rayfiyo/llms/dialogue/cmd/flags"
	"github.com/rayfiyo/llms/dialogue/cmd/generate"
	"github.com/rayfiyo/llms/dialogue/models"
)

func main() {
	flags.Parse()
	prompt := flag.Arg(0)

	fileName := generate.FileName()

	// ログの Markdown に ヘッダー情報として書き込む
	if err := files.Header(fileName, prompt); err != nil {
		log.Fatal("Error writing options in file header: %w", err)
	}

	client := api.NewClient("http://172.27.167.204:11434")

	var messages []models.Message
	var content string
	var err error

	for i := 1; i < *flags.CyclesLimit+1; i++ {

		// 整形
		if *flags.Init != "" {
			i--
			*flags.Init = ""
			generate.Prompt(i, prompt)
		} else {
			generate.Prompt(i, prompt)
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
		content, err = client.Chat(request)
		if err != nil {
			log.Fatalf("Error in switch@%d: %v", i, err)
		}

		// ファイルに保存
		if err := files.Append(fileName,
			"## "+fmt.Sprint(i)+"\n",
		); err != nil {
			log.Fatalf("Error appending to file 1@%d: %v", i, err)
		}
		if err := files.Append(fileName,
			content+"\n",
		); err != nil {
			log.Fatalf("Error appending to file 2@%d: %v", i, err)
		}

		// 次のサイクルに繋げる後処理
		prompt = content
	}
}
