package variable

import "fmt"

func Variable2() {
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 250
	weight = 352.2
	isRunning = true

	fmt.Println("name: ", name, "height: ", height, "weight: ", weight, "isRunning: ", isRunning)
}
