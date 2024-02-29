package flowcontrol

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	newPow := math.Pow(x, n)
	fmt.Println(newPow)
	if v := newPow; v < lim {
		return v
	}
	return lim
}

func IfWithAShortStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
