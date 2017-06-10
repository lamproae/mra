package ccase

type Condition struct {
	Name        string
	Assertions  []*Assertion
	Description string
}
