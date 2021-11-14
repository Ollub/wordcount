package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func wordcount(input io.Reader, output io.Writer) error {
	text := os.Args[1]
	if text == "" {
		fmt.Println(0)
		return nil
	}
	words := strings.Split(text, " ")
	fmt.Println(len(words))
	return nil
}

func main() {
	wordcount(os.Stdin, os.Stdout)
}
