package variable

import "fmt"

func Const2() {
	const a, b int = 10, 20
	const c, d, e = true, 0.85, "Hello"
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 773
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)
}
