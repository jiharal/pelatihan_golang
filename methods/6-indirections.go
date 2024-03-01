package methods

import "fmt"

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Indirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	fmt.Println("--------")

	p := &Vertex{4, 3}
	p.Scale(3)
	fmt.Println(p)
	ScaleFunc(p, 8)
	fmt.Println(p)
}
