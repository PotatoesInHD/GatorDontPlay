package main

import (
	"fmt"
	"errors"
)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("command has no arguments")
	}
	if err := s.configPtr.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Printf("User: %s has been set", cmd.args[0])
	return nil
}
