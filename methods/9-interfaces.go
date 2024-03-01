package methods

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func Interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())
	a = v
	fmt.Println(a.Abs())
}
