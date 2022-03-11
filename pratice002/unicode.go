package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// unicode use case

	f := func(c rune) bool {
		return (!unicode.IsLetter(c) &&
			!unicode.IsNumber(c) &&
			!unicode.Is(unicode.Han, c))
		// ||
		// !unicode.IsPunct(c) ||
		// !unicode.IsControl(c) ||
		// !unicode.IsSymbol(c) ||
		// !unicode.IsSpace(c))
	}
	fmt.Printf("f %d \n", strings.IndexFunc("我的[Smirk]@#123！@#", f))
	fmt.Printf("f %d \n", strings.IndexFunc("我的Smirk]12", f))
	fmt.Printf("f %d \n", strings.IndexFunc("channel", f))
}
