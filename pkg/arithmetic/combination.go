package arithmetic

import (
	"fmt"
)

// Combination represents an arithmetic combination of numbers
type Combination struct {
	a  int
	b  int
	op Operation
}

// NewCombination is a constructor
func NewCombination(a, b int, op Operation) *Combination {
	return &Combination{
		a:  a,
		b:  b,
		op: op,
	}
}

func (c *Combination) String() string {
	return fmt.Sprintf("%d %s %d = %d", c.a, c.op.Sign(), c.b, c.op.Apply(c.a, c.b))
}
