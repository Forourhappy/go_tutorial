package statement

import (
	"fmt"
	"time"
)

func GoSync3() {
	// 뮤텍스: 상호 배제 -> Thread(고루틴)들이 서로 running time에 서로 영향을 주지 않게 즉, 단독으로 실행되게 하는 기술

	// 동기화 사용하지 않은 경우

	data := 0
	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write: ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read1: ", data)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2: ", data)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
}
