package build

import "context"

type ContextHandler interface {
	context.Context
	events.Emitter

	WatchFile(string) error
	WatchFiles(string, bool) error
	IgnoreFile(string)

	AddFileHook(func() error) Context
}

type Context struct {
	WatchedFiles []string
	IgnoredFiles []string

	FileOperation FileOperation
}
