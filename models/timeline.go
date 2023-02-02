package models

type Timeline struct {
	Id           int
	Slug         string
	Label        string
	Descriptiuon string
	Tags         []Tag
	LazyLoading  bool
	Events       []Event
}
