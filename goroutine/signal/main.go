package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//サイズが１より大きいチャンネルを作成
	signals := make(chan os.Signal, 1)
	// SIGINT(Ctrl+C)を受け取る
	signal.Notify(signals, syscall.SIGINT)
	//シグナルがくるまで待つ
	fmt.Println("waiting SIGINT (ctrl+c)")
	<-signals
	fmt.Println("SIGINT arrived")
}
