//Antonio Perez
package main

func minimum(numbers []float64) float64 {
	min := numbers[0]
	for i := 1; i < len(numbers); x++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}

func maximum(numbers []float64) float64 {
	max := numbers[0]
	for i := 1; i < len(numbers); x++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min
}
