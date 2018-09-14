package main

import (
	"fmt"
	"time"
)

// 新しくつkジュられるgorutineが呼ぶ関数
func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("start sub()")
	// gorutineを作って関数を実行
	// go sub()

	// インラインで無名関数を作ってその場でgoroutineで実行
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()

	time.Sleep(2 * time.Second)
}
