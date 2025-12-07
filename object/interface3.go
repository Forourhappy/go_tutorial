package object

import "fmt"

type Cat struct {
	name   string
	weight int
}

func (d Dog) sounds() {
	fmt.Println(d.name, "dog sounds!")
}

func (d Dog) run() {
	fmt.Println(d.name, "dog run!")
}

func (c Cat) bite() {
	fmt.Println(c.name, "cat bite!")
}

func (c Cat) sounds() {
	fmt.Println(c.name, "cat meow!")
}

func (c Cat) run() {
	fmt.Println(c.name, "cat run!")
}

// 동물의 행동 인터페이스 선언
type Behavior2 interface {
	bite()
	sounds()
	run()
}

// 인터페이스 파라미터 받는다
func act(animal Behavior2) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func Interface3() {
	/*
		인터페이스 구현 예제
		인터페이스 규격화 역할 이해
		인터페이스에 정의된 메서드 사용 유도
		코드의 가독성 및 유지보수 증가

		덕타이핑 예제
		덕타이핑: 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메서드로만 판단하는 방식
		Go의 덕타이핑 중요한 특징: 오리처럼 걷고, 소리내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다. -> 의미
	*/

	dog := Dog{"poll", 10}
	cat := Cat{"marry", 12}

	// 개 행동 실행
	act(dog)

	// 고양이 행동 실행
	act(cat)
}
