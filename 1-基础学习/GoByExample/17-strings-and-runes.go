package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s是一个字符串，包含了一些Unicode字符
	const s = "สวัสดี"

	// len(s)返回字符串s的字节数
	fmt.Println("len(s) =", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	// utf8.RuneCountInString 返回字符串s中的Unicode字符数量
	fmt.Println("rune count =", utf8.RuneCountInString(s))

	for index, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	fmt.Println("\n Using DecodeRuneInString:")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}

}
