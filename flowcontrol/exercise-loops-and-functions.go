package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	fmt.Printf("Sqrt approximations of %v\n", x)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z
}

func Sqrt1(x float64) float64 {
	var z float64 = 1
	var t float64
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-6 {
			break
		}
	}
	return z
}

func ExerciseLoopsAndFunctions() {
	fmt.Println(fmt.Println(Sqrt(2)))

	fmt.Println("Next step")
	guess := Sqrt1(2)
	expected := math.Sqrt(2)
	fmt.Printf("Guess: %v, Expected: %v, Error: %v", guess, expected, math.Abs(guess-expected))
}
