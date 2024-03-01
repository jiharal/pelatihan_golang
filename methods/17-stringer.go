package methods

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func Stringer() {
	a := Person{"Jhon", 32}
	z := Person{"Zain", 9001}
	fmt.Println(a, z)
}
