package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/gulducat/go-rainbow-logs"
)

/* TODO:
 * help
 * custom regex
 */

func main() {
	var src string
	var reader io.ReadCloser
	var err error

	if hasStdin() {
		src = "stdin"
		reader = os.Stdin

	} else if len(os.Args) > 1 {
		src = os.Args[1]
		reader, err = os.Open(src)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		defer reader.Close()
	}

	if reader == nil {
		fmt.Println("no rainbows to read")
		os.Exit(1)
	}

	fmt.Printf("reading rainbows from %s...\n", src)

	c := rainbow.NewColorWriter(os.Stdout)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		c.Write(scanner.Bytes())
		c.Write([]byte("\n"))
	}
}

func hasStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
