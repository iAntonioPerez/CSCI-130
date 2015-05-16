//Antonio Perez
package main

import (
	"fmt"
)

func max(numbers ...int) int {
	max = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}
