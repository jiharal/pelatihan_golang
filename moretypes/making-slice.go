package moretypes

import "fmt"

func MakingSlice() {
	a := make([]int, 5)
	printSlice1("a", a)

	b := make([]int, 4, 5)
	printSlice1("b", b)
	newVal := []int{1, 2, 3, 4, 5}
	b = append(b, newVal...)
	printSlice1("b.a", b)

	c := b[:2]
	printSlice1("c", c)

	d := c[2:5]
	printSlice1("d", d)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
