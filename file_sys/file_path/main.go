package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Tmp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
}
