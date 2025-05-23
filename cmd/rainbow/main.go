package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gulducat/go-rainbow-logs"
)

func main() {
	// TODO: allow custom regex via CLI flags
	c := rainbow.NewColorWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("reading from stdin...")
	for scanner.Scan() {
		c.Write(scanner.Bytes())
		c.Write([]byte("\n"))
	}
}
