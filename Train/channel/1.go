package channel

import (
	"fmt"
	"sync"
	"time"
)

func RandomSelect() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}

func DeadLock() {
	echo := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		echo <- "aaa"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-echo
		fmt.Println(x)
	}()
	wg.Wait()

}

func SelectChan() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	default:
		fmt.Println("default")
	}
}
