package flowcontrol

import "fmt"

func ForContinued() {
	sum := 1
	// ignore
	for ; sum < 1000; {
		sum += sum
		fmt.Println("sum:", sum)
	}

	fmt.Println(sum)
}
