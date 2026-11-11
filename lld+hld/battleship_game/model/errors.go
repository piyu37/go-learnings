package model

import "errors"

// Sentinel errors. Callers wrap them with context via fmt.Errorf("%w: ...")
// and check them with errors.Is.
var (
	ErrInvalidSize    = errors.New("invalid size")
	ErrInvalidShip    = errors.New("invalid ship")
	ErrNotInitialized = errors.New("game not initialized")
	ErrAlreadyStarted = errors.New("game already started")
	ErrOutOfBounds    = errors.New("ship out of player's field")
	ErrOverlap        = errors.New("ship overlaps another ship")
	ErrDuplicateShip  = errors.New("duplicate ship id")
	ErrNoShips        = errors.New("no ships added")
	ErrNoTargetsLeft  = errors.New("no targets left")
)
