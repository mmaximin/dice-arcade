package games

import (
	"dice-arcade/internal/dice"
	"fmt"
)

type Other struct{}

func (Other) Name() string { return "other" }

func (Other) PlayOnce() string {
	n := dice.D6()
	n1 := dice.D6()
	if n == n1 {
		return fmt.Sprintf("Other: rolled %d and %d → BUST", n, n1)
	}
	return fmt.Sprintf("Other: rolled %d and %d → +points", n, n1)
}
