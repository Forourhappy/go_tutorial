package condition

import "fmt"

func For2() {
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1
	}
	fmt.Println("ex1: ", sum1)

	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
		// j := i++ // Go에서 증감 연산자는 반환 값 X
	}
	fmt.Println("ex2: ", sum2)

	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3: ", sum3)

	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4: ", i, j)
	}

	/** 에러 발생
	for i, j := 0, 0; i <= 10; i++, j+=10 {
		fmt.Println("ex5: ", i, j)
	}
	*/
}
