package faq

type QuestionState int

const (
	UNANSWERED QuestionState = iota
	ANSWERED
)

type Question struct {
	Title    string
	Body     string
	Views    int
	Answsers []Answer
}

func (self Question) HTML() string {
	return ""
}
