package build

import "time"

type Build struct {
	CompileTime time.Time
	Binary      Binary
	Version     Version
	Checksum    string
}
