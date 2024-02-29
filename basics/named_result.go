package basics

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	fmt.Println("x:", x)
	y = sum - x
	fmt.Println("y:", y)
	return
}

func NamedResult() {
	x, y := split(17)
	fmt.Println(x, y)
}
