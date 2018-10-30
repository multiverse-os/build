package main

type LanguageType int

const (
	Go LanguageType = iota
	Ruby
	Rust
	C
)

type Compiler struct {
	Name string
	Path string
	Version
}

type Interpreter struct {
	Name string
	Path string
	Version
}

type Language struct {
	Name string
	Type LanguageType
	Compiler
	Executable string
	Path       string
	Compiled   bool
}

func (self *Language) Compile() ([]byte, error) {
	return nil, nil
}

func (self *Language) Interpret() (string, error) {
	return nil, nil
}
