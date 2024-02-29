package moretypes

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func ValueFunction() {
	mul := func(x, y float64) float64 {
		return x * y
	}
	fmt.Println(mul(10, 10))
	// x := mul(10, 10)

	fmt.Println(compute(mul))
	// fmt.Println(compute(math.Pow))

	hypot := func(x, y float64) float64 {
		i := math.Sqrt(x*x + y*y)
		fmt.Println(i)
		return x
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))

}
