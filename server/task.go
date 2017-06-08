package main

type Case struct {
	Group     string
	Feature   string
	Name      string
	Tasks     []*Task
	DUT       []*Dut
	TaskCount int
}

type Dut struct {
	Name string
}

type Task struct {
	Name          string
	Seq           int
	PreCondition  Condition
	Routine       Routine
	PostCondition Condition
	Description   string
}

type Condition struct {
	Assertions  []*Assertion
	Description string
}

type Routine struct {
	Assertions  []*Assertion
	Description string
}

type Command struct {
	Mode string
	Cli  string
}

type Assertion struct {
	DUT      Dut
	Command  Command
	Expected string
}
