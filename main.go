package main

import (
	"os"

	"github.com/maiyama18/immonkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
