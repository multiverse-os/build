package project

type TaskType int
type TaskStage int

const (
	BugTask TaskType = iota
	FeatureTask
)

const (
	UNASSIGNED TaskStage = iota
	ASSIGNED
)

type Task struct {
	Type        TaskType
	Stage       TaskStage
	Priority    Priority
	Name        string
	Description string
}
