package interview

import "fmt"

var (
	dog  = make(chan struct{})
	cat  = make(chan struct{})
	fish = make(chan struct{})
)

func Dog() {
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
}

func Cat() {
	<-dog
	fmt.Println("cat")
	cat <- struct{}{}
}

func Fish() {
	<-cat
	fmt.Println("fish")
	fish <- struct{}{}
}
