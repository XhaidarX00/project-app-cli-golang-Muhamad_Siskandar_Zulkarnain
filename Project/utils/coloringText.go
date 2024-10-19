package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func ErrorMessage(err error) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s%s\n\n", red("Error : "), err.Error())
}

func SuccesMessage(text string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s%s\n", green("Succes : "), text)
}

func ColorMessage(color_ string, text string) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()

	switch color_ {
	case "red":
		// fmt.Printf("%s\n", red(text))
		return red(text)
	case "green":
		// fmt.Printf("%s\n", green(text))
		return green(text)
	case "blue":
		// fmt.Printf("%s\n", blue(text))
		return blue(text)
	case "yellow":
		// fmt.Printf("%s\n", yellow(text))
		return yellow(text)
	}

	return "0"
}
