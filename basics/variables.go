package basics

import "fmt"

var c, python, java bool

func Variables() {
	var i int
	fmt.Println(i, c, python, java)
	fmt.Printf("%T %T %T %T\n", i, c, python, java)
}
