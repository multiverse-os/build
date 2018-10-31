package faq

type Answer struct {
	Rank     int
	Response string
}

func (self Answer) HTML() string {
	return ""
}
