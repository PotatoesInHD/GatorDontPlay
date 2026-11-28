package main

import (
	"github.com/PotatoesInHD/GatorDontPlay/internal/database"
	"context"
	"fmt"
)


func middlewareLoggedIn(handler func(s *state, cmd command, dbUser database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		dbUser, err := s.dbPtr.GetUser(context.Background(), s.configPtr.CurrentUserName)
		if err != nil {
			return fmt.Errorf("db get users failed: %w", err)
		}
		return handler(s, cmd, dbUser)
	}
}
