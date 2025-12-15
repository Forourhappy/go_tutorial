package statement

import (
	"fmt"
	"time"
)

func GoChannel3() {
	// 동기: 버퍼 미사용
	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := range cnt {
			ch <- true
			fmt.Println("Go: ", i)
			time.Sleep(time.Second) // Sleep 주석 처리 후 테스트!
		}
	}()

	for i := range cnt {
		<-ch
		fmt.Println("Main: ", i)
	}
}
