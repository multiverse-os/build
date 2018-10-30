package main

import (
	"fmt"
	"os"

	build "github.com/multiverse-os/build"
	versioncontrol "github.com/multiverse-os/build/project/versioncontrol"
	log "github.com/multiverse-os/log"
	file "github.com/multiverse-os/os/file"
	terminal "github.com/multiverse-os/os/terminal"
)

func main() {
	var path string
	terminal.PrintBanner("[WATCH] developer file watch", "v0.1.0")

	if len(os.Args) > 1 {
		path = os.Args[1]
		log.Info("os.Args[1]: " + path + "'")
		if path == "." {
			log.Info("path is '.', need to convert to absolute path")
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
