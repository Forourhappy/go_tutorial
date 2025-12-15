package object

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func Struct1() {
	kim := Account{
		number:   "245-901",
		balance:  10_000_000,
		interest: 0.015,
	}
	lee := Account{
		number:  "245-902",
		balance: 12_000_000,
	}
	park := Account{
		number:   "245-903",
		interest: 0.025,
	}

	cho := Account{
		"245-904",
		15_000_000,
		0.03,
	}

	fmt.Println(kim)
	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(cho)

	fmt.Println()

	fmt.Println(int(kim.Calculate()))
	fmt.Println(int(lee.Calculate()))
	fmt.Println(int(park.Calculate()))
	fmt.Println(int(cho.Calculate()))
}
