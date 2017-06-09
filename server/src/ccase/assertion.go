package ccase

type Assertion struct {
	DUT      string
	Mode     string
	Cli      string
	Command  Command
	Seq      string
	Expected string
}
