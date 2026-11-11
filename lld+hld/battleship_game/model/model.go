// Package model is the data layer: plain in-memory structs with no game rules.
package model

import "fmt"

// Coordinate is a grid cell. X is the row, Y is the column (both 0-based).
type Coordinate struct {
	X, Y int
}

// String formats the coordinate as "(x, y)"; fmt calls it automatically for %v.
func (c Coordinate) String() string { return fmt.Sprintf("(%d, %d)", c.X, c.Y) }

// Ship is a square ship occupying rows [MinX..MaxX] and columns [MinY..MaxY].
type Ship struct {
	ID                     string
	MinX, MinY, MaxX, MaxY int
	Destroyed              bool
}

// NewShip builds a square ship of the given size centered at (x, y).
// Example: size 4 at (2, 2) occupies cells 0..3 on both axes,
// i.e. its corners are (0, 0) and (4, 4).
func NewShip(id string, size, x, y int) *Ship {
	minX, minY := x-size/2, y-size/2
	return &Ship{
		ID:   id,
		MinX: minX, MinY: minY,
		MaxX: minX + size - 1, MaxY: minY + size - 1,
	}
}

// Overlaps reports whether two ships share at least one cell.
// Ships that only touch boundaries do not overlap.
func (s *Ship) Overlaps(o *Ship) bool {
	return s.MinX <= o.MaxX && o.MinX <= s.MaxX && s.MinY <= o.MaxY && o.MinY <= s.MaxY
}

// Player owns all rows and the columns [MinY..MaxY] of the battlefield,
// plus a fleet of ships.
type Player struct {
	Name       string
	MinY, MaxY int                  // first and last column of this player's half
	Ships      []*Ship              // fleet, in the order ships were added
	cells      map[Coordinate]*Ship // every occupied cell -> ship, for O(1) hit lookup
	alive      int                  // number of ships not yet destroyed
}

// NewPlayer creates a player with an empty fleet who owns columns minY..maxY.
func NewPlayer(name string, minY, maxY int) *Player {
	return &Player{Name: name, MinY: minY, MaxY: maxY, cells: make(map[Coordinate]*Ship)}
}

// AddShip stores the ship in the fleet and maps every cell it covers to it.
// It does no validation; the service checks bounds and overlaps first.
func (p *Player) AddShip(s *Ship) {
	p.Ships = append(p.Ships, s)
	p.alive++
	for x := s.MinX; x <= s.MaxX; x++ {
		for y := s.MinY; y <= s.MaxY; y++ {
			p.cells[Coordinate{x, y}] = s
		}
	}
}

// ShipAt returns the ship covering cell c, or false if the cell is empty (O(1)).
func (p *Player) ShipAt(c Coordinate) (*Ship, bool) {
	s, ok := p.cells[c]
	return s, ok
}

// HasShip reports whether the fleet already has a ship with this id.
func (p *Player) HasShip(id string) bool {
	for _, s := range p.Ships {
		if s.ID == id {
			return true
		}
	}
	return false
}

// Destroy marks the ship destroyed and decrements the alive count.
// Calling it again on an already destroyed ship does nothing.
func (p *Player) Destroy(s *Ship) {
	if !s.Destroyed {
		s.Destroyed = true
		p.alive--
	}
}

// AllDestroyed reports whether every ship in the fleet is destroyed (player has lost).
func (p *Player) AllDestroyed() bool { return p.alive == 0 }

// Battlefield is the NxN sea. Players[0] (A) owns the left half (columns
// 0..N/2-1), Players[1] (B) owns the right half (columns N/2..N-1).
type Battlefield struct {
	N       int
	Players [2]*Player
	Fired   map[Coordinate]bool // every coordinate a missile has hit, across the whole game
}

// NewBattlefield creates an empty NxN battlefield, giving PlayerA the left
// half of the columns and PlayerB the right half.
func NewBattlefield(n int) *Battlefield {
	return &Battlefield{
		N: n,
		Players: [2]*Player{
			NewPlayer("PlayerA", 0, n/2-1),
			NewPlayer("PlayerB", n/2, n-1),
		},
		Fired: make(map[Coordinate]bool),
	}
}
