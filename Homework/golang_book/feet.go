//Antonio Perez

package main

import (
	"fmt"
)

func convert(feet float32) float32 {
	return feet * 0.3048
}

func main() {
	var feet float32 = 32
	meter := convert(feet)
	fmt.Println(meter)
}
