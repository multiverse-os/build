package main

type KeyType int

const (
	GPG KeyType = iota
	SSH
	Bitcoin
)

type Key struct {
	Name        string
	Type        KeyType
	Public      string
	Fingerprint string
}
