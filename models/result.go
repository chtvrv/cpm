package models

type Result struct {
	WorkName    string
	Duration    int
	EarlyStart  int
	LateStart   int
	EarlyFinish int
	LateFinish  int
	TimeMargin  int
}
