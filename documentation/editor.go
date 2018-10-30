package doc

type Editor struct {
	Avatar     string
	Name       string
	Background string
	Links      []Link
	Contacts   []Contact
	Questions  []Question
	Answers    []Question
	Keys       []build.Key
}

type Contact struct {
	icon     string
	platform string
	name     string
}

func (self Editor) VCard() string {
	return ""
}
