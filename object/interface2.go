package object

import "fmt"

type Dog struct {
	name   string
	weight int
}

// bite 메서드 구현
func (d Dog) bite() {
	fmt.Println(d.name, "bite!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func Interface2() {
	// 인터페이스 구현

	dog1 := Dog{"poll", 10}

	var inter1 Behavior = dog1
	// dog1.bite()
	inter1.bite()

	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)

	inter2.bite()

	inters := []Behavior{dog1, dog2}

	for idx := range inters {
		inters[idx].bite()
	}

	// 값 형태로 실행(인터페이스)
	for _, val := range inters {
		val.bite()
	}
}
