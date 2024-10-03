package flags

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		prompt []string
		want   string
	}{
		// コマンドライン引数（プロンプトの設定用）
		{
			name:   "コマンドライン引数が１つ（期待する用途）",
			prompt: []string{"", "今日の天気は？"},
			want:   "今日の天気は？",
		},
		{
			name:   "コマンドライン引数が１つで空（期待する用途）",
			prompt: []string{"", ""},
			want:   "",
		},
		{
			name:   "コマンドライン引数が複数",
			prompt: []string{"", "今日の天気は？", "うわあああ"},
			want:   "今日の天気は？",
		},
		{
			name:   "コマンドライン引数がない",
			prompt: []string{""},
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.prompt

			if got := Parse(); got != tt.want {
				t.Errorf("hoge() = %v, want %v", got, tt.want)
			}
		})
	}
}
