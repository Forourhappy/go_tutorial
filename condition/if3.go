package condition

import "fmt"

func If3() {
	var a int = 50

	if a >= 65 {
		fmt.Println("65 이상")
	} else if a >= 50 {
		fmt.Println("50 이상")
	} else {
		fmt.Println("50 미만")
	}
}
