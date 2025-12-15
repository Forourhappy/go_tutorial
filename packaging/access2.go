package packaging

import (
	"fmt"

	checkUp "go_tutorial/packaging/lib"
	_ "go_tutorial/packaging/lib2"
)

func Access2() {
	/**
	패키지 접근제어
	별칭 사용
	빈 식별자 사용
	*/

	fmt.Println("10보다 큰 수?: ", checkUp.CheckNum(11))
}
