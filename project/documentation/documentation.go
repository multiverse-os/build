package docs

import (
	"time"

	faq "github.com/multiverse-os/build/project/documentation/faq"
	volunteer "github.com/multiverse-os/build/project/volunteer"
)

type Documentation struct {
	Tile         string
	Contributors []volunteer.Contributor
	Description  string
	FAQ          []faq.Question
	Wiki         []Page
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
