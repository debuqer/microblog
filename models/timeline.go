package models

type Timeline struct {
	Id           int
	Slug         string
	Title        string
	Descriptiuon string
	CreatedBy    int
	CreatedDate  string
	Tags         []Tag
	LazyLoading  bool
	Events       []Event
}
