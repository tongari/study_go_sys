package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	fmt.Println("clean: " + filepath.Clean("./path/filepath/../path.go"))

	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println("abspath: " + abspath)

	relPath, _ := filepath.Rel("/Users/tongari/", "/Users/tongari/go/src/github.com")
	fmt.Println("relPath: " + relPath)

	fmt.Println("${GOPATH}: " + clean(os.ExpandEnv("${GOPATH}")))
}

func clean(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}
