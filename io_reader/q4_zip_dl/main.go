package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment;filename=ascii_sample.zip")

	//zipファイル
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	//ファイルの数だけ書き込み
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("１つめのファイルのテキストです"))
	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("２つめのファイルのテキストです"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
