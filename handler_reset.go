package main

import (
	"fmt"
	"context"
)


func handlerReset(s *state, cmd command) error {
	err := s.dbPtr.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db delete users failed: %w", err)
	}
	fmt.Println("Database: Users table was reset successfully!")

	return nil
}
