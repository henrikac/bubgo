package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

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
