package methods

import "fmt"

func EmptyInterface() {
	var i interface{}
	describev1(i)

	i = 42
	describev1(i)

	i = "hello"
	describev1(i)
}

func describev1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
