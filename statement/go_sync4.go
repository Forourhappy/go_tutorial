package statement

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func GoSync4() {
	// RWMutex: 쓰기 Lock -> 쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 X, 읽기 락, 쓰기 락 전부 방어
	// RMutex: 읽기 Lock -> 읽기 시도 중에 값이 변경 방지 즉, 쓰기 락 방어

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)
	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 뮤텍스 잠금
			mutex.Lock()
			data += 1
			fmt.Println("Write: ", data)
			time.Sleep(200 * time.Millisecond)
			// 쓰기 뮤텍스 해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1: ", data)
			time.Sleep(time.Second)
			mutex.RUnlock()
			// 읽기 뮤텍스 해제
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2: ", data)
			time.Sleep(time.Second)
			mutex.RUnlock()
			// 읽기 뮤텍스 해제
		}
	}()

	time.Sleep(5 * time.Second)
}
