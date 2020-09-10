package command

// Command defines a funcion which takes command line args and processes
// them - potentially encountering an error.
type Command func(args []string) error
