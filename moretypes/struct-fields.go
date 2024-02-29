package moretypes

import "fmt"

func StructField() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
