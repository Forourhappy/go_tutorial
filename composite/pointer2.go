package composite

import "fmt"

func Pointer2() {
	i := 7
	p := &i

	fmt.Println(i, *p, &i, p)

	*p++

	fmt.Println(i, *p, &i, p)

	*p = 77777

	fmt.Println(i, *p, &i, p)

	i = 77

	fmt.Println(i, *p, &i, p)
}
