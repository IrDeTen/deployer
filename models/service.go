package models

type Service struct {
	Name        string
	Description string
	Path        string
	Type        string
	After       string
	Requires    string
}
