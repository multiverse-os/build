package build

import "os"

type Binary struct {
	filename    string
	path        string
	size        int
	data        []byte
	errors      []string
	sourceFiles []SourceFile
}

func (self Binary) Execute() (*os.Process, error) {
	return nil, nil
}
