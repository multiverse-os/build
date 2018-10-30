package main

import (
	"fmt"
	"log"
	"time"

	watch "github.com/multiverse-os/os/file/watch"
)

type Project struct {
	Name           string
	URL            string
	RootKey        Keypair
	Path           string
	language       string
	builds         []Build
	versionControl VersionControl
	repositories   []Repository
	developers     Developer
	sourceFiles    []SourceFile
	binaries       []Binary
	documentation  Documentation
	subprojects    []Subproject
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
	return project
}

func (self *Project) Watch() {
	w := watch.New()
	w.RateLimit(1)

	go func() {
		for {
			select {
			case event := <-w.Events:
				bash(`go build ` + project.Path)
			case err := <-w.Errors:
				log.Fatalln(err)
			case <-w.Shutdown:
				return
			}
		}
	}()

	if err := w.AddRecursive("."); err != nil {
		log.Fatalln(err)
	}

	for path, file := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, file.Name())
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
