package messages

import (
	"github.com/fatih/color"
)

func Error() string {
	return color.New(color.FgRed).Sprintf("\n[error]")
}
