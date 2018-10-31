package build

import (
	"errors"
	"os"

	docs "github.com/multiverse-os/build/project/documentation"
	source "github.com/multiverse-os/build/project/source"
	versioncontrol "github.com/multiverse-os/build/project/versioncontrol"
	volunteer "github.com/multiverse-os/build/project/volunteer"
	log "github.com/multiverse-os/log"
	watch "github.com/multiverse-os/os/fs/watch"
)

type Project struct {
	Name          string
	Path          string
	Locale        string
	Language      source.ProgrammingLanguage
	Documentation docs.Documentation
	Developers    []volunteer.Developer
	Files         []os.FileInfo
	Git           versioncontrol.Repository
	Builds        []Build
	FileHooks     map[string]*Action
}

type Action struct {
	Name     string
	Function func(event *watch.Event)
}

type ProjectHandler interface {
	InstallDependencies() error
	Clean() error
	Build() error
	Publish() error
	Tasks() error
}

func New(project Project) Project {
	// TODO: Use this to validate, transform, and test things
	if project.Name == "" {
		log.FatalError(errors.New("build project name must be specified"))
	}
	// Check if path is real
	return Project{
		Name: project.Name,
		Path: project.Path,
		Git:  project.Git,
	}
}
