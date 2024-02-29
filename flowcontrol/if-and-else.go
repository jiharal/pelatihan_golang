package flowcontrol

import (
	"fmt"
	"math"
)

func pow1(x, n, lim float64) float64 {
	newPow := math.Pow(x, n)
	fmt.Println(newPow)
	if v := newPow; v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func IfAndElse() {
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)
}
