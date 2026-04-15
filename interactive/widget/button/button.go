package button

import (
	"fmt"
	"matman0497/sshboy/interactive/style"
)

func NewFocused(text string) string {

	return style.FocusedStyle.Render(fmt.Sprintf("[ %s ]", text))

}

func NewBlurred(text string) string {

	return fmt.Sprintf("[ %s ]", style.BlurredStyle.Render(text))

}
