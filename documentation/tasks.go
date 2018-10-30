package doc

type Priority int

const (
	Normal Priority = iota
	Low
	Important
	Critical
)

type TaskType int
type TaskStage int

const (
	Bug TaskType = iota
	Feature
)

const (
	Unassigned TaskStage = iota
	Assigned
)

type Task struct {
	Title       string
	Description string
	Assets      []string
	Developers  []Developer
	SourceCode  []SourceReference
	Links       []Link
	Features    []Feature
	Priority    int
	Bugs        []Bug
}
