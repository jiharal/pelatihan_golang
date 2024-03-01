package methods

import "fmt"

func NilInterfaceValues() {
	fmt.Println("---------")
	var x I1
	describe(x)
	x.M1()
}
