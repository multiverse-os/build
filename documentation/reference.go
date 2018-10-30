package doc

type Reference struct {
	title    string
	content  string
	fromLine int
	toLine   int
}

type RemoteReference struct {
	link      Link
	reference Reference
}

type SourceReference struct {
	source    SourceFile
	reference Reference
}
