package statement

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func onceTest() {
	fmt.Println("Once Test")
}

func GoSyncEx1() {
	/*
		Once: 한 번만 실행(주로 초기화에 사용)
		Do로 실행
	*/

	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once) // Once 객체 생성

	for i := range 5 {
		go func(n int) {
			fmt.Println("Goroutine: ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
