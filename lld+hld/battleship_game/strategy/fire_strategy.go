// Package strategy holds the pluggable missile-targeting strategies.
package strategy

import (
	"math/rand"

	"practice/lld+hld/battleship_game/model"
)

// FireStrategy picks the next coordinate to fire at inside the opponent's
// field (rows 0..n-1, columns minY..maxY). It must not return a fired coordinate.
// Add a new strategy by implementing this interface.
type FireStrategy interface {
	NextTarget(minY, maxY, n int, fired map[model.Coordinate]bool) (model.Coordinate, error)
}

// RandomFireStrategy fires at a random coordinate that hasn't been fired yet.
type RandomFireStrategy struct {
	rng *rand.Rand
}

// NewRandomFireStrategy takes a seed so games can be reproduced in tests.
func NewRandomFireStrategy(seed int64) *RandomFireStrategy {
	return &RandomFireStrategy{rng: rand.New(rand.NewSource(seed))}
}

func (r *RandomFireStrategy) NextTarget(minY, maxY, n int, fired map[model.Coordinate]bool) (model.Coordinate, error) {
	var candidates []model.Coordinate
	for x := 0; x < n; x++ { // every row
		for y := minY; y <= maxY; y++ { // opponent's columns only
			if c := (model.Coordinate{X: x, Y: y}); !fired[c] {
				candidates = append(candidates, c)
			}
		}
	}
	if len(candidates) == 0 {
		return model.Coordinate{}, model.ErrNoTargetsLeft
	}
	return candidates[r.rng.Intn(len(candidates))], nil
}
