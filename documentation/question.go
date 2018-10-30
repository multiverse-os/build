package doc

type QuestionType int

const (
	Unanswered QuestionType = iota
	Answered
)

type Question struct {
	Title    string
	Body     string
	Creator  Editor
	Views    int
	Answsers []Answer
	History  []Edit
}

type Answer struct {
	Rank      int
	Response  string
	Responder Editor
	History   []Edit
}

func (self Question) HTML() string {
	return ""
}

func (self Answer) HTML() string {
	return ""
}
