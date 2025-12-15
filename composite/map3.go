package composite

import "fmt"

func Map3() {
	map1 := map[string]string{
		"daum":   "https://www.daum.net",
		"google": "https://www.google.com",
		"naver":  "https://www.naver.com",
		"home1":  "https://test1.com",
	}

	fmt.Println(map1)

	// 추가
	map1["home2"] = "https://test2.com"
	fmt.Println(map1)
	fmt.Println()

	// 수정
	map1["home2"] = "https://test2-2.com"
	fmt.Println(map1)
	fmt.Println()

	delete(map1, "home2")
	fmt.Println(map1)

	delete(map1, "google")
	fmt.Println(map1)
}
