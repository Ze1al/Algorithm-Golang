package interview

import "fmt"

// 启动两个协程，一个打奇数 一个负责打偶数

func PrintNum2() {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			c <- i
		}(i)

		fmt.Println(<-c)
	}
}
