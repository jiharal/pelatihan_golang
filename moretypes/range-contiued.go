package moretypes

import "fmt"

func RangeContinued() {
	pow := make([]int, 10)
	fmt.Println(pow)
	for i := range pow {
		x := 1 << uint(i) // == 2**i (operasi bitwise)
		fmt.Println(x)
		pow[i] = x
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
