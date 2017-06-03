package task

import (
	"github.com/lamproae/mra/routine"
)

type Task struct {
	Name          string
	preCondition  *routine.Routine
	postCondition *routine.Routine
	Routines      []*routine.Routine
}

func (t *Task) CheckPreCondition() error {
	return nil

}

func (t *Task) CheckPostCondition() error {
	return nil

}
