package main

import (
	// "bufio"
	// "compress/gzip"
	// "io"
	// "net"
	// "encoding/json"
	"os"
	// "bytes"
	// "fmt"
	// "time"
	"net/http"
)

// func main() {
// 	// file, err := os.Create("test.txt")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// file.Write([]byte("日本語書き込み\n"))
// 	// file.Close()

// 	// os.Stdout.Write([]byte("os.Stdout example\n"))

// 	// var buffer bytes.Buffer
// 	// buffer.Write([]byte("bytes.Buffer example\n"))
// 	// fmt.Println(buffer.String())

// 	conn, err := net.Dial("tcp", "ascii.jp:80")
// 	if err != nil {
// 		panic(err)
// 	}

// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
// 	io.Copy(os.Stdout, conn)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("http.ResponseWriter sample"))
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	file, err := os.Create("multiwriter.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	writer := io.MultiWriter(file, os.Stdout)
// 	io.WriteString(writer, "io.MultiWriter example\n")
// }

// func main() {
// 	file, err := os.Create("test.txt.gz")
// 	if err != nil {
// 		panic(err)
// 	}
// 	writer := gzip.NewWriter(file)
// 	writer.Header.Name = "test.txt"
// 	writer.Write([]byte("gzip.Writer example\n"))
// 	writer.Close()
// }

// func main() {
// 	buffer := bufio.NewWriter(os.Stdout)
// 	buffer.WriteString("bufio.Writer ")
// 	buffer.Flush()
// 	buffer.WriteString("example\n")
// 	buffer.Flush()
// }

// func main() {
// 	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
// }

// func main() {
// 	encoder := json.NewEncoder(os.Stdout)
// 	encoder.SetIndent("", " ")
// 	encoder.Encode(map[string]string{
// 		"example": "encoding/json",
// 		"hello":   "world",
// 	})
// }

func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
