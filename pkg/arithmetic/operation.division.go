package arithmetic

// DivisionOperation subtracts numbers
type DivisionOperation struct{}

// Check we implement the interface
var _ Operation = &DivisionOperation{}

// Sign implements Operation
func (do *DivisionOperation) Sign() string {
	return "÷"
}

// Commutative implements Operation
func (do *DivisionOperation) Commutative() bool {
	return false
}

// UsableOn implements Operation
func (do *DivisionOperation) UsableOn(a, b int) bool {
	if a == 0 || b == 0 {
		return false
	}
	div := a / b
	return div*b == a
}

// Apply implements Operation
func (do *DivisionOperation) Apply(a, b int) int {
	return a / b
}
