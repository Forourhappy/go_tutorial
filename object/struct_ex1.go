package object

import "fmt"

func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest}
}

func StructEx1() {
	// 구조체 생성자 패턴
	kim := Account{number: "235-901", balance: 1_000_000, interest: 0.015}

	var lee *Account = new(Account)
	lee.number = "235-902"
	lee.balance = 2_000_000
	lee.interest = 0.015

	fmt.Println(kim)
	fmt.Println(lee)

	park := NewAccount("235-903", 3_000_000, 0.015)
	fmt.Println(park)
}
