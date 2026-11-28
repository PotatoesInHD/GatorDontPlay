package main

import (
	"errors"
	"context"
	"fmt"
)


func handlerAgg(s *state, cmd command) error {
  	if len(cmd.args) == 0 {
		return errors.New("command has no argument")
		}
	_, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
		if err != nil {
			return fmt.Errorf("fetch feed failed: %w", err)
		}

	return nil
}
