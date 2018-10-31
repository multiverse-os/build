package project

import (
	docs "github.com/multiverse-os/build/project/documentation"
	source "github.com/multiverse-os/build/project/source"
	versioncontrol "github.com/multiverse-os/build/project/versioncontrol"
)

type Subproject struct {
	Name           string
	Path           string
	Version        Version
	Language       source.ProgrammingLanguage
	Documentation  docs.Documentation
	VersionControl []versioncontrol.Repository
}

func New(subproject Subproject) Subproject {
	return Subproject{
		Name:           subproject.Name,
		Path:           subproject.Path,
		Version:        subproject.Version,
		VersionControl: subproject.VersionControl,
	}
}
