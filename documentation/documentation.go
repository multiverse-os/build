package doc

import "time"

type Documentation struct {
	tile         string
	contributors []Editor
	description  string
	links        []Link
	contacts     []Contact
	faq          []Question
	pages        []Page
	issues       []Issues
	bugs         []Bug
	roadmap      []Feature
	tasks        []Task
	features     []Feature
	template     func(string) string
}

type Link struct {
	icon      string
	name      string
	url       string
	createdAt time.Time
}

type Feature struct {
	Title        string
	Description  string
	Developers   string
	Links        []Link
	Submitted    bool
	PeerReviewed bool
	Completed    bool
}

func (self Documentation) HTML() string {
	return ""
}
