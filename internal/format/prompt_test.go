package format

import (
	"fmt"
	"os"
	"testing"

	"github.com/rayfiyo/dialogueLLM/models"
)

func TestPrompt(t *testing.T) {
	type args struct {
		cycle  int
		prompt string
	}
	type wants struct {
		want1 string
		want2 string
	}

	tests := []struct {
		name   string
		args   args
		flags  models.Flags // 初期値で試すので省略
		prompt []string
		wants  wants
	}{
		// サイクル
		{
			name:   "サイクルが負",
			args:   args{cycle: -1, prompt: "今日の天気は？"},
			prompt: []string{"", "今日の天気は？"},
			wants: wants{
				want1: "",
				want2: "Cycle is negative",
			},
		},
		{
			name:   "サイクルが０",
			args:   args{cycle: 0, prompt: "今日の天気は？"},
			prompt: []string{"", "今日の天気は？"},
			wants: wants{
				want1: "\n\n今日の天気は？\n\n\n",
				want2: "<nil>",
			},
		},
		{
			name:   "サイクルが正",
			args:   args{cycle: 100000, prompt: "今日の天気は？"},
			prompt: []string{"", "今日の天気は？"},
			wants: wants{
				want1: "\n\n今日の天気は？\n\n\n",
				want2: "<nil>",
			},
		},

		// コマンドライン引数（プロンプトの設定用）
		{
			name:   "コマンドライン引数が１つ（期待する用途）",
			args:   args{cycle: 6, prompt: "今日の天気は？"},
			prompt: []string{"", "今日の天気は？"},
			wants: wants{
				want1: "\n\n今日の天気は？\n\n\n",
				want2: "<nil>",
			},
		},
		{
			name:   "コマンドライン引数が１つで空（期待する用途）",
			args:   args{cycle: 6, prompt: ""},
			prompt: []string{"", ""},
			wants: wants{
				want1: "\n\n\n\n\n",
				want2: "<nil>",
			},
		},
		{
			name:   "コマンドライン引数が複数",
			args:   args{cycle: 6, prompt: "今日の天気は？"},
			prompt: []string{"", "今日の天気は？", "うわあああ"},
			wants: wants{
				want1: "\n\n今日の天気は？\n\n\n",
				want2: "<nil>",
			},
		},
		{
			name:   "コマンドライン引数がない",
			args:   args{cycle: 6},
			prompt: []string{""},
			wants: wants{
				want1: "\n\n\n\n\n",
				want2: "<nil>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.prompt

			if got1, got2 := Prompt(
				tt.args.cycle, tt.args.prompt,
			); got1 != tt.wants.want1 || fmt.Sprint(got2) != tt.wants.want2 {
				t.Errorf("got1: %v, want1: %v\ngot2: %v, want2: %v",
					got1, tt.wants.want1, got2, tt.wants.want2)
			}
		})
	}
}
