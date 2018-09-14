package main

import (
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan bool)
	go func() {
		// 終了を通知
		fmt.Println("sub() is finished")
		done <- true
	}()
	// 終了を待つ
	<-done
	fmt.Println("all tasks are finished")
}
