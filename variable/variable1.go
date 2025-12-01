package variable

import "fmt"

func Variable1() {
	/**
	기본 초기화
	정수타입 : 0, 실수(소수점) : 0.0, 문자열 : "", Boolean : true, false
	변수명 : 숫자 첫글자x, 대소문자 구분O, 문자, 숫자, 밑줄, 특수기호 사용 가능
	변수 및 상수 : 함수 내외 사용 가능
	*/

	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!"
	var k = 4.72
	var l = "Hello, World!"
	var m = true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)

}
