package doc

type Readme struct {
	title        string
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
