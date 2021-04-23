package main

import (
	"fmt"
)

func main() {
	// host := os.Args[1]

	var openPorts []int

	// if host == "" {
	// 	fmt.Print(("error"))
	// 	return
	// }

	// 容量10のchannel作成
	c2 := make(chan int, 10)
	words := []string{"hoge", "fuga"}
	for w := range words {
		fmt.Print(w)
	}
	c := make(chan int, 10)

	// キューに送信
	c <- 10

	// キューから受信し、値を出力
	// キューに値がなければロック
	fmt.Println(<-c) //10

	for i := 0; i < 4; i++ {
		//time.Sleep(1 * time.Second)
		openPorts = append(openPorts, i)
		//fmt.Printf(i)
	}
	for _, p := range openPorts {
		fmt.Printf("%d\n", p)
	}
}
