package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Server is running at localhost:8888")

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			panic(err)
// 		}
// 		go func() {
// 			fmt.Printf("Accept %v\n", conn.RemoteAddr())
// 			//リクエストを読み込む
// 			request, err := http.ReadRequest(
// 				bufio.NewReader(conn),
// 			)
// 			if err != nil {
// 				panic(err)
// 			}
// 			dump, err := httputil.DumpRequest(request, true)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(string(dump))
// 			//レスポンスを書き込む
// 			responce := http.Response{
// 				StatusCode: 200,
// 				ProtoMajor: 1,
// 				ProtoMinor: 0,
// 				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
// 			}
// 			responce.Write(conn)
// 			conn.Close()
// 		}()

// 	}
// }

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			var responseBuffer bytes.Buffer
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			for {
				//タイムアウトを設定
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				//リクエストを読み込む
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					//タイムアウトもしくはソケットクローズ時は終了
					//それ以外はエラーにする
					neterr, ok := err.(net.Error) //ダウンキャスト
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				//リクエスト表示
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				// content := "Hello World\n"
				responseBuffer.WriteString(string(dump))
				content := responseBuffer.String()
				//レスポンスを書き込む
				// HTTP/1.1 かつ、ContentLengthの設定が必要
				responce := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				responce.Write(conn)
			}
			conn.Close()
		}()

	}
}
