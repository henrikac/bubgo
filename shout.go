package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

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
