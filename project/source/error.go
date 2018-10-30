package source

type SyntaxError struct {
	SourceFile File
	Line       int
	Message    string
	Syntax     bool
	Warning    bool
}

// TODO: Here we will use YAML to define code generation templates to make our
// development process much faster
