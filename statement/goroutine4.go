package statement

import (
	"fmt"
	"runtime"
	"time"
)

func Goroutine4() {
	// 클로저 사용 예제
	runtime.GOMAXPROCS(1)

	s := "Goroutine Cloure: "
	for i := range 1000 {
		//  반복분 클로저는 일반적으로 즉시 실행 But 고루틴 클로저는 가장 나중에 실행(반복문 종료 후)
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i)
	}

	time.Sleep(time.Second * 5)
}
