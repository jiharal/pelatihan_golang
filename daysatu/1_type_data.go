package daysatu

import "fmt"

func NumericalType() {
	var a int = 42
	var b float64 = 3.14
	var c uint = 10

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)

	// 42 int
	// 3.14 float64
	// 10 uint
}

func StringType() {
	var greeting string = "Hello, Go!"
	fmt.Printf("%v %T\n", greeting, greeting)
	// Hello, Go! string
	fmt.Println("hello word")
}

func BoleanType() {
	isTrue := true
	isFalse := false

	fmt.Printf("%v %T\n", isTrue, isTrue)
	fmt.Printf("%v %T\n", isFalse, isFalse)
	// true bool
	// false bool
}

func SpecialType() {
	var ptr *int
	var err error

	fmt.Printf("%v %T\n", ptr, ptr)
	fmt.Printf("%v %T\n", err, err)

	// <nil> *int
	// <nil> <nil>
}
