package math_test

import (
	"fmt"
	"math"
)

func ExampleCeil() {
	// Ceil returns the least integer greater than or equal to x
	fmt.Println(math.Ceil(1.49))
	fmt.Println(math.Ceil(1.59))
	// Floor returns the greatest integer value less than or equal to x
	fmt.Println(math.Floor(1.49))
	fmt.Println(math.Floor(1.59))
	// Round returns the nearest integer, rounding half away from zero
	fmt.Println(math.Round(1.4))
	fmt.Println(math.Round(-1.5))
	// Output:
	// 2
	// 2
	// 1
	// 1
	// 1
	// -2
}
