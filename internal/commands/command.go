package commands

import "context"

type Command struct {
	cronExpression string
	name           string
	runFunc        func(ctx context.Context)
}

func NewCommand(cronExpression string, name string, runFunc func(ctx context.Context)) *Command {
	return &Command{
		cronExpression: cronExpression,
		name:           name,
		runFunc:        runFunc,
	}
}

func (c *Command) GetCronExpression() string {
	return c.cronExpression
}

func (c *Command) Run(ctx context.Context) {
	c.runFunc(ctx)
}

func (c *Command) GetName() string {
	return c.name
}
