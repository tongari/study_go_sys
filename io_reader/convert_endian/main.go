package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32bitのビッグエンディアンのデータ(10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	//エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Println("data: %d\n", i)
}
