package object

import "fmt"

func Struct4() {
	// 구조체 익명 선언 및 활용
	// type 구조체명 타입
	car1 := struct{ name, color string }{"520d", "red"}
	fmt.Println(car1)
	fmt.Printf("%#v\n", car1)

	cars := []struct{ name, color string }{{"520d", "red"}, {"sonata", "white"}, {"avante", "black"}}
	fmt.Println(cars)

	for _, c := range cars {
		fmt.Printf("(%s, %s)----(%#v)\n", c.name, c.color, c)
	}
}
