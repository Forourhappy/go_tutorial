package object

import "fmt"

type shoppingBasket struct {
	cnt,
	price int
}

// 결제 함수
func (d shoppingBasket) purchase() int {
	return d.cnt * d.price
}

// 원본 수정(참조 형식)
func (d *shoppingBasket) rePurchageP(cnt, price int) {
	d.cnt += cnt
	d.price += price
}

// 원본 수정X(값 전달 형식)
func (d shoppingBasket) rePurchaseD(cnt, price int) {
	d.cnt += cnt
	d.price += price
}

func UserStruct4() {
	/*
		리시버(구조체 메서드) 전달(값, 참조) 형식
		함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달
		리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능
	*/

	bs1 := shoppingBasket{3, 5000}
	fmt.Println("bs1: ", bs1.purchase())

	// 참조 전달(원본 값 수정)
	bs1.rePurchageP(7, 5000)
	fmt.Println("bs1: ", bs1.purchase())

	bs1.rePurchaseD(10, 0)
	fmt.Println("bs1: ", bs1.purchase())
}
