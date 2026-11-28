package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/PotatoesInHD/GatorDontPlay/internal/database"
	"github.com/google/uuid"
)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("command has no arguments")
	}

	userName := cmd.args[0]
	dbUser, err := s.dbPtr.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("db query get user failed: %w", err)
	}

	if err := s.configPtr.SetUser(dbUser.Name); err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Printf("User: %s has been set\n", dbUser.Name)

	return nil
}


func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("command has no argument")
	}

	userName := cmd.args[0]
	timeNow := time.Now()
	dBUserParams := database.CreateUserParams {
		ID:	uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name: userName,
	}

	dbUser, err := s.dbPtr.CreateUser(context.Background(), dBUserParams)
	if err != nil {
		return fmt.Errorf("db create user failed: %w", err)
	}

	if err := s.configPtr.SetUser(dbUser.Name); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}


func handlerUsers(s *state, cmd command) error {
	dbUsers, err := s.dbPtr.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db get list of users failed: %w", err)
	}

	for _, user := range dbUsers {
		if user.Name == s.configPtr.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
