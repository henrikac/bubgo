package main

import (
	"fmt"
	"os"
)

var chars = map[rune]rune{
	'a': 'ⓐ',
	'b': 'ⓑ',
	'c': 'ⓒ',
	'd': 'ⓓ',
	'e': 'ⓔ',
	'f': 'ⓕ',
	'g': 'ⓖ',
	'h': 'ⓗ',
	'i': 'ⓘ',
	'j': 'ⓙ',
	'k': 'ⓚ',
	'l': 'ⓛ',
	'm': 'ⓜ',
	'n': 'ⓝ',
	'o': 'ⓞ',
	'p': 'ⓟ',
	'q': 'ⓠ',
	'r': 'ⓡ',
	's': 'ⓢ',
	't': 'ⓣ',
	'u': 'ⓤ',
	'v': 'ⓥ',
	'w': 'ⓦ',
	'x': 'ⓧ',
	'y': 'ⓨ',
	'z': 'ⓩ',
	'A': 'Ⓐ',
	'B': 'Ⓑ',
	'C': 'Ⓒ',
	'D': 'Ⓓ',
	'E': 'Ⓔ',
	'F': 'Ⓕ',
	'G': 'Ⓖ',
	'H': 'Ⓗ',
	'I': 'Ⓘ',
	'J': 'Ⓙ',
	'K': 'Ⓚ',
	'L': 'Ⓛ',
	'M': 'Ⓜ',
	'N': 'Ⓝ',
	'O': 'Ⓞ',
	'P': 'Ⓟ',
	'Q': 'Ⓠ',
	'R': 'Ⓡ',
	'S': 'Ⓢ',
	'T': 'Ⓣ',
	'U': 'Ⓤ',
	'V': 'Ⓥ',
	'W': 'Ⓦ',
	'X': 'Ⓧ',
	'Y': 'Ⓨ',
	'Z': 'Ⓩ',
	'0': '⓪',
	'1': '①',
	'2': '②',
	'3': '③',
	'4': '④',
	'5': '⑤',
	'6': '⑥',
	'7': '⑦',
	'8': '⑧',
	'9': '⑨',
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bubgo \"your text here\"")
		os.Exit(1)
	}
	str := os.Args[1]
	for _, c := range str {
		if char, found := chars[c]; found {
			fmt.Printf("%c", char)
			continue
		}
		fmt.Printf("%s", string(c))
	}
	fmt.Printf("\n")
}
