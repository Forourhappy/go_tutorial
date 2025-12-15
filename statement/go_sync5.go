package statement

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func GoSync5() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	mutex := new(sync.Mutex)
	condition := sync.NewCond(mutex)

	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Wating: ", n)
			condition.Wait() //고루틴 대기
			fmt.Println("Wating End: ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c

	}

	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("Wake Goroutine(Signal): ", i)
	// 	condition.Signal() // 한 개씩 깨움(모든 고루틴 생성 후)
	// 	mutex.Unlock()
	// }

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
