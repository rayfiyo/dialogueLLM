package format

import (
	"fmt"

	"github.com/rayfiyo/dialogueLLM/internal/flags"
)

// func (int, string) string
func Prompt(cycle int, prompt string) (string, error) {
	if cycle < 0 {
		return "", fmt.Errorf("Cycle is negative")
	}

	if cycle%2 != 0 {
		// 1 odd
		if *flags.Model1 != "" {
			flags.Model = flags.Model1
		}
		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
			*flags.Head, *flags.Head1, prompt, *flags.Tail, *flags.Tail1), nil
	} else {
		// 2 even
		if *flags.Model2 != "" {
			flags.Model = flags.Model2
		}
		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
			*flags.Head, *flags.Head2, prompt, *flags.Tail, *flags.Tail2), nil
	}
}
