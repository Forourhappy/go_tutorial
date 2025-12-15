package composite

import "fmt"

func Map2() {
	// 맵 조회 및 순회(Iterator)
	map1 := map[string]string{
		"daum":   "https://www.daum.net",
		"google": "https://www.google.com",
		"naver":  "https://www.naver.com",
	}

	// 조회
	fmt.Println(map1["daum"])
	fmt.Println(map1["google"])
	fmt.Println(map1["naver"])
	fmt.Println()

	// 순회
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println(v)
	}
}
