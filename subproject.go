package main

import "time"

type Subproject struct {
	name          string
	version       Version
	language      Language
	path          string
	version       Version
	lastBuild     time.Time
	documentation Documentation
	repositories  []Repository
	sourceFiles   []string
	errors        []string
	warnings      []string
	binary        Binary
}

func New(subproject Subproject) Subproject {
	return subproject
}
