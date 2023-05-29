package interview

import "fmt"

// 两个协程交替打印数字和letter

var numChan = make(chan struct{})
var letterChan = make(chan struct{})

func PrintNum() {
	for i := 0; i < 10; i++ {
		<-letterChan
		fmt.Println(i)
		numChan <- struct{}{}
	}
}

func PrintLetter() {
	for i := 0; i < 10; i++ {
		<-numChan
		fmt.Println(string(rune('a' + i)))
		letterChan <- struct{}{}
	}
}
