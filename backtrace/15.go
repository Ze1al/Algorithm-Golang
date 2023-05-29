package backTrace

import "fmt"

func b() (i int) {
	defer func() {
		i++
		fmt.Printf("defer2:%d\n", i)
	}()
	defer func() {
		i++
		fmt.Printf("defer1:%d\n", i)
	}()
	return i
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Printf("defer2:%d\n", i)
	}()
	defer func() {
		i++
		fmt.Printf("defer1:%d\n", i)
	}()
	return &i
}
