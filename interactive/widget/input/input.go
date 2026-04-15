package input

import (
	"matman0497/sshboy/interactive/style"

	"charm.land/bubbles/v2/textinput"
)

func New(prompt string) textinput.Model {

	//server name input
	var nameInput textinput.Model
	nameInput = textinput.New()
	nameInput.Prompt = prompt
	nameInput.CharLimit = 32

	textinputStyle := nameInput.Styles()
	textinputStyle.Cursor.Color = style.PrimaryColor
	textinputStyle.Focused.Prompt = style.FocusedStyle
	textinputStyle.Focused.Text = style.FocusedStyle
	textinputStyle.Blurred.Prompt = style.BlurredStyle
	textinputStyle.Focused.Text = style.FocusedStyle
	nameInput.SetStyles(textinputStyle)

	return nameInput

}
