package packaging

import "fmt"

var num int32

func init() {
	num = 30
	fmt.Println("Main init start!")
}

func Init3() {
	fmt.Println("init 3")
}
