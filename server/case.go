package ccase

import (
	"fmt"
)

type Case struct {
	Group     string
	Feature   string
	Name      string
	Tasks     []*Task
	DUT       []*Dut
	TaskCount int
}

type Task struct {
	Name          string
	Seq           int
	PreCondition  Condition
	Routine       Routine
	PostCondition Condition
	Description   string
}

func CreateNewCase(in url.Values) (*Case, error) {
	for k, v := range in {
		fmt.Println(k, v)
	}
}
