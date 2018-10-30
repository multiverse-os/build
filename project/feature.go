package project

type Features []Feature

type Feature struct {
	Subject   string
	Body      string
	Author    string
	Completed bool
}
