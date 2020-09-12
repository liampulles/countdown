package arithmetic

// AdditionOperation adds numbers
type AdditionOperation struct{}

// Check we implement the interface
var _ Operation = &AdditionOperation{}

// Sign implements Operation
func (ao *AdditionOperation) Sign() string {
	return "+"
}

// Commutative implements Operation
func (ao *AdditionOperation) Commutative() bool {
	return true
}

// UsableOn implements Operation
func (ao *AdditionOperation) UsableOn(a, b int) bool {
	return true
}

// Apply implements Operation
func (ao *AdditionOperation) Apply(a, b int) int {
	return a + b
}
