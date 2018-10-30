package build

import (
	"time"

	source "github.com/multiverse-os/build/project/source"
)

type Build struct {
	CompileTime   time.Time
	Path          string
	Binary        Binary
	Checksum      string
	SourceFiles   []source.File
	SyntaxErrors  []source.SyntaxError
	CompileErrors []CompileError
	RuntimeErrors []RuntimeError
}
