package moretypes

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Struct() {
	fmt.Println(Vertex{1, 2})
}
