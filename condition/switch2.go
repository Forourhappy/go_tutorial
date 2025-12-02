package condition

import (
	"fmt"
	"math/rand"
	"time"
)

func Switch2() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	switch i := rand.Intn(100); {
	case i < 100:
		fmt.Println("100 미만", i)
	case i < 50:
		fmt.Println("50 미만", i)
	case i < 25:
		fmt.Println("25 미만", i)
	default:
		fmt.Println("i", i)
	}
}
