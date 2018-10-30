package project

type TaskType int
type TaskStage int

const (
	Bug TaskType = iota
	Feature
)

const (
	UNASSIGNED TaskStage = iota
	ASSIGNED
)

type Task struct {
	Type        TaskType
	Stage       TaskStage
	Priority    Priority
	Title       string
	Description string
}
