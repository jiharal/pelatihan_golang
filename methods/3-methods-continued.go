package methods

import (
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	fmt.Println(f)
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MethodsContinued() {
	f := MyFloat(-5)
	fmt.Println(f.Abs())
}
