package basics

import (
	"fmt"
	"math"
)

func TypeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("%T %T %T\n", x, y, z)
}
