package project

import (
	versioncontrol "github.com/multiverse-os/build/project/versioncontrol"
)

type Subproject struct {
	Name           string
	Path           string
	Version        Version
	Language       Language
	Documentation  Documentation
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
