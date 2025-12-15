package statement

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func GoSyncEx4() {
	/*
		원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작 (모두 성공하거나, 모두 실패)
		모든 조작이 완료될 때까지 다른 프로세스 개입 불가
		sync/atomic에서 원자적 연산자 제공
	*/

	// 원자성 사용하지 않는 경우
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := range 5000 {
		wg.Add(1)
		go func(n int) {
			// cnt += 1
			atomic.AddInt64(&cnt, 1)
			defer wg.Done()
		}(i)
	}

	for i := range 2000 {
		wg.Add(1)
		go func(n int) {
			atomic.AddInt64(&cnt, -1)
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("WaitGroup End! Cnt? >>>", cnt)
	fmt.Println("WaitGroup End! Cnt? >>>", finalCnt)
}
