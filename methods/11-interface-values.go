package methods

import (
	"fmt"
	"math"
)

type I1 interface {
	M1()
}
type T1 struct {
	S string
}

func (t *T1) M1() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F1 float64

func (f F1) M1() {
	fmt.Println(f)
}

func InterfaceValues() {
	var i I1

	i = &T1{"Hello"}
	describe(i)
	i.M1()

	i = F1(math.Pi)
	describe(i)
	i.M1()
}

func describe(i I1) {
	fmt.Printf("(%v, %T)\n", i, i)
}
