package source

type EditType int

const (
	Add EditType = iota
	Remove
)

type Edit struct {
	SourceFile File
	Line       int
	Type       EditType
	Update     string
}
