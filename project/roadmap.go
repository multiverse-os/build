package project

type Roadmap struct {
	Feature []Feature
	Issue   []Issue
}

type MajorUpdate struct {
	Codename string
	Version
	Feature []Feature
	Issue   []Issue
	Bug     []Bug
}

type MinorUpdate struct {
	Version
	Features []Feature
	Issues   []Issue
	Bug      []Bug
}
