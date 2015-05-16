//Antonio Perez

package main

import (
	"fmt"
	"time"
)

type game struct {
	title  string
	genre  string
	price  float64
	rating int
}

func (g game) game_method() string {
	return g.title
}

func main() {
	//Time Package
	current := time.Now()

	fmt.Println("Current TIme: ", current)

	var game1 = game{"Killing Floor 2", "FPS", 29.99, 82}
	var game2 = game{"Grand Theft Auto 5", "Sandbox", 59.99, 98}
	var game3 = game{"Counter-Strike: Global Offensive", "FPS", 14.99, 83}
	var game4 = game{"The Witcher 3: Wild Hunt Game", "RPG", 59.99, 89}
	var game5 = game{"Galactic Civilizations 3", "RTS", 49.99, 78}

	var games []game
	games = make([]game, 3, 3)

	games[0] = game1
	games[1] = game2
	games[2] = game3

	fmt.Println("The length of our slice is:", len(games))
	fmt.Println("These are the games in our slice:")
	for i := 0; i < len(games); i++ {
		temp := games[i].game_method()
		fmt.Println("	", temp)
	}
	fmt.Println("Lets append to the slice.")

	games = append(games, game4, game5)
	fmt.Println("The length of our slice is now:", len(games))
	fmt.Println("These are the games in that are now in our slice:")
	for i := 0; i < len(games); i++ {
		temp := games[i].game_method()
		fmt.Println("	", temp)
	}
	fmt.Println("Now lets delete from the slice.")

	games = games[:2]
	fmt.Println("The length of our slice is now:", len(games))
	fmt.Println("These are the games that are left in our slice:")
	for i := 0; i < len(games); i++ {
		temp := games[i].game_method()
		fmt.Println("	", temp)
	}
}
