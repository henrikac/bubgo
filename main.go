package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
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

type shout []string

func (s *shout) String() string {
	return fmt.Sprint(*s)
}

func (s *shout) Set(value string) error {
	if len(*s) > 0 {
		return errors.New("shout flag already set")
	}
	words := append([]string{value}, flag.Args()...)
	for _, word := range words {
		*s = append(*s, strings.ToUpper(word))
	}
	return nil
}

var shoutFlag shout

func init() {
	const usage = "transforms text to uppercase text: %s <words to shout>"
	flag.Var(&shoutFlag, "shout", fmt.Sprintf(usage, "--shout"))
	flag.Var(&shoutFlag, "s", fmt.Sprintf(usage, "-s"))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bubgo \"your text here\"")
		os.Exit(1)
	}
	flag.Parse()
	var str string
	if len(shoutFlag) > 0 {
		str = strings.Join(shoutFlag, " ")
	} else {
		str = strings.Join(os.Args[1:], " ")
	}
	for _, c := range str {
		if char, found := chars[c]; found {
			fmt.Printf("%c", char)
			continue
		}
		fmt.Printf("%s", string(c))
	}
	fmt.Printf("\n")
}
