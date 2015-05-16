package main

import (
	"fmt"
)

func convert(temp float32) float32 {
	return (temp - 32) * 5 / 9
}

func main() {
	var fah float32 = 52
	cel := convert(fah)
	fmt.Println(cel)
}
