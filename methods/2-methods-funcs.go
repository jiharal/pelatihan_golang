package methods

import (
	"fmt"
	"math"
)

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodsFunc() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
