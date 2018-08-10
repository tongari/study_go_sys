package main

import (
	"fmt"
	"time"
)

//新しく作られるgoroutineが呼ぶ関数
// func sub() {
// 	fmt.Println("sub() is running")
// 	time.Sleep(time.Second)
// 	fmt.Println("sub() is finished")
// }

// func main() {
// 	fmt.Println("start sub()")
// 	//goroutineを作って関数を実行
// 	go sub()
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("end main()")
// }

func main() {
	fmt.Println("start sub()")
	//インラインで無名関数を作ってその場でgoroutineを実行
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(2 * time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
