package arithmetic

// MultiplicationOperation subtracts numbers
type MultiplicationOperation struct{}

// Check we implement the interface
var _ Operation = &MultiplicationOperation{}

// Sign implements Operation
func (mo *MultiplicationOperation) Sign() string {
	return "×"
}

// Commutative implements Operation
func (mo *MultiplicationOperation) Commutative() bool {
	return true
}

// UsableOn implements Operation
func (mo *MultiplicationOperation) UsableOn(a, b int) bool {
	return true
}

// Apply implements Operation
func (mo *MultiplicationOperation) Apply(a, b int) int {
	return a * b
}
