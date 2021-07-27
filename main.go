package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
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
	bubbles := ""
	for _, c := range str {
		if char, found := chars[c]; found {
			bubbles += string(char)
			continue
		}
		bubbles += string(c)
	}
	fmt.Printf("%s\n", bubbles)

	copyToClipboard(bubbles)
}

func copyToClipboard(s string) error {
	var c1 *exec.Cmd
	var c2 *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		_, err := exec.LookPath("clip.exe")
		if err != nil {
			return err
		}
		c1 = exec.Command("echo", s)
		c2 = exec.Command("clip.exe")
	case "darwin":
		_, err := exec.LookPath("pbcopy")
		if err != nil {
			return err
		}
		c1 = exec.Command("printf", s)
		c2 = exec.Command("pbcopy")
	default: // linux
		_, err := exec.LookPath("xclip")
		if err != nil {
			return err
		}
		c1 = exec.Command("printf", s)
		c2 = exec.Command("xclip", "-sel", "clip")
	}

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	if err := c1.Start(); err != nil {
		return err
	}
	if err := c2.Start(); err != nil {
		return err
	}
	if err := c1.Wait(); err != nil {
		return err
	}
	w.Close()
	return c2.Wait()
}
