package main

import (
	"io"
	"os"
)

func main() {
	data, _ := os.Open(os.Args[1])
	io.Copy(os.Stdout, data)
}
