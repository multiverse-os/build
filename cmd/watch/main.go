package main

import term "github.com/multiverse-os/os/terminal"

func main() {
	term.PrintBanner("build project watch", "v0.1.0")
	project := New(Project{
		Name: "cli-framework",
		Path: "~/Development/cli",
		Git:  "https://github.com/multiverse-os/cli",
	})
}
