package source

type LanguageType int

const (
	GO LanguageType = iota
	RUBY
	RUST
	C
	CPP
	HTML
	CSS
	JAVASCRIPT
)

type ProgrammingLanguage struct {
	Name        string
	Type        LanguageType
	Path        string
	Binary      string
	Compiled    bool
	Interpreted bool
}
