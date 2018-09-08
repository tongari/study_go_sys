package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("%s", os.Getenv("GOPATH"))
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
}
