package moretypes

import "fmt"

func StructPointers() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 * 100
	fmt.Println(v)
}
