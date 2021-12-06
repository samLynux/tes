package main

import (
	"fmt"
	"strings"
)

func main() {
	x := findFirstStringInBracket("string diluar bracket (string pertama) string diluar bracket (string kedua)")
	fmt.Println(x)
}

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		openBracket := strings.Index(str, "(")
		if openBracket >= 0 {
			closeBracket := strings.Index(str[openBracket:], ")")
			if closeBracket >= 0 {
				return str[(openBracket + 1):(openBracket + closeBracket)]
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}

}
