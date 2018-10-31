package main

import (
	"fmt"
	"os"
	"path/filepath"

	build "github.com/multiverse-os/build"
	project "github.com/multiverse-os/build/project"
	versioncontrol "github.com/multiverse-os/build/project/versioncontrol"
	log "github.com/multiverse-os/log"
	file "github.com/multiverse-os/os/fs/file"
	terminal "github.com/multiverse-os/os/terminal"
)

func main() {
	var path string
	version := project.Version{Major: 0, Minor: 1, Patch: 0}
	terminal.PrintBanner((terminal.Light("[WATCH]") + terminal.LightPurple(" Developer File Event Hook")), terminal.Strong(version.String()))

	if len(os.Args) > 1 {
		path = os.Args[1]
		log.Info("os.Args[1]: " + path + "'")
		log.Info("path is '.', need to convert to absolute path")
		path, err := filepath.Abs(path)
		if err != nil {
			log.Fatal("could not parse path")
			os.Exit(1)
		}
		if file.Exists(path) {
			log.Info("build project path exists: '" + path + "'")
		} else {
			log.Fatal("specified project path does not exist, exiting...")
			os.Exit(1)
		}
	} else {
		log.Fatal(" watch cli requires project path: 'e.g. watch {PATH}'")
		fmt.Println(terminal.Strong("  USAGE ") + terminal.Light("watch /home/user/Developer/cli") + "\n")
		os.Exit(1)
	}
	project := build.Project{
		Name: "cli",
		Path: path,
		Git: versioncontrol.Repository{
			Protocol: versioncontrol.Git,
			URL:      "https://github.com/multiverse-os/cli",
			Branch:   "master",
		},
	}
	log.Info("Watching project files for changes...")
	project.Watch()
}
