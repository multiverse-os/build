package build

import "os"

type Binary struct {
	filename string
	path     string
	size     int
	data     []byte
}

type RuntimeError struct {
	Message string
}

func (self Binary) Execute() (*os.Process, error) {
	return nil, nil
}
