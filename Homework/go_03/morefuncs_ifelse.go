//Antonio Perez

package main

import (
	"fmt"
)

type search func(game, game) string

type game struct {
	title  string
	genre  string
	price  float64
	rating int
}

func pick_func(b bool) search {
	if b == true {
		return func(g1 game, g2 game) string {
			if g1.rating >= g2.rating {
				return g1.title
			} else {
				return g2.title
			}
		}
	} else {
		return func(g1 game, g2 game) string {
			if g1.price <= g2.price {
				return g1.title
			} else {
				return g2.title
			}

		}
	}
}

func use_func(game1 game, game2 game, s search) {
	result := s(game1, game2)
	fmt.Println(result)
}

func main() {
	var game1 = game{"Killing Floor 2", "FPS", 29.99, 89}
	var game2 = game{"Grand Theft Auto 5", "Sandbox", 59.99, 98}

	fmt.Println("Highes Rated Game: ")
	f1 := pick_func(true)
	use_func(game1, game2, f1)
	f2 := pick_func(false)
	fmt.Println("Cheapest Game: ")
	use_func(game1, game2, f2)
}
