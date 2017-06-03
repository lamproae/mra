package routine

import (
	"github.com/lamproae/mra/assert"
	"github.com/lamproae/mra/command"
)

type Routine struct {
	Name   string
	CMD    []*command.Command
	Assert []*assert.Assert
}
