//Antonio Perez
package main

import (
	"fmt"
)

type game struct {
	title  string
	genre  string
	price  float64
	rating int
}

func main() {
	var game1 = game{"Killing Floor 2", "FPS", 29.99, 89}

	var title string = "Grand Theft Auto 5"
	genre := "Sandbox"
	price := 59.99
	var rating int = 98
	var game2 = game{title, genre, price, rating}

	var msg string = "Here are the available games:"
	var display *string = &msg
	fmt.Println(msg)
	*display = "GAME 1"
	fmt.Println(msg)
	fmt.Println("Title: ", game1.title, "	Genre:  ", game1.genre, "	Price: ", game1.price, "	Rating: ", game1.rating)
	*display = "GAME 2"
	fmt.Println(msg)
	fmt.Println("Title: ", game2.title, "	Genre:  ", game2.genre, "	Price: ", game2.price, "	Rating: ", game2.rating)

}
