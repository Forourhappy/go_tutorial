package datatype

import (
	"fmt"
	"math"
)

func NumberOperation1() {
	/*
		산술 연산(산술, 비교)
		타입이 같아야 가능
		다른 타입끼리는 반드시 형 변환 후 계산
		형 변환 없을 경우 예외(에러) 발생
		+, -, *, /, %, <<, >>, &, ^
	*/

	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	n5 := 100_000
	n6 := int16(10000)

	// fmt.Println(n5 + n6) // 예외 발생
	fmt.Println(n5 + int(n6))

}
