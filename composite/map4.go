package composite

import "fmt"

func Map4() {
	// 맵 조회할 경우에 주의할 점

	map1 := map[string]int{
		"apple":  12,
		"banana": 114,
		"orange": 1123,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3, ok)

	if value, ok := map1["kewi"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("kiwi is not in map1")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("kiwi is not in map1")
	}

}
