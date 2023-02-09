package models

type Date struct {
	Year  int
	Month int
	Day   int
}

const AD = 1
const BC = 2
const MYA = 3

type ComplexDate struct {
	Era  int
	Date Date
}
