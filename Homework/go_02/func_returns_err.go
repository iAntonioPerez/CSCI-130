//Antonio Perez

package main

import (
	"fmt"
)

const (
	message = "Title: "
)

type game struct {
	title  string
	genre  string
	price  float64
	rating int
}

func get_title_price(g game) (string, float64) {
	return g.title, g.price
}

func print_game(title string) {
	fmt.Println(message, title)
}

func main() {
	var game1 = game{"Killing Floor 2", "FPS", 29.99, 89}
	var game2 = game{"Grand Theft Auto 5", "Sandbox", 59.99, 98}

	title1, _ := get_title_price(game1)
	title2, _ := get_title_price(game2)
	print_game(title1)
	print_game(title2)

}
