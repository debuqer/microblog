package models

type Timeline struct {
	Label        string
	Descriptiuon string
	Tags         []Tag
	LazyLoading  bool
	Events       []Event
}
