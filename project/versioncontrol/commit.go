package versioncontrol

import (
	"time"
)

type Commit struct {
	CreatedAt time.Time
	Hash      string
	Message   string
	Size      int
}
