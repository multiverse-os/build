package docs

type Page struct {
	title    string
	sections []PageSection
}

type PageSection struct {
	title   string
	content string
}

func (self Page) HTML() string {
	return ""
}
