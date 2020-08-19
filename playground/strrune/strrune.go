package strrune

import (
	"fmt"
	"unicode/utf8"
)

func LaunchStrrune() {
	const nihongo = "日本語"
	hhBytes := []byte(nihongo)
	for _, bb := range hhBytes {
		fmt.Println("byte ", bb)
	}
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
