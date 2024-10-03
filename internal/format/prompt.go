package format

import (
	"fmt"

	"github.com/rayfiyo/llms/dialogue/internal/flags"
)

// func (int, string) string
func Prompt(cycle int, prompt string) string {
	if cycle%2 != 0 {
		// 1 odd
		if *flags.Model1 != "" {
			flags.Model = flags.Model1
		}
		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
			*flags.Head, *flags.Head1, prompt, *flags.Tail, *flags.Tail1)
	} else {
		// 2 even
		if *flags.Model2 != "" {
			flags.Model = flags.Model2
		}
		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
			*flags.Head, *flags.Head2, prompt, *flags.Tail, *flags.Tail2)
	}
}
