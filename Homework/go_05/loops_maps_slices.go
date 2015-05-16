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
	var game1 = game{"Killing Floor 2", "FPS", 29.99, 82}
	var game2 = game{"Grand Theft Auto 5", "Sandbox", 59.99, 98}
	var game3 = game{"Counter-Strike: Global Offensive", "FPS", 14.99, 83}
	var game4 = game{"The Witcher 3: Wild Hunt Game", "RPG", 59.99, 89}
	var game5 = game{"Galactic Civilizations 3", "RTS", 49.99, 78}

	games := []game{game1, game2, game3, game4, game5}

	targetprice := 14.99

	for i := 0; i < 5; i++ {
		fmt.Println("searching for game loop:", i+1)
		if games[i].price == targetprice {
			fmt.Println("This Game Cost $14.99: ")
			fmt.Println(games[i].title)
			fmt.Println("Found game. Exiting loop")
			break
		}
	}
	fmt.Println("Lets Make a Map!!")
	gamemap := map[string]float64{game1.title: game1.price, game3.title: game3.price}

	for i := 0; i < 5; i++ {
		if games[i].title == game1.title || games[i].title == game3.title {
			continue
		}
		gamemap[games[i].title] = games[i].price
	}

	for key, value := range gamemap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
