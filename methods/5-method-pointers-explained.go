package methods

import "fmt"

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodsPointersExplained() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
