package arithmetic

var operations = []Operation{
	&AdditionOperation{},
	&SubtractionOperation{},
	&MultiplicationOperation{},
	&DivisionOperation{},
}

// Operation defines a form of arithmetic which can be applie don two ints
type Operation interface {
	Sign() string
	Commutative() bool
	UsableOn(a, b int) bool
	Apply(a, b int) int
}
