package command

type Command struct {
	mode   string
	cmd    string
	result string
}

func (c *Command) CMD() string {
	return c.cmd
}

func (c *Command) Result() string {
	return c.result
}

func (c *Command) RequiredMode() string {
	return c.mode
}
