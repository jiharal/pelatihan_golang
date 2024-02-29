package moretypes

import "fmt"

func SlicesPointer() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:3]
	b := names[1:4]
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	b[1] = "YYY"
	fmt.Println("new a:", a)
	fmt.Println("new b:", b)
	fmt.Println(names)
}
