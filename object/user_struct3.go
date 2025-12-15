package object

import "fmt"

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt: %d, price: %d, oderPrice: %d\n", cnt, price, fn(cnt, price))
}

func UserStruct3() {
	// 함수 사용자 정의 타입
	var oderPrice totCost
	oderPrice = func(cnt int, price int) int {
		return (cnt * price) + 100_000
	}

	describe(2, 5_000, oderPrice)

}
