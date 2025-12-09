package statement

import "fmt"

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}

func GoChannelEx2() {
	// 채널 또한 함수의 반환 값으로 사용 가능
	c := sum(100)

	fmt.Println("1~100 합: ", <-c)

}
