package main

import (
	"fmt"
	"os"
	"time"

	"practice/lld+hld/battleship_game/service"
	"practice/lld+hld/battleship_game/strategy"
)

func main() {
	seed := time.Now().UnixNano()
	game := service.NewGameService(os.Stdout,
		strategy.NewRandomFireStrategy(seed), strategy.NewRandomFireStrategy(seed+1))

	check(game.InitGame(5)) // invalid: N must be even
	check(game.InitGame(6))

	// Coordinates are (x = row, y = column) of the ship's center.
	check(game.AddShip("SH1", 2, 1, 1, 4, 4))
	check(game.AddShip("SH2", 2, 4, 1, 1, 4))
	check(game.AddShip("SH3", 2, 1, 3, 4, 4)) // invalid: crosses into PlayerB's half
	check(game.AddShip("SH4", 2, 1, 2, 5, 2)) // invalid: overlaps SH1 for PlayerA
	check(game.AddShip("SH1", 1, 0, 5, 5, 5)) // invalid: duplicate id

	fmt.Println("Battlefield:")
	check(game.ViewBattleField())

	winner, err := game.StartGame()
	check(err)
	fmt.Println("Winner:", winner)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
