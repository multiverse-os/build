package issues

type Issues []Issue

type Issue struct {
	Title    string
	Body     string
	Comments []string
}
