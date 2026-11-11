// Package service is the service layer: all game rules, turns and console output.
package service

import (
	"fmt"
	"io"
	"strings"

	"practice/lld+hld/battleship_game/model"
	"practice/lld+hld/battleship_game/strategy"
)

type GameService struct {
	bf         *model.Battlefield
	strategies [2]strategy.FireStrategy // index 0 = PlayerA, 1 = PlayerB
	out        io.Writer
	started    bool
}

func NewGameService(out io.Writer, strategyA, strategyB strategy.FireStrategy) *GameService {
	return &GameService{out: out, strategies: [2]strategy.FireStrategy{strategyA, strategyB}}
}

// InitGame creates an NxN battlefield; left half to PlayerA, right half to PlayerB.
func (g *GameService) InitGame(n int) error {
	if n < 2 || n%2 != 0 {
		return fmt.Errorf("%w: N must be an even number >= 2, got %d", model.ErrInvalidSize, n)
	}
	g.bf = model.NewBattlefield(n)
	g.started = false
	return nil
}

// AddShip adds a ship with the same id and size to both fleets. (x, y) is the
// ship's center. Both placements are validated before either is saved, so the
// fleets always stay equal.
func (g *GameService) AddShip(id string, size, xA, yA, xB, yB int) error {
	if g.bf == nil {
		return model.ErrNotInitialized
	}
	if g.started {
		return fmt.Errorf("%w: cannot add ship %q", model.ErrAlreadyStarted, id)
	}
	if id == "" {
		return fmt.Errorf("%w: ship id must not be empty", model.ErrInvalidShip)
	}
	if size <= 0 {
		return fmt.Errorf("%w: ship size must be > 0, got %d", model.ErrInvalidSize, size)
	}
	if g.bf.Players[0].HasShip(id) {
		return fmt.Errorf("%w: %q", model.ErrDuplicateShip, id)
	}

	ships := [2]*model.Ship{model.NewShip(id, size, xA, yA), model.NewShip(id, size, xB, yB)}
	for i, p := range g.bf.Players {
		if err := g.validatePlacement(p, ships[i]); err != nil {
			return err
		}
	}
	for i, p := range g.bf.Players {
		p.AddShip(ships[i])
	}
	return nil
}

func (g *GameService) validatePlacement(p *model.Player, s *model.Ship) error {
	if s.MinY < p.MinY || s.MaxY > p.MaxY || s.MinX < 0 || s.MaxX >= g.bf.N {
		return fmt.Errorf("%w: %s's ship %q spans (%d,%d)-(%d,%d), field is rows x[0..%d] cols y[%d..%d]",
			model.ErrOutOfBounds, p.Name, s.ID, s.MinX, s.MinY, s.MaxX, s.MaxY, g.bf.N-1, p.MinY, p.MaxY)
	}
	for _, other := range p.Ships {
		if s.Overlaps(other) {
			return fmt.Errorf("%w: %s's ship %q overlaps %q", model.ErrOverlap, p.Name, s.ID, other.ID)
		}
	}
	return nil
}

// StartGame plays until one player loses all ships. PlayerA fires first and the
// turn always passes to the other player. Returns the winner's name.
func (g *GameService) StartGame() (string, error) {
	if g.bf == nil {
		return "", model.ErrNotInitialized
	}
	if g.started {
		return "", model.ErrAlreadyStarted
	}
	if len(g.bf.Players[0].Ships) == 0 {
		return "", model.ErrNoShips
	}
	g.started = true

	for turn := 0; ; turn = 1 - turn {
		attacker, defender := g.bf.Players[turn], g.bf.Players[1-turn]

		target, err := g.strategies[turn].NextTarget(defender.MinY, defender.MaxY, g.bf.N, g.bf.Fired)
		if err != nil {
			return "", fmt.Errorf("%s could not fire: %w", attacker.Name, err)
		}
		if g.bf.Fired[target] || target.Y < defender.MinY || target.Y > defender.MaxY || target.X < 0 || target.X >= g.bf.N {
			return "", fmt.Errorf("%s's strategy returned invalid target %v", attacker.Name, target)
		}
		g.bf.Fired[target] = true

		g.fire(attacker, defender, target)

		if defender.AllDestroyed() {
			fmt.Fprintf(g.out, "GameOver. %s wins.\n", attacker.Name)
			return attacker.Name, nil
		}
	}
}

func (g *GameService) fire(attacker, defender *model.Player, target model.Coordinate) {
	prefix := fmt.Sprintf("%s's turn: Missile fired at %v.", attacker.Name, target)
	ship, hit := defender.ShipAt(target)
	switch {
	case !hit:
		fmt.Fprintf(g.out, "%s \"Miss\"\n", prefix)
	case ship.Destroyed:
		fmt.Fprintf(g.out, "%s \"Hit\". %s's ship with id %q already destroyed.\n", prefix, defender.Name, ship.ID)
	default:
		defender.Destroy(ship)
		fmt.Fprintf(g.out, "%s \"Hit\". %s's ship with id %q destroyed.\n", prefix, defender.Name, ship.ID)
	}
}

// ViewBattleField prints the NxN grid, labelling every cell a ship occupies
// (A-SH1, B-SH2, ...). Row x = 0 is printed first.
func (g *GameService) ViewBattleField() error {
	if g.bf == nil {
		return model.ErrNotInitialized
	}
	const width = 8
	var sb strings.Builder
	for x := 0; x < g.bf.N; x++ { // rows
		for y := 0; y < g.bf.N; y++ { // columns
			label := "."
			for i, p := range g.bf.Players {
				if s, ok := p.ShipAt(model.Coordinate{X: x, Y: y}); ok {
					label = string(rune('A'+i)) + "-" + s.ID
				}
			}
			fmt.Fprintf(&sb, "%-*s", width, label)
		}
		sb.WriteString("\n")
	}
	_, err := io.WriteString(g.out, sb.String())
	return err
}
