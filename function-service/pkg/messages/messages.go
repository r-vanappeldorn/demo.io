package messages

import (
	"github.com/fatih/color"
)

func Error() string {
	return color.New(color.FgRed).Sprintf("\n[error]")
}

func Succes() string {
	return color.New(color.FgGreen).Sprintf("[success]")
}