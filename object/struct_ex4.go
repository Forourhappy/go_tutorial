package object

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee     // is a 관계
	specialBonus float64
}

func StructEx4() {
	/*
		구조체 임베디드 패턴
		다른 관점으로 메소드를 재사용하는 장점 제공
		상속을 허용하지 않는 Go 언어에서 메서드 상속 활용을 위한 패턴
	*/

	ep1 := Employee{"kim", 2_000_000, 200_000}
	ep2 := Employee{"lee", 1_500_000, 300_000}

	ex := Executives{
		Employee{"lee", 5_000_000, 500_000},
		1_000_000,
	}

	fmt.Println(int(ep1.Calculate()))
	fmt.Println(int(ep2.Calculate()))

	// Employee 구조체 통해서 메서드 호출
	fmt.Println(int(ex.Calculate() + ex.specialBonus))
}
