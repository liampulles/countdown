package arithmetic

// SubtractionOperation subtracts numbers
type SubtractionOperation struct{}

// Check we implement the interface
var _ Operation = &SubtractionOperation{}

// Sign implements Operation
func (so *SubtractionOperation) Sign() string {
	return "-"
}

// Commutative implements Operation
func (so *SubtractionOperation) Commutative() bool {
	return false
}

// UsableOn implements Operation
func (so *SubtractionOperation) UsableOn(a, b int) bool {
	return a-b > 0
}

// Apply implements Operation
func (so *SubtractionOperation) Apply(a, b int) int {
	return a - b
}
