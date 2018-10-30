package build

type SourceFile struct {
	filename string
	path     string
	lines    int
	size     int
	language Language
	warnings []string
	errors   []string
}

type SourceError struct {
	Line    int
	Message string
	Syntax  bool
	Compile bool
}

// TODO: Here we will use YAML to define code generation templates to make our
// development process much faster
