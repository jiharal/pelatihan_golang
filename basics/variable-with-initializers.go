package basics

import "fmt"

var i, j int = 1, 2

func VariableWithInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println("i, j, c, python, java")
	fmt.Println(i, j, c, python, java)
	fmt.Printf("%T %T %T %T %T\n", i, j, c, python, java)
}
