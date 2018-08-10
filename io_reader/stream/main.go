package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("---- HEADER ----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	fotter := bytes.NewBufferString("---- FOOTER ----\n")

	reader := io.MultiReader(header, content, fotter)
	io.Copy(os.Stdout, reader)
}
