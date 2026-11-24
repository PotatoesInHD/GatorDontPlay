package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if c == nil {
		return errors.New("nil commands pointer")
	}

	callback, ok := c.registeredCmds[cmd.name];
	if !ok {
		return errors.New("command not found")
	}

	err := callback(s, cmd)
	return err
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.registeredCmds[name] = f
}
