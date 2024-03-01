package methods

import "fmt"

func MethodWithPointerReceiver() {
	v := &Vertex{3, 4}
	fmt.Printf("Sebelum scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("Setelah scaling %+v, Abs: %v\n", v, v.Abs())
}
