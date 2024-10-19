package utils

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	s = strings.ToLower(s)

	capitalizeNext := true
	result := []rune(s)

	for i, r := range result {
		if capitalizeNext && unicode.IsLetter(r) {
			result[i] = unicode.ToUpper(r)
			capitalizeNext = false
		} else if unicode.IsSpace(r) {
			capitalizeNext = true
		}
	}

	return string(result)
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		module := exec.Command("cmd", "/c", "cls")
		module.Stdout = os.Stdout
		module.Run()
	} else {
		module := exec.Command("clear")
		module.Stdout = os.Stdout
		module.Run()
	}
}
