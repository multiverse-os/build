package project

type Bug struct {
	Name             string
	Issues           []Issue
	Path             string
	AffectedVersions []Version
	Reproducible     bool
	Resolved         bool
	VersionResolved  Version
}
