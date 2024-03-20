package main

import (
	"fmt"
	"time"
)

// 接收一个字符串参数，使用两个协程按顺序交替输出每个字符。
// 利用无缓冲通道的阻塞特性
func printStr(s string) {

	ch1 := make(chan int)

	go func() {
		for i := 0; i < len(s); i++ {
			ch1 <- i
			if i%2 == 0 {
				fmt.Println("g1:", string(s[i]))
			}
		}
	}()

	go func() {
		for i := 0; i < len(s); i++ {
			<-ch1
			if i%2 != 0 {
				fmt.Println("g2:", string(s[i]))
			}
		}
	}()

}

func main() {
	printStr("hello,world")
	time.Sleep(time.Second)
}
