package statement

import (
	"fmt"
	"sync"
)

func GoSyncEx2() {
	/*
		대기 그룹
		실행 흐름 변경(고루틴이 종료될 때가지 대기 기능)
		WatingGroup: Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시까지 대기)
		Add로 추가된 고루틴 개수와 Done으로 종료되는 알림 횟수가 같아야 정확하게 동작한다.
	*/

	wg := new(sync.WaitGroup)

	for i := range 100 {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup: ", n)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("WaitGroup End!")
}
