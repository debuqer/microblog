package models

type Date struct {
	Year  int
	Month int
	Day   int
}

type ComplexDate struct {
	AD   bool
	Date Date
}
