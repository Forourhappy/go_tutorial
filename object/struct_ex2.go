package object

import "fmt"

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func StructEx2() {
	kim := Account{"235-901", 1_000_000, 0.015}
	lee := Account{"235-902", 2_000_000, 0.015}

	fmt.Println(kim)
	fmt.Println(lee)

	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println(int(kim.balance))
	fmt.Println(int(lee.balance))

}
