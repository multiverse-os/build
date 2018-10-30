package versioncontrol

type Repository struct {
	Name         string
	URL          string
	Protocol     Protocol
	Contributors []Contributor
	Branch       string
	Branches     []string
	Commits      []Commit
	Files        []File
}

type RepositoryHandler interface {
	Save(message string) error
	Ignore(path string) error
	Add(path string) error
	Remove(path string) error
}

// Since add all changes to the repository automatically, save does the commit and push commands
func (self *Repository) Save(message string) error {
	// TODO: Validate message, maybe even use a ruleset file so a project can
	// define rules for commits
	// Adds occur automatically since we are watching all files for changes always
	// to rebuild and such sop we just need to commit
	return nil
}

// Add path to .gitignore in root directory
func (self *Repository) Ignore(path string) error { return nil }

// Add a file or folder and all subfolders to project repository
func (self *Repository) Add(path string) error { return nil }

// Remove file from project and rebase to remove the file from all versions
func (self *Repository) Remove(path string) error { return nil }
