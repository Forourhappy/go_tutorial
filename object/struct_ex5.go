package object

import "fmt"

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func StructEx5() {
	// 메서드 오버라이딩 패턴
	ex := Executives{
		Employee{"lee", 5_000_000, 500_000},
		1_000_000,
	}

	fmt.Println(int(ex.Calculate()))
}
