package statement

import (
	"fmt"
	"runtime"
	"sync"
)

type count2 struct {
	num   int
	mutex sync.Mutex
}

func (c *count2) increment() {
	// 공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num++
	// 공유 데이터 수정 후 보호 해제
	c.mutex.Unlock()
}

func (c *count2) result() {
	fmt.Println("Result: ", c.num)
}

func GoSync2() {
	/*
		뮤텍스(Mutex): 여러 고루틴에서 작업하는 공유 데이터 보호
		sync.Mutex 선언 후 Lock, Unlock 사용
	*/

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count2{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10_000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	for i := 1; i < 10_000; i++ {
		<-done
	}

	c.result()
}
