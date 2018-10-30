package project

type Roadmap struct {
	Features
	Issues
}

type MajorUpdate struct {
	Codename string
	Version
	Features
	Issues
	Bugs
}

type MinorUpdate struct {
	Version
	Features
	Issues
	Bugs
}
