package games

import (
	"dice-arcade/internal/dice"
	"fmt"
)

type Other struct{}

func (Other) Name() string { return "other" }

func (Other) PlayOnce() string {
	n := dice.D6()
	if n == 1 {
		return fmt.Sprintf("Other: rolled %d → BUST", n)
	}
	return fmt.Sprintf("Other: rolled %d → +points", n)
}
