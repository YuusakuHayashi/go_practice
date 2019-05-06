package main

import (
	"fmt"
	"time"
)

func main () {
	//hys_goroutine()
	hys_channel()
}

func hys_goroutine () {
	async1()
	async2()
	go async1()
	go async2()
	time.Sleep(time.Second*5)
}
func async1 () {
	time.Sleep(time.Second*3)
	fmt.Println("async1 end")
}
func async2 () {
	fmt.Println("async2 end")
}
func async3 ( ch_data chan string ) {
	time.Sleep(time.Second*3)
	ch_data<-"hello channel"
}
func hys_channel () {
	ch_data := make(chan string)
	go async3(ch_data)
	//？？？引数として与える場合	
	go async2()
	//先に終了
	fmt.Println(<-ch_data)
	//？？？データを渡す場合
	//ch_dataの参照先にデータが受信されるまで、処理が完了しない
}
