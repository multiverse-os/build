package build

import (
	source "github.com/multiverse-os/build/project/source"
)

type Compiler struct {
	Executable string
	Path       string
}

func (self *Project) Compile() ([]byte, error) {
	return nil, nil
}

type CompileError struct {
	SourceFile source.File
	Line       int
	Message    int
}
