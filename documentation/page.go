package doc

type Page struct {
	title    string
	authors  []Developer
	assets   []string
	sections []PageSection
}

type PageSection struct {
	title   string
	content string
}

func (self Page) HTML() string {
	return ""
}
