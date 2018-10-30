package versioncontrol

type File struct {
	Extension    string
	Language     ProgrammingLanguage
	Filename     string
	Path         string
	Contributors []string
	Commits      []Commit
	Lines        int
	Size         int
}
