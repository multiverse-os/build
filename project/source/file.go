package source

type File struct {
	Filename     string
	Path         string
	Lines        int
	Size         int
	Language     ProgrammingLanguage
	SyntaxErrors []SyntaxError
}
