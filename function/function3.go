package function

import "fmt"

func multiply(x int, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func Function3() {
	// 다중 값 반환

	a, b := multiply(10, 5)
	fmt.Println(a, b)

	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	fmt.Println(x1, x2, x3, x4, x5)
}
