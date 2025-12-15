package object

import "fmt"

func Struct2() {
	/*
		다양한 선언 방법
		&struct, struct: &struct 포인터를 받아오기 역참조를 또 하기 때문에 속도는 조금 느리다
		인터페이스 메소드를 선언만 해둔 후 -> 오버라이딩 해서 메서드에 포인터 리시버를 사용할 경우 반드시 &struct 사용
	*/

	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10_000_000
	kim.interest = 0.015

	hong := &Account{
		"245-902",
		12_000_000,
		0,
	}

	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 15_000_000
	lee.interest = 0.02

	fmt.Printf("%#v\n", kim)
	fmt.Printf("%#v\n", hong)
	fmt.Printf("%#v\n", lee)
	fmt.Println()

	fmt.Println(kim.Calculate())
	fmt.Println(hong.Calculate())
	fmt.Println(lee.Calculate())
}
