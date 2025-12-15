package statement

import "fmt"

func GoChannel6() {
	// Close: 채널 닫기

	ch := make(chan int)

	go func() {
		for range 3 {
			ch <- 7777
		}
	}()

	val2, ok1 := <-ch
	fmt.Println(val2, ok1)
	val2, ok2 := <-ch
	fmt.Println(val2, ok2)
	val3, ok3 := <-ch
	fmt.Println(val3, ok3)

	close(ch)
	val4, ok4 := <-ch
	fmt.Println(val4, ok4)
}
