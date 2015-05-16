//Antonio Perez
package main

import (
	"fmt"
	"github.com/iAntonioPerez/calculator"
)

func main() {
	x := 5
	y := 10
	result := add(x, y)
	fmt.Println(result)
	result = subtract(x, y)
	fmt.Println(result)
	result = multiply(x, y)
	fmt.Println(result)
	result = divide(x, y)
	fmt.Println(result)
}
