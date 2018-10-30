package versioncontrol

type Protocol int

const (
	Git Protocol = iota
	Mercurial
	SVN
	Fossil
)
