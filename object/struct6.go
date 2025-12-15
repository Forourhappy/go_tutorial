package object

import "fmt"

type Car2 struct {
	name    string `tag:"차량명"`
	color   string `tag:"색상"`
	company string `tag:"제조사"`
	detail  spec   `tag:"상세"`
}

type spec struct {
	length int `tag:"전장"`
	height int `tag:"전고"`
	width  int `tag:"전축"`
}

func Struct6() {
	car1 := Car2{
		name:    "520d",
		color:   "red",
		company: "bmw",
		detail: spec{
			length: 4000,
			height: 1700,
			width:  1800,
		},
	}

	fmt.Println(car1.name)
	fmt.Println(car1.color)
	fmt.Println(car1.company)
	fmt.Printf("%#v\n", car1.detail)

	fmt.Println(car1.detail.length)
	fmt.Println(car1.detail.height)
	fmt.Println(car1.detail.width)
}
