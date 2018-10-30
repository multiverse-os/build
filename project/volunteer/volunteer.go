package volunteer

type Role int

const (
	DEVELOPER Role = iota
	EDITOR
	TRANSLATOR
)

type Contributor interface {
	Name() string
	Email() string
}

type Volunteer struct {
	name  string
	email string
	roles []Role
}
