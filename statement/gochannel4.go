package statement

import (
	"fmt"
	"runtime"
)

func GoChannel4() {
	// 비동기: 버퍼 사용
	// 버퍼: 발신 -> 가득차면 대기, 비어 있으면 작동 / 수신 -> 비어있으면 대기, 가득차 있으면 작동
	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 2)
	cnt := 12

	go func() {
		for i := range cnt {
			ch <- true
			fmt.Println("Go: ", i)
		}
	}()

	for i := range cnt {
		<-ch
		fmt.Println("Main: ", i)
	}
}
