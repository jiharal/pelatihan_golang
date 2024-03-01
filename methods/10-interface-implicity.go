package methods

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func InterfaceImplicity() {
	var i I = T{"hello"}
	i.M()
}
