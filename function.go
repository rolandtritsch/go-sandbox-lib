package play

import (
	"fmt"
)

func CleanUp() {
	fmt.Println("Hello")
}

func Translate(n int) (string, string) {
	switch n {
	default: return "What?", "Was?"
	case 0: return "Zero", "Null"
	case 1: return "One", "Eins"
	}
}

func RunFunction(n int) string {
	english, german := Translate(n)
	return fmt.Sprintf("%d -> %s/%s", n, english, german)
}

func RunFunction2(s string) string {
	defer CleanUp()
	return s
}
