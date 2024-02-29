package basics

import "fmt"

func ShortVariableDeclarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println("i, j, k, c, python, java")
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("%T %T %T %T %T %T\n", i, j, k, c, python, java)
}
