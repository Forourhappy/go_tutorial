package object

import "fmt"

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateF(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func StructEx3() {
	kim := Account{"235-901", 1_000_000, 0.015}
	lee := Account{"235-902", 2_000_000, 0.015}

	fmt.Println(int(kim.balance))
	fmt.Println(int(lee.balance))

	fmt.Println()

	kim.CalculateD(1_000_000)
	lee.CalculateF(1_000_000)

	fmt.Println(int(kim.balance))
	fmt.Println(int(lee.balance))

}
