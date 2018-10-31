package readme

type File struct {
	title        string
	markdown     bool
	contributors []Editor
	description  string
	sections     []PageSection
	history      []Edit
}

func (self Readme) String() string {
	return ""
}

func (self Readme) HTML() string {
	return ""
}

func (self Readme) Markdown() string {
	return ""
}
