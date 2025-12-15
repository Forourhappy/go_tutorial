package object

func act2(animal interface{ run() }) {
	animal.run()
}

func Interface4() {
	dog := Dog{"poll", 10}
	cat := Cat{"marry", 12}

	// 개 행동 실행
	act2(dog)

	// 고양이 행동 실행
	act2(cat)
}
